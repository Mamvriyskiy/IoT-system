-- +goose Up
-- +goose StatementBegin
CREATE INDEX index_login ON client (login);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
