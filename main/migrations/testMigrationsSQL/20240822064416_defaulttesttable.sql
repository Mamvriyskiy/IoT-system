-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS client (
    clientID UUID,
    password varchar(255),
    login varchar(255),
    email varchar(255)
);

CREATE TABLE IF NOT EXISTS access (
    accessID UUID,
    clientID UUID,
    homeID UUID,
    accessStatus varchar(15),
    accessLevel int
);

CREATE TABLE IF NOT EXISTS home (
    homeID UUID,
    latitude REAL,
    longitude REAL,
    name varchar(20)
);

CREATE TABLE IF NOT EXISTS device (
    deviceID UUID,
    homeID UUID,
    name varchar(20),
    typeDevice varchar(20),
    status varchar(10),
    brand varchar(15)
);

CREATE TABLE IF NOT EXISTS deviceCharacteristics (
    characterid UUID,
    deviceID UUID,
    valuesChar DECIMAL,
    typecharacterid UUID
);

CREATE TABLE IF NOT EXISTS typeCharacter (
    typecharacterid UUID,
    typecharacter varchar(25),
    unitmeasure varchar(15)
);

CREATE TABLE IF NOT EXISTS historyDev (
    historyDevID UUID,
    timeWork int,
    AverageIndicator decimal,
    EnergyConsumed int
);

CREATE TABLE IF NOT EXISTS historyDevice (
    historyDevID UUID,
    deviceID UUID
);

CREATE TABLE IF NOT EXISTS resetPswrd (
    resetCode varchar(6),
    clientID UUID,
    token text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE historyDevice, historydev, device, access, home, resetPswrd, client, devicecharacteristics, typecharacter;
-- +goose StatementEnd
