package serialization

import (
	"shared/model"
	"strings"
)

func WriteAll(works []model.Work, composers []model.Composer, perfs []model.Performer, recs []model.Recording) string {
	var b strings.Builder

	b.WriteString(writeComposers(composers))
	b.WriteString(writeWorks(works, composers))

	b.WriteString(writePerformers(perfs))
	b.WriteString(writeRecordings(recs))

	return b.String()
}
