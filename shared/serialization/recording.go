package serialization

import (
	"fmt"
	"shared/model"
	"strings"
)

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

func writeRecordedMovements(recMovs []model.RecordedMovement) string {
	var b strings.Builder
	b.WriteString(getRecordedMovementsInsertHeader())

	total := len(recMovs)
	for i, r := range recMovs {
		b.WriteString(writeRecording(r))

		if i < total - 1 {
			b.WriteByte(',')
		} else {
			b.WriteByte(';')
		}
	}

	return b.String()
}

func writeRecodedMovement(recMov model.RecordedMovement) string {
	return fmt.Sprintf("()")
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
