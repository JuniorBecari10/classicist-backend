package query

import (
	"database/sql"
	"shared/model"
	"shared/option"
	"slices"
	"strings"

	_ "embed"
)

//go:embed sql/search/work/work.sql
var searchWorkQuery string

//go:embed sql/search/work/composer.sql
var searchComposerQuery string

//go:embed sql/search/recording/performer.sql
var searchPerformerQuery string

func SearchDatabase(db *sql.DB, term string) ([]model.SearchResult, error) {
	works, err := searchWorks(db, term)
	if err != nil {
		return nil, err
	}

	composers, err := searchComposers(db, term)
	if err != nil {
		return nil, err
	}

	performers, err := searchPerformers(db, term)
	if err != nil {
		return nil, err
	}

	// to ensure it's not nil, and rather an empty list
	concat := slices.Concat(works, composers, performers)
	if concat == nil {
		concat = []model.SearchResult{}
	}

	return concat, nil
}

func searchWorks(db *sql.DB, term string) ([]model.SearchResult, error) {
	var termIds [][]int

	// separate by words, because in the query order matters, and here it doesn't.
	// perform one query per word, in a list of results (list of list of IDs)
	for _, part := range strings.Split(term, " ") {
		list, err := queryWorkTerm(db, part)
		if err != nil {
			return nil, err
		}

		termIds = append(termIds, list)
	}

	// take the intersection of them, ensuring there are no duplicates
	// (take the common items only in all of the lists)
	filtered := Intersection(termIds...)
	var works []model.Work

	// use the list of IDs to properly query each work
	for _, id := range filtered {
		work, err := QueryWork(db, id)
		if err != nil {
			return nil, err
		}

		works = append(works, work)
	}

	// convert each work to a SearchResult, and then, return it
	results := Map(works, func (work model.Work) model.SearchResult {
		return model.NewWorkSR(work)
	})

	return results, nil
}

func searchComposers(db *sql.DB, term string) ([]model.SearchResult, error) {
	rows, err := db.Query(searchComposerQuery, term)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	composers := []model.Composer{}

	for rows.Next() {
		var (
			composer model.Composer
			deathYear sql.NullInt64
		)
		
		if err := rows.Scan(&composer.Id, &composer.Name, &composer.BirthYear, &deathYear, &composer.PhotoPath); err != nil {
			if err == sql.ErrNoRows {
				// it's ok not to find!
				continue
			}

			return nil, err
		}

		composer.DeathYear = option.FromSQL(int(deathYear.Int64), deathYear.Valid)
		composers = append(composers, composer)
	}

	results := Map(composers, func (composer model.Composer) model.SearchResult {
		return model.NewComposerSR(composer)
	})

	return results, nil
}

func searchPerformers(db *sql.DB, term string) ([]model.SearchResult, error) {
	rows, err := db.Query(searchPerformerQuery, term)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	performers := []model.Performer{}

	for rows.Next() {
		var perf model.Performer
		
		if err := rows.Scan(&perf.Id, &perf.Name); err != nil {
			if err == sql.ErrNoRows {
				// it's ok not to find!
				continue
			}

			return nil, err
		}

		performers = append(performers, perf)
	}

	results := Map(performers, func (perf model.Performer) model.SearchResult {
		return model.NewPerformerSR(perf)
	})

	return results, nil
}

// ---

// List of IDs
func queryWorkTerm(db *sql.DB, term string) ([]int, error) {
	rows, err := db.Query(searchWorkQuery, term)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	ids := []int{}

	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			if err == sql.ErrNoRows {
				// it's ok not to find!
				continue
			}

			return nil, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}
