-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS buildings (
id  SERIAL PRIMARY KEY,
name VARCHAR(255) NOT NULL,
city VARCHAR(255) NOT NULL,
year_completed INT NOT NULL,
floors INT NOT NULL
);

CREATE INDEX idx_city ON Buildings (city);
CREATE INDEX idx_year_completed ON Buildings (year_completed);
CREATE INDEX idx_floors ON Buildings (floors);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE IF EXISTS buildings;
-- +goose StatementEnd
