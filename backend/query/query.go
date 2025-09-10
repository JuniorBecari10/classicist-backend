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
	
	if err := row.Scan(&id, &composer.Name, &composer.BirthYear, &deathYear, &composer.PhotoPath); err != nil {
		if err == sql.ErrNoRows {
			return model.Composer{}, fmt.Errorf("Composer not found")
		}

		return model.Composer{}, err
	}

	composer.DeathYear = option.FromSQL(int(deathYear.Int64), deathYear.Valid)
	return composer, nil
}
