package serialization

import (
	"fmt"
	"shared/model"
	"slices"
	"strings"
)

// --- Works ---

func writeWorks(works []model.Work, composers []model.Composer) string {
	var b strings.Builder

	b.WriteString(writeWorksData(works, composers))
	b.WriteString(writeMovements(works))
	b.WriteString(writeTempoMarkings(works))
	b.WriteString(writeLyrics(works))

	return b.String()
}

func writeWorksData(works []model.Work, composers []model.Composer) string {
	var b strings.Builder
	b.WriteString(getWorkDataInsertHeader())

	total := len(works)

	for i, w := range works {
		b.WriteString(writeWork(w, composers))

		if i < total - 1 {
			b.WriteByte(',')
		} else {
			b.WriteByte(';')
		}
	}

	return b.String()
}

func writeWork(w model.Work, composers []model.Composer) string {
	return fmt.Sprintf("(%s, %s, %s, %d, %d, %d, %s, %s, %s, %d, %s)",
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
	)
}

func writeMovements(works []model.Work) string {
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
			b.WriteString(writeMovement(mov, i, j))

			if count < total {
				b.WriteByte(',')
			} else {
				b.WriteByte(';')
			}
		}
	}

	return b.String()
}

func writeMovement(mov model.Movement, i, j int) string {
	return fmt.Sprintf("(%d, %d, %s, %s, %s)", i + 1, j + 1,
		OptionalToSQL(mov.Form), OptionalToSQL(mov.Nickname), StringToSQL(mov.Sheet.Path))
}

func writeTempoMarkings(works []model.Work) string {
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
	movId := 0

	for _, w := range works {
		for _, mov := range w.Movements {
			movId++

			for j, tempo := range mov.TempoMarkings {
				count++
				b.WriteString(writeTempoMarking(tempo, movId, j))

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

func writeTempoMarking(tempo model.TempoMarking, movId, j int) string {
	return fmt.Sprintf("(%d, %d, %s, %s)", movId, j + 1, OptionalToSQL(tempo.Form), StringToSQL(tempo.Name))
}

func writeLyrics(works []model.Work) string {
	var b strings.Builder
	b.WriteString(getLyricInsertHeader())

	// count total of lyrics
	total := 0
	for _, w := range works {
		for _, mov := range w.Movements {
			if lyrics, ok := mov.Lyrics.TryUnwrap(); ok {
				total += len(lyrics)
			}
		}
	}

	count := 0
	movId := 0
	for _, w := range works {
		for _, mov := range w.Movements {
			movId++
			
			if lyrics, ok := mov.Lyrics.TryUnwrap(); ok {
				for j, lyric := range lyrics {
					count++
					b.WriteString(writeLyric(lyric, movId, j))

					if count < total {
						b.WriteByte(',')
					} else {
						b.WriteByte(';')
					}
				}
			}
		}
	}

	return b.String()
}

func writeLyric(lyric string, movId, j int) string {
	return fmt.Sprintf("(%d, %d, %s)", movId, j + 1, StringToSQL(lyric))
}

// --- Composers ---

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

// --- Headers ---

func getWorkDataInsertHeader() string {
	return `INSERT INTO works (
    title_kind, title_number, title_nickname,
    key_note, key_mode, composer_id,
    catalog_prefix, catalog_number, catalog_subnumber,
    composition_start_year, composition_end_year) VALUES `
}

func getMovementInsertHeader() string {
	return "INSERT INTO movements (work_id, order_num, form, nickname, sheet_path) VALUES "
}

func getTempoMarkingInsertHeader() string {
	return "INSERT INTO tempo_markings (movement_id, order_num, form, name) VALUES "
}

func getLyricInsertHeader() string {
	return "INSERT INTO lyrics (movement_id, order_num, line) VALUES "
}

func getComposerInsertHeader() string {
	return "INSERT INTO composers (name, birth_year, death_year, photo_path) VALUES "
}
