package serialization

import (
	"fmt"
	"shared/model"
	"strings"
)

// --- Performers ---

func writePerformers(perfs []model.Performer) string {
	var b strings.Builder
	b.WriteString(getPerformersInsertHeader())

	total := len(perfs)
	for i, p := range perfs {
		b.WriteString(writePerformer(p))

		if i < total - 1 {
			b.WriteByte(',')
		} else {
			b.WriteByte(';')
		}
	}

	return b.String()
}

func writePerformer(perf model.Performer) string {
	return fmt.Sprintf("(%s)", StringToSQL(perf.Name))
}

// --- Recordings ---

func writeRecordings(recs []model.Recording) string {
	var b strings.Builder

	b.WriteString(writeRecordingsData(recs))

	for _, rec := range recs {
		b.WriteString(writeRecordedMovements(rec.Id, rec.Movements))
		b.WriteString(writeRecordingPerformers(rec.Id, rec.Performers))
	}

	return b.String()
}

func writeRecordingsData(recs []model.Recording) string {
	var b strings.Builder
	b.WriteString(getRecordingsInsertHeader())

	total := len(recs)
	for i, r := range recs {
		b.WriteString(writeRecording(r))

		if i < total - 1 {
			b.WriteByte(',')
		} else {
			b.WriteByte(';')
		}
	}

	return b.String()
}

func writeRecording(rec model.Recording) string {
	return fmt.Sprintf("(%d, %d, %s)", rec.WorkId, rec.Year, StringToSQL(rec.PhotoPath))
}

// --- Recording Movements ---

func writeRecordedMovements(recordingId int, recMovs []model.RecordedMovement) string {
	var b strings.Builder
	b.WriteString(getRecordedMovementsInsertHeader())

	total := len(recMovs)
	for i, r := range recMovs {
		b.WriteString(writeRecordedMovement(recordingId, r))

		if i < total - 1 {
			b.WriteByte(',')
		} else {
			b.WriteByte(';')
		}
	}

	return b.String()
}

func writeRecordedMovement(recordingId int, recMov model.RecordedMovement) string {
	return fmt.Sprintf("(%d, %d, %s, %d)",
		recMov.MovementId,
		recordingId,
		StringToSQL(recMov.AudioFile.Path),
		recMov.AudioFile.Duration,
	)
}

// --- Recording Performers ---

func writeRecordingPerformers(recordingId int, recPerfs []model.RecordingPerformer) string {
	var b strings.Builder
	b.WriteString(getRecordingPerformersInsertHeader())

	total := len(recPerfs)
	for i, rp := range recPerfs {
		b.WriteString(writeRecordingPerformer(recordingId, rp))

		if i < total - 1 {
			b.WriteByte(',')
		} else {
			b.WriteByte(';')
		}
	}

	return b.String()
}

func writeRecordingPerformer(recordingId int, rp model.RecordingPerformer) string {
	return fmt.Sprintf("(%d, %d, %s)",
		recordingId,
		rp.Performer.Id,
		StringToSQL(rp.Role),
	)
}

// --- Headers ---

func getPerformersInsertHeader() string {
	return "INSERT INTO performers (name) VALUES "
}

func getRecordingsInsertHeader() string {
	return "INSERT INTO recordings (work_id, year, photo_path) VALUES "
}

func getRecordedMovementsInsertHeader() string {
	return "INSERT INTO recorded_movements (movement_id, recording_id, audio_path, duration) VALUES "
}

func getRecordingPerformersInsertHeader() string {
	return "INSERT INTO recording_performers (recording_id, performer_id, role) VALUES "
}
