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
	var b strings.Builder
	b.WriteString(GetWorkInsertHeader())

}

func WriteWorkMovements(w model.Work) string {
	var b strings.Builder
	b.WriteString(GetWorkInsertHeader())

}

func WriteWorkTempoMarkings(w model.Work) string {
	var b strings.Builder
	b.WriteString(GetWorkInsertHeader())

}

func WriteWorks(works []model.Work) string {
	var b strings.Builder

	for _, w := range works {
		b.WriteString(WriteWork(w))
	}

	return b.String()
}

func GetWorkInsertHeader() string {
	return `INSERT INTO works (
    title_kind, title_number, title_nickname,
    key_note, key_mode, composer_id,
    catalog_prefix, catalog_number, catalog_subnumber,
    composition_start_year, composition_end_year) VALUES `
}

func GetMovementInsertHeader() string {
	return "INSERT INTO movements (work_id, order_num, kind) VALUES "
}

func GetTempoMarkingInsertHeader() string {
	return "INSERT INTO tempo_markings (movement_id, order_num, name) VALUES "
}

// ---

func WriteComposer(composer model.Composer) string {
	var b strings.Builder

	b.WriteString(GetComposerInsertHeader())
	b.WriteString(WriteComposerData(composer))
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
