package kobomatic

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

type ReadingState struct {
	EntitlementID                string
	LastModified                 string
	Status                       string
	SpentReadingMinutes          int
	RemainingTimeMinutes         int
	ProgressPercent              int
	ContentSourceProgressPercent int
	Value                        string
	Type                         string
	Source                       string
}

func GetReadingState(db *sql.DB, entitlementID string) (ReadingState, error) {
	query := sq.
		Select(
			"entitlement_id",
			"last_modified",
			"status",
			"spent_reading_minutes",
			"remaining_time_minutes",
			"progress_percent",
			"content_source_progress_percent",
			"value",
			"type",
			"source",
		).
		From("reading_state").
		Where(sq.Eq{"entitlement_id": entitlementID})

	row := query.RunWith(db).QueryRow()
	var i ReadingState
	err := row.Scan(
		&i.EntitlementID,
		&i.LastModified,
		&i.Status,
		&i.SpentReadingMinutes,
		&i.RemainingTimeMinutes,
		&i.ProgressPercent,
		&i.ContentSourceProgressPercent,
		&i.Value,
		&i.Type,
		&i.Source,
	)
	return i, err
}

func GetReadingStatesByLastModified(db *sql.DB, lastModified string) ([]ReadingState, error) {
	query := sq.
		Select(
			"entitlement_id",
			"last_modified",
			"status",
			"spent_reading_minutes",
			"remaining_time_minutes",
			"progress_percent",
			"content_source_progress_percent",
			"value",
			"type",
			"source").
		From("reading_state").
		Where(sq.Gt{"last_modified": lastModified})

	rows, err := query.RunWith(db).Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ReadingState
	for rows.Next() {
		var i ReadingState
		if err := rows.Scan(
			&i.EntitlementID,
			&i.LastModified,
			&i.Status,
			&i.SpentReadingMinutes,
			&i.RemainingTimeMinutes,
			&i.ProgressPercent,
			&i.ContentSourceProgressPercent,
			&i.Value,
			&i.Type,
			&i.Source,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func GetReadingStateLastModified(db *sql.DB, entitlementID string) (string, error) {
	query := sq.
		Select("last_modified").
		From("reading_state").
		Where(sq.Eq{"entitlement_id": entitlementID})

	row := query.RunWith(db).QueryRow()
	var lastModified string
	err := row.Scan(&lastModified)
	return lastModified, err
}

func InsertReadingState(db *sql.DB, s ReadingState) error {
	query := sq.
		Replace("reading_state").
		Columns(
			"entitlement_id",
			"last_modified",
			"status",
			"spent_reading_minutes",
			"remaining_time_minutes",
			"progress_percent",
			"content_source_progress_percent",
			"value",
			"type",
			"source",
		).
		Values(
			s.EntitlementID,
			s.LastModified,
			s.Status,
			s.SpentReadingMinutes,
			s.RemainingTimeMinutes,
			s.ProgressPercent,
			s.ContentSourceProgressPercent,
			s.Value,
			s.Type,
			s.Source,
		)
	_, err := query.
		RunWith(db).
		Exec()
	return err
}
