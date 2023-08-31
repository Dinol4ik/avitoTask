-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_segment
(
    id SERIAL NOT NULL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES "user" ON DELETE CASCADE ,
    segment_id INTEGER NOT NULL REFERENCES "segment" ON DELETE CASCADE,
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp NOT NULL DEFAULT NOW(),
    is_removed BOOLEAN DEFAULT false,


    UNIQUE (user_id, segment_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_segment;
-- +goose StatementEnd
