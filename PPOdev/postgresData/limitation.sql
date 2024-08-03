ALTER TABLE client
ALTER COLUMN password
SET NOT NULL,
ALTER COLUMN login
SET NOT NULL,
ALTER COLUMN email
SET NOT NULL,
ADD check (password != ''),
ADD check (login != ''),
ADD check (email != ''),
ADD primary key (clientID);

ALTER TABLE resetPswrd
ALTER COLUMN resetcode
SET NOT NULL,
ALTER COLUMN token
SET NOT NULL,
ADD FOREIGN KEY (clientID) REFERENCES client (clientID) ON DELETE CASCADE;

ALTER TABLE access
ALTER COLUMN accessStatus
SET NOT NULL,
ALTER COLUMN accessLevel
SET NOT NULL,
ADD check (accessStatus != ''),
ADD check (accessLevel > 0),
ADD primary key (accessID);

ALTER TABLE accessClient
ALTER COLUMN accessID
SET NOT NULL,
ALTER COLUMN clientID
SET NOT NULL,
ADD FOREIGN KEY (accessID) REFERENCES access (accessID) ON DELETE CASCADE,
ADD FOREIGN KEY (clientID) REFERENCES client (clientID) ON DELETE CASCADE;

ALTER TABLE home
ALTER COLUMN name
SET NOT NULL,
ALTER COLUMN ownerID
SET NOT NULL,
ADD check (name != ''),
ADD primary key (homeID);

ALTER TABLE accessHome
ALTER COLUMN accessID
SET NOT NULL,
ALTER COLUMN homeID
SET NOT NULL,
ADD FOREIGN KEY (accessID) REFERENCES access (accessID) ON DELETE CASCADE,
ADD FOREIGN KEY (homeID) REFERENCES home (homeID) ON DELETE CASCADE;

ALTER TABLE device
ALTER COLUMN name
SET NOT NULL,
ALTER COLUMN typeDevice
SET NOT NULL,
ALTER COLUMN status
SET NOT NULL,
ALTER COLUMN brand
SET NOT NULL,
ALTER COLUMN maxParametr
SET NOT NULL,
ALTER COLUMN minParametr
SET NOT NULL,
ALTER COLUMN powerConsumption
SET NOT NULL,
ADD check (name != ''),
ADD check (typeDevice != ''),
ADD check (status != ''),
ADD check (minParametr > 0),
ADD check (maxParametr > minParametr),
ADD primary key (deviceID);

ALTER TABLE deviceHome
ALTER COLUMN deviceID
SET NOT NULL,
ALTER COLUMN homeID
SET NOT NULL,
ADD FOREIGN KEY (deviceID) REFERENCES device (deviceID) ON DELETE CASCADE,
ADD FOREIGN KEY (homeID) REFERENCES home (homeID) ON DELETE CASCADE;

ALTER TABLE historyDev
ALTER COLUMN timeWork
SET NOT NULL,
ALTER COLUMN AverageIndicator
SET NOT NULL,
ALTER COLUMN EnergyConsumed
SET NOT NULL,
ADD check (AverageIndicator > 0),
ADD check (EnergyConsumed > 0),
ADD primary key (historyDevID);

ALTER TABLE historyDevice
ALTER COLUMN deviceID
SET NOT NULL,
ALTER COLUMN historyDevID
SET NOT NULL,
ADD FOREIGN KEY (deviceID) REFERENCES device (deviceID) ON DELETE CASCADE,
ADD FOREIGN KEY (historyDevID) REFERENCES historydev (historyDevID) ON DELETE CASCADE;
