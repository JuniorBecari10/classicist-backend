package serialization

import (
	"fmt"
	"shared/model"
	"slices"
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

func writeRecordings(recs []model.Recording, performers []model.Performer) string {
	var b strings.Builder

	b.WriteString(writeRecordingsData(recs))

	for i, rec := range recs {
		b.WriteString(writeRecordedMovements(rec.Movements, i + 1))
		b.WriteString(writeRecordingPerformers(rec.Performers, performers, i + 1))
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

func writeRecordedMovements(recMovs []model.RecordedMovement, recordingId int) string {
	var b strings.Builder
	b.WriteString(getRecordedMovementsInsertHeader())

	total := len(recMovs)
	for i, r := range recMovs {
		b.WriteString(writeRecordedMovement(r, recordingId))

		if i < total - 1 {
			b.WriteByte(',')
		} else {
			b.WriteByte(';')
		}
	}

	return b.String()
}

func writeRecordedMovement(recMov model.RecordedMovement, recordingId int) string {
	return fmt.Sprintf("(%d, %d, %s, %d)",
		recMov.MovementId,
		recordingId,
		StringToSQL(recMov.AudioFile.Path),
		recMov.AudioFile.Duration,
	)
}

// --- Recording Performers ---

func writeRecordingPerformers(recPerfs []model.RecordingPerformer, performers []model.Performer, recordingId int) string {
	var b strings.Builder
	b.WriteString(getRecordingPerformersInsertHeader())

	total := len(recPerfs)
	for i, rp := range recPerfs {
		b.WriteString(writeRecordingPerformer(rp, performers, recordingId))

		if i < total - 1 {
			b.WriteByte(',')
		} else {
			b.WriteByte(';')
		}
	}

	return b.String()
}

func writeRecordingPerformer(rp model.RecordingPerformer, performers []model.Performer, recordingId int) string {
	return fmt.Sprintf("(%d, %d, %s)",
		recordingId,
		slices.Index(performers, rp.Performer) + 1, // assumes it's not -1
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
