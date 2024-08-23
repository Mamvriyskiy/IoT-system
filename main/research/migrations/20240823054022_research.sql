-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS client(
    clientID SERIAL, password varchar(255), login varchar(255), email varchar(255)
);

ALTER TABLE client
    ALTER COLUMN password SET NOT NULL,
    ALTER COLUMN login SET NOT NULL,
    ALTER COLUMN email SET NOT NULL;

ALTER TABLE client
    ADD CHECK (password != ''),
    ADD CHECK (login != ''),
    ADD CHECK (email != ''),
    ADD PRIMARY KEY (clientID);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE client;
-- +goose StatementEnd
