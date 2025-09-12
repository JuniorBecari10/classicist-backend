package serialization

import (
	"fmt"
	"shared/model"
	"slices"
	"strings"
)

// -------------------- Works --------------------

func writeWorks(works []model.Work, composers []model.Composer) string {
	var b strings.Builder

	b.WriteString(writeWorksData(works, composers))
	b.WriteString(writeWorksMovements(works))
	b.WriteString(writeWorksTempoMarkings(works))

	return b.String()
}

func writeWorksData(works []model.Work, composers []model.Composer) string {
	var b strings.Builder
	b.WriteString(getWorkDataInsertHeader())

	total := len(works)

	for i, w := range works {
		b.WriteString(
			fmt.Sprintf("(%s, %s, %s, %d, %d, %d, %s, %s, %s, %d, %s, %s)",
				StringToSQL(w.Title.Kind),
				OptionalToSQL(w.Title.Number),
				OptionalToSQL(w.Title.Nickname),
				w.Key.Note, w.Key.Mode,
				slices.Index(composers, w.Composer) + 1, // assumes composer is in the list
				StringToSQL(w.Catalog.Prefix),
				OptionalToSQL(w.Catalog.Number),
				OptionalToSQL(w.Catalog.Subnumber),
				w.Year.StartYear,
				OptionalToSQL(w.Year.EndYear),
				StringToSQL(w.Sheet.Path),
			),
		)

		if i < total - 1 {
			b.WriteByte(',')
		} else {
			b.WriteByte(';')
		}
	}

	return b.String()
}

func writeWorksMovements(works []model.Work) string {
	var b strings.Builder
	b.WriteString(getMovementInsertHeader())

	// count total movements
	total := 0
	for _, w := range works {
		total += len(w.Movements)
	}

	count := 0
	for i, w := range works {
		for j, mov := range w.Movements {
			count++
			b.WriteString(fmt.Sprintf("(%d, %d, %s, %s)", i + 1, j + 1, OptionalToSQL(mov.Form), OptionalToSQL(mov.Nickname)))

			if count < total {
				b.WriteByte(',')
			} else {
				b.WriteByte(';')
			}
		}
	}

	return b.String()
}

func writeWorksTempoMarkings(works []model.Work) string {
	var b strings.Builder
	b.WriteString(getTempoMarkingInsertHeader())

	// count total tempo markings
	total := 0
	for _, w := range works {
		for _, mov := range w.Movements {
			total += len(mov.TempoMarkings)
		}
	}

	count := 0
	movementID := 0
	for _, w := range works {
		for _, mov := range w.Movements {
			movementID++
			for j, tempo := range mov.TempoMarkings {
				count++
				b.WriteString(fmt.Sprintf("(%d, %d, %s)", movementID, j + 1, StringToSQL(tempo.Name)))

				if count < total {
					b.WriteByte(',')
				} else {
					b.WriteByte(';')
				}
			}
		}
	}

	return b.String()
}

// -------------------- Composers --------------------

func writeComposers(composers []model.Composer) string {
	var b strings.Builder
	b.WriteString(getComposerInsertHeader())

	total := len(composers)
	for i, c := range composers {
		b.WriteString(writeComposerData(c))

		if i < total - 1 {
			b.WriteByte(',')
		} else {
			b.WriteByte(';')
		}
	}

	return b.String()
}

func writeComposerData(c model.Composer) string {
	return fmt.Sprintf("(%s, %d, %s, %s)",
		StringToSQL(c.Name),
		c.BirthYear,
		OptionalToSQL(c.DeathYear),
		StringToSQL(c.PhotoPath),
	)
}

// -------------------- Headers --------------------

func getWorkDataInsertHeader() string {
	return `INSERT INTO works (
    title_kind, title_number, title_nickname,
    key_note, key_mode, composer_id,
    catalog_prefix, catalog_number, catalog_subnumber,
    composition_start_year, composition_end_year, sheet_path) VALUES `
}

func getMovementInsertHeader() string {
	return "INSERT INTO movements (work_id, order_num, form, nickname) VALUES "
}

func getTempoMarkingInsertHeader() string {
	return "INSERT INTO tempo_markings (movement_id, order_num, name) VALUES "
}

func getComposerInsertHeader() string {
	return "INSERT INTO composers (name, birth_year, death_year, photo_path) VALUES "
}

// -------------------- Entry Point --------------------

func WriteAll(works []model.Work, composers []model.Composer) string {
	var b strings.Builder

	b.WriteString(writeComposers(composers))
	b.WriteString(writeWorks(works, composers))

	return b.String()
}
