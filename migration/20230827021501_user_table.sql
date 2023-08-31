-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "user"
(
    id   SERIAL NOT NULL PRIMARY KEY,
    name text   NOT NULL
);



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "user"
-- +goose StatementEnd
