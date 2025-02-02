-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS reading_state (
	entitlement_id TEXT PRIMARY KEY,
	created TEXT NOT NULL,
	last_modified TEXT NOT NULL,
	status TEXT NOT NULL,
	spent_reading_minutes INTEGER NOT NULL,
	remaining_time_minutes INTEGER NOT NULL,
	progress_percent INTEGER NOT NULL,
	content_source_progress_percent INTEGER NOT NULL,
	value TEXT NOT NULL,
	type TEXT NOT NULL,
	source TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS reading_state;
-- +goose StatementEnd
