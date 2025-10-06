package query

import (
	"database/sql"
	"fmt"
	"shared/model"

	_ "embed"
)

// --- Queries ---

//go:embed sql/recording/recorded_movements.sql
var recordedMovementsQuery string

//go:embed sql/recording/recordings_for_work.sql
var recordingsByWorkIdQuery string

//go:embed sql/recording/recording_performers.sql
var recordingPerformersQuery string

//go:embed sql/recording/performer.sql
var performerQuery string

func QueryRecordingsForWork(db *sql.DB, workId int) ([]model.Recording, error) {
	rows, err := db.Query(recordingsByWorkIdQuery, workId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	recs := []model.Recording{}

	for rows.Next() {
		var rec model.Recording

		if err := rows.Scan(&rec.Id, &rec.WorkId, &rec.Year, &rec.PhotoPath); err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("Recording not found")
			}
			
			return nil, err
		}

		err := queryRecordingDetails(db, &rec)
		if err != nil {
			return nil, err
		}

		recs = append(recs, rec)
	}

	if len(recs) == 0 {
		return nil, fmt.Errorf("Recording not found")
	} else {
		return recs, nil
	}
}

func queryRecordingDetails(db *sql.DB, rec *model.Recording) error {
	recMovs, err := queryRecordedMovements(db, rec)
	if err != nil {
		return err
	}

	rec.Movements = recMovs

	recPerfs, err := queryRecordingPerformers(db, rec)
	if err != nil {
		return err
	}

	rec.Performers = recPerfs
	return nil
}

// in/out parameter
func queryRecordedMovements(db *sql.DB, rec *model.Recording) ([]model.RecordedMovement, error) {
	rows, err := db.Query(recordedMovementsQuery, rec.Id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	recMovs := []model.RecordedMovement{}

	for rows.Next() {
		var (
			recMov model.RecordedMovement
			recId int // unused
		)

		if err := rows.Scan(&recMov.Id, &recMov.MovementId, &recId, &recMov.AudioFile.Path, &recMov.AudioFile.Duration); err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("Recorded movements not found")
			}
			
			return nil, err
		}

		recMovs = append(recMovs, recMov)
	}

	if len(recMovs) == 0 {
		return nil, fmt.Errorf("Recorded movements not found")
	} else {
		return recMovs, nil
	}
}

// in/out parameter
func queryRecordingPerformers(db *sql.DB, rec *model.Recording) ([]model.RecordingPerformer, error) {
	rows, err := db.Query(recordingPerformersQuery, rec.Id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	recPerfs := []model.RecordingPerformer{}

	for rows.Next() {
		var (
			recPerf model.RecordingPerformer
			recId int // unused
		)

		if err := rows.Scan(&recPerf.Id, &recId, &recPerf.Performer.Id, &recPerf.Role); err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("Recorded performers not found")
			}
			
			return nil, err
		}

		perf, err := QueryPerformer(db, recPerf.Performer.Id)
		if err != nil {
			return nil, err
		}

		recPerf.Performer = perf
		recPerfs = append(recPerfs, recPerf)
	}

	if len(recPerfs) == 0 {
		return nil, fmt.Errorf("Recorded performers not found")
	} else {
		return recPerfs, nil
	}
}

func QueryPerformer(db *sql.DB, id int) (model.Performer, error) {
	row := db.QueryRow(performerQuery, id)

	var perf model.Performer
	if err := row.Scan(&perf.Id, &perf.Name); err != nil {
		if err == sql.ErrNoRows {
			return model.Performer{}, fmt.Errorf("Performer not found")
		}

		return model.Performer{}, err
	}

	return perf, nil
}
