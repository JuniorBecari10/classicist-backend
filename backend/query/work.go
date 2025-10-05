package query

import (
	"database/sql"
	"errors"
	"fmt"
	"shared/model"

	_ "embed"
	"shared/option"
)

// --- Queries ---

//go:embed sql/work/composer_by_id.sql
var composerByIdQuery string

//go:embed sql/work/work_by_id.sql
var workByIdQuery string

//go:embed sql/work/movements_by_id.sql
var movementsByIdQuery string

//go:embed sql/work/tempo_markings_by_id.sql
var tempoMarkingsByIdQuery string

//go:embed sql/work/lyrics_by_id.sql
var lyricsByIdQuery string

func GetComposerById(db *sql.DB, id int) (model.Composer, error) {
	var (
		composer model.Composer
		deathYear sql.NullInt64
	)

	row := db.QueryRow(composerByIdQuery, id)
	
	if err := row.Scan(&composer.Id, &composer.Name, &composer.BirthYear, &deathYear, &composer.PhotoPath); err != nil {
		if err == sql.ErrNoRows {
			return model.Composer{}, fmt.Errorf("Composer not found")
		}

		return model.Composer{}, err
	}

	composer.DeathYear = option.FromSQL(int(deathYear.Int64), deathYear.Valid)
	return composer, nil
}

func GetWorkById(db *sql.DB, id int) (model.Work, error) {
	work, err := queryWork(db, id)
	if err != nil {
		return model.Work{}, err
	}

	movs, err := queryMovements(db, id)
	if err != nil {
		return model.Work{}, err
	}

	work.Movements = movs
	return work, nil
}

// ---

// it will be incomplete. complete it using another function.
// remaining field: Movements
func queryWork(db *sql.DB, id int) (model.Work, error) {
	var (
		work model.Work
		composerId int

		titleNumber sql.NullInt32
		titleNickname sql.NullString

		catalogNumber sql.NullString
		catalogSubnumber sql.NullString

		compositionEndYear sql.NullInt32
	)

	row := db.QueryRow(workByIdQuery, id)

	if err := row.Scan(
		&work.Id, &work.Title.Kind, &titleNumber, &titleNickname,
		&work.Key.Note, &work.Key.Mode, &composerId,
		&work.Catalog.Prefix, &catalogNumber, &catalogSubnumber,
		&work.Year.StartYear, &compositionEndYear,
	); err != nil {
		if err == sql.ErrNoRows {
			return model.Work{}, fmt.Errorf("Composer not found")
		}

		return model.Work{}, err
	}

	work.Title.Number = option.FromSQL(int(titleNumber.Int32), titleNumber.Valid)
	work.Title.Nickname = option.FromSQL(titleNickname.String, titleNickname.Valid)
	work.Catalog.Number = option.FromSQL(catalogNumber.String, catalogNumber.Valid)
	work.Catalog.Subnumber = option.FromSQL(catalogSubnumber.String, catalogSubnumber.Valid)
	work.Year.EndYear = option.FromSQL(int(compositionEndYear.Int32), compositionEndYear.Valid)

	composer, err := GetComposerById(db, composerId)
	if err != nil {
		return model.Work{}, err
	}

	work.Composer = composer
	return work, nil
}

func queryMovements(db *sql.DB, id int) ([]model.Movement, error) {
	rows, err := db.Query(movementsByIdQuery, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	movements := []model.Movement{}

	for rows.Next() {
		var (
			nickname, form sql.NullString
			sheetPath string
			movId, workId, orderNum int // unused for now.
		)

		if err := rows.Scan(&movId, &workId, &orderNum, &form, &nickname, &sheetPath); err != nil {
			return nil, err
		}

		tm, err := queryTempoMarkings(db, id, movId)
		if err != nil {
			return nil, err
		}

		lyrics, err := queryLyrics(db, id, movId)
		if err != nil {
			return nil, err
		}

		mov := model.Movement{
			Nickname: option.FromSQL(nickname.String, nickname.Valid),
			Form: option.FromSQL(form.String, form.Valid),
			TempoMarkings: tm,
			Lyrics: lyrics,

			Sheet: model.SheetMusic{
				Path: sheetPath,
			},
		}

		movements = append(movements, mov)
	}

	if len(movements) == 0 {
		return nil, errors.New("Movements not found")
	} else {
		return movements, nil
	}
}

func queryTempoMarkings(db *sql.DB, workId, movId int) ([]model.TempoMarking, error) {
	rows, err := db.Query(tempoMarkingsByIdQuery, workId, movId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	tempos := []model.TempoMarking{}

	for rows.Next() {
		var (
			tempo model.TempoMarking
			form sql.NullString
			id, movementId, orderNum int // unused for now.
		)

		if err := rows.Scan(&id, &movementId, &orderNum, &form, &tempo.Name); err != nil {
			return nil, err
		}

		tempo.Form = option.FromSQL(form.String, form.Valid)
		tempos = append(tempos, tempo)
	}

	if len(tempos) == 0 {
		return nil, errors.New("Tempo markings not found")
	} else {
		return tempos, nil
	}
}

// this one can return no lyrics, because it's an optional list
func queryLyrics(db *sql.DB, workId, movId int) (option.Option[[]string], error) {
	rows, err := db.Query(lyricsByIdQuery, workId, movId)
	if err != nil {
		return option.Option[[]string]{}, err
	}

	defer rows.Close()
	lyrics := []string{}

	for rows.Next() {
		var (
			line string
			id, movementId, orderNum int // unused for now.
		)

		if err := rows.Scan(&id, &movementId, &orderNum, &line); err != nil {
			return option.Option[[]string]{}, err
		}

		lyrics = append(lyrics, line)
	}

	if len(lyrics) == 0 {
		return option.None[[]string](), nil
	} else {
		return option.Some(lyrics), nil
	}
}
