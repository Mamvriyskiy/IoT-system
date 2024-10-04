-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS client (
    clientID SERIAL,
    password varchar(255),
    login varchar(255),
    email varchar(255)
);

CREATE TABLE IF NOT EXISTS access (
    accessID SERIAL,
    clientID int,
    homeID int,
    accessStatus varchar(15),
    accessLevel int
);

CREATE TABLE IF NOT EXISTS home (
    homeID SERIAL,
    latitude REAL,
    longitude REAL,
    name varchar(20)
);

CREATE TABLE IF NOT EXISTS device (
    deviceID SERIAL,
    homeID int,
    name varchar(20),
    typeDevice varchar(20),
    status varchar(10),
    brand varchar(15)
);

CREATE TABLE IF NOT EXISTS deviceCharacteristics (
    characterid serial,
    deviceID int,
    valuesChar DECIMAL,
    typecharacterid int
);

CREATE TABLE IF NOT EXISTS typeCharacter (
    typecharacterid serial,
    typecharacter varchar(25),
    unitmeasure varchar(15)
);

CREATE TABLE IF NOT EXISTS historyDev (
    historyDevID SERIAL,
    timeWork int,
    AverageIndicator decimal,
    EnergyConsumed int
);

CREATE TABLE IF NOT EXISTS historyDevice (
    historyDevID int,
    deviceID int
);

CREATE TABLE IF NOT EXISTS resetPswrd (
    resetCode varchar(6),
    clientID int,
    token text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- DROP TABLE accessClient,
-- accessHome,
-- deviceHome,
-- historyDevice,
-- historydev,
-- device,
-- home,
-- access,
-- resetPswrd,
-- client,
-- devicecharacteristics,
-- typecharacter,
-- clientHome;
-- +goose StatementEnd
