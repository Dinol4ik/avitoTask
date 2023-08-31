-- +goose Up
-- +goose StatementBegin
CREATE TABLE segment
(
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(255) UNIQUE,
    is_removed BOOLEAN DEFAULT false
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE segment
-- +goose StatementEnd
