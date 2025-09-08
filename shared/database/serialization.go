package database

import (
	"fmt"
	"shared/model"
	"strings"
)

func WriteWork(w model.Work) string {
	var b strings.Builder
	
	
}

func WriteWorkData(w model.Work) string {

}

func WriteWorkMovements(w model.Work) string {

}

func WriteWorkTempoMarkings(w model.Work) string {

}

func WriteWorks(works []model.Work) string {
	var b strings.Builder

	for _, w := range works {
		b.WriteString(WriteWork(w))
	}

	return b.String()
}

// ---

func WriteComposer(c model.Composer) string {
	var b strings.Builder

	b.WriteString(GetComposerInsertHeader())
	b.WriteString(WriteComposerData(c))
	b.WriteByte(';')

	return b.String()
}

func WriteComposers(composers []model.Composer) string {
	var b strings.Builder
	b.WriteString(GetComposerInsertHeader())

	for i, c := range composers {
		b.WriteString(WriteComposerData(c))

		if i < len(composers) - 1 {
			b.WriteByte(',')
		} else {
			b.WriteByte(';')
		}
	}

	return b.String()
}

func GetComposerInsertHeader() string {
	return "INSERT INTO composers (name, birth_year, death_year, photo_path) VALUES "
}

func WriteComposerData(c model.Composer) string {
	return fmt.Sprintf("('%s', %d, %d, '%s')",
		c.Name, c.BirthYear, c.DeathYear, c.PhotoPath)
}