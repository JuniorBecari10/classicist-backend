package serialization

import (
	"fmt"
	"shared/model"
	"strings"
)

func writePerformers(perfs []model.Performer) string {
	var b strings.Builder
	b.WriteString(getComposerInsertHeader())

	total := len(perfs)
	for i, p := range perfs {
		b.WriteString(writePerformerData(p))

		if i < total - 1 {
			b.WriteByte(',')
		} else {
			b.WriteByte(';')
		}
	}

	return b.String()
}

func writePerformerData(perf model.Performer) string {
	return fmt.Sprintf("(%s)", StringToSQL(perf.Name))
}

// --- Recordings ---

func writeRecordings(recs []model.Recording) string {

}

// --- Headers ---

func getPerformersHeader() string {
	return "INSERT INTO performers (name) VALUES"
}
