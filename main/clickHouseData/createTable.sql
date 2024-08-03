CREATE TABLE IF NOT EXISTS client (
    clientID UInt32,
    password String,
    login String,
    email String
) ENGINE = MergeTree()
ORDER BY clientID;

CREATE TABLE IF NOT EXISTS access (
    accessID UInt32,
    accessStatus String,
    accessLevel Int32
) ENGINE = MergeTree()
ORDER BY accessID;

CREATE TABLE IF NOT EXISTS accessClient (
    clientID UInt32,
    accessID UInt32
) ENGINE = MergeTree()
ORDER BY clientID;

CREATE TABLE IF NOT EXISTS accessHome (
    accessID UInt32,
    homeID UInt32
) ENGINE = MergeTree()
ORDER BY accessID;

CREATE TABLE IF NOT EXISTS home (
    homeID UInt32,
    ownerID UInt32,
    name String
) ENGINE = MergeTree()
ORDER BY homeID;

CREATE TABLE IF NOT EXISTS device (
    deviceID UInt32,
    name String,
    typeDevice String,
    status String,
    brand String,
    maxParametr Int32,
    minParametr Int32,
    powerConsumption Int32
) ENGINE = MergeTree()
ORDER BY deviceID;

CREATE TABLE IF NOT EXISTS deviceHome (
    homeID UInt32,
    deviceID UInt32
) ENGINE = MergeTree()
ORDER BY homeID;

CREATE TABLE IF NOT EXISTS historyDev (
    historyDevID UInt32,
    timeWork Int32,
    AverageIndicator Decimal32(2),
    EnergyConsumed Int32
) ENGINE = MergeTree()
ORDER BY historyDevID;

CREATE TABLE IF NOT EXISTS historyDevice (
    historyDevID UInt32,
    deviceID UInt32
) ENGINE = MergeTree()
ORDER BY historyDevID;
