-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS goods (
	id SERIAL PRIMARY KEY,
	name VARCHAR(100),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS goods;
-- +goose StatementEnd
