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
	rows, err := db.Query(workByIdQuery, id)
	if err != nil {
		return model.Work{}, err
	}

	defer rows.Close()

	var (
		work model.Work
		composerId int

		titleNumber sql.NullInt32
		titleNickname sql.NullString

		catalogNumber sql.NullString
		catalogSubnumber sql.NullString

		compositionEndYear sql.NullInt32
	)

	row := db.QueryRow(composerByIdQuery, id)
	
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

	composer, err := GetComposerById(db, composerId)
	if err != nil {
		return model.Work{}, err
	}

	work.Composer = composer
	return work, nil
}

func queryMovements(db *sql.DB, id int) ([]model.Movement, error) {

}
