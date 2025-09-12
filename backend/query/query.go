package query

import (
	"database/sql"
	"fmt"
	"shared/model"
	"shared/option"

	_ "embed"
)

//go:embed sql/composer_by_id.sql
var composerByIdQuery string

//go:embed sql/work_by_id.sql
var workByIdQuery string

//go:embed sql/movements_by_id.sql
var movementsByIdQuery string

//go:embed sql/tempo_markings_by_id.sql
var tempoMarkingsByIdQuery string

func GetComposerById(db *sql.DB, id int) (model.Composer, error) {
	rows, err := db.Query(composerByIdQuery, id)
	if err != nil {
		return model.Composer{}, err
	}

	defer rows.Close()

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
	rows, err := db.Query(composerByIdQuery, id)
	if err != nil {
		return model.Work{}, err
	}

	defer rows.Close()

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
		&work.Year.StartYear, &compositionEndYear, &work.Sheet.Path,
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
			id, workId, orderNum int // unused for now.
		)

		if err := rows.Scan(&id, &workId, &orderNum, &form, &nickname); err != nil {
			return nil, err
		}

		tm, err := queryTempoMarkings(db, id)
		if err != nil {
			return nil, err
		}

		mov := model.Movement{
			Nickname: option.FromSQL(nickname.String, nickname.Valid),
			Form: option.FromSQL(form.String, form.Valid),
			TempoMarkings: tm,
		}

		movements = append(movements, mov)
	}

	return movements, nil
}

func queryTempoMarkings(db *sql.DB, id int) ([]model.TempoMarking, error) {
	rows, err := db.Query(tempoMarkingsByIdQuery, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	tempos := []model.TempoMarking{}

	for rows.Next() {
		var (
			tempo model.TempoMarking
			id, movementId, orderNum int // unused for now.
		)

		if err := rows.Scan(&id, &movementId, &tempo.Name, &orderNum); err != nil {
			return nil, err
		}


		tempos = append(tempos, tempo)
	}

	return tempos, nil
}
