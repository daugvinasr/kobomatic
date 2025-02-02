-- +goose Up
-- +goose StatementBegin
ALTER TABLE reading_state DROP COLUMN created;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE reading_state ADD COLUMN created TEXT NOT NULL DEFAULT '';
-- +goose StatementEnd