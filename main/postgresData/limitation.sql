ALTER TABLE client
    ALTER COLUMN password SET NOT NULL,
    ALTER COLUMN login SET NOT NULL,
    ALTER COLUMN email SET NOT NULL;

ALTER TABLE client
    ADD CHECK (password != ''),
    ADD CHECK (login != ''),
    ADD CHECK (email != ''),
    ADD PRIMARY KEY (clientID);

ALTER TABLE resetPswrd
    ALTER COLUMN resetcode SET NOT NULL,
    ALTER COLUMN token SET NOT NULL;

ALTER TABLE resetPswrd 
    ADD FOREIGN KEY (clientID) REFERENCES client (clientID) ON DELETE CASCADE;

ALTER TABLE access
    ALTER COLUMN accessStatus SET NOT NULL,
    ALTER COLUMN accessLevel SET NOT NULL;

ALTER TABLE access
    ADD CHECK (accessStatus != ''),
    ADD CHECK (accessLevel > 0),
    ADD PRIMARY KEY (accessID);

ALTER TABLE accessClient
    ALTER COLUMN accessID SET NOT NULL,
    ALTER COLUMN clientID SET NOT NULL;

ALTER TABLE accessClient 
    ADD FOREIGN KEY (accessID) REFERENCES access (accessID) ON DELETE CASCADE,
    ADD FOREIGN KEY (clientID) REFERENCES client (clientID) ON DELETE CASCADE;

ALTER TABLE home
    ALTER COLUMN name SET NOT NULL,
    ALTER COLUMN ownerID SET NOT NULL;

ALTER TABLE home 
    ADD CHECK (name != ''),
    ADD PRIMARY KEY (homeID);

ALTER TABLE accessHome
    ALTER COLUMN accessID SET NOT NULL,
    ALTER COLUMN homeID SET NOT NULL;

ALTER TABLE accessHome 
    ADD FOREIGN KEY (accessID) REFERENCES access (accessID) ON DELETE CASCADE,
    ADD FOREIGN KEY (homeID) REFERENCES home (homeID) ON DELETE CASCADE;

ALTER TABLE device
    ALTER COLUMN name SET NOT NULL,
    ALTER COLUMN typeDevice SET NOT NULL,
    ALTER COLUMN status SET NOT NULL,
    ALTER COLUMN brand SET NOT NULL;
 
ALTER TABLE device 
    ADD CHECK (name != ''),
    ADD CHECK (typeDevice != ''),
    ADD CHECK (status != ''),
    ADD PRIMARY KEY (deviceID);

ALTER TABLE deviceHome
    ALTER COLUMN deviceID SET NOT NULL,
    ALTER COLUMN homeID SET NOT NULL;

ALTER TABLE deviceHome 
    ADD FOREIGN KEY (deviceID) REFERENCES device (deviceID) ON DELETE CASCADE,
    ADD FOREIGN KEY (homeID) REFERENCES home (homeID) ON DELETE CASCADE;

ALTER TABLE historyDev
    ALTER COLUMN timeWork SET NOT NULL,
    ALTER COLUMN AverageIndicator SET NOT NULL,
    ALTER COLUMN EnergyConsumed SET NOT NULL;

ALTER TABLE historyDev 
    ADD CHECK (AverageIndicator > 0),
    ADD CHECK (EnergyConsumed > 0),
    ADD PRIMARY KEY (historyDevID);

ALTER TABLE historyDevice
    ALTER COLUMN deviceID SET NOT NULL,
    ALTER COLUMN historyDevID SET NOT NULL;

ALTER TABLE historyDevice 
    ADD FOREIGN KEY (deviceID) REFERENCES device (deviceID) ON DELETE CASCADE,
    ADD FOREIGN KEY (historyDevID) REFERENCES historydev (historyDevID) ON DELETE CASCADE;

ALTER TABLE TypeCharacter
    ALTER COLUMN typecharacter SET NOT NULL,
    ALTER COLUMN unitmeasure SET NOT NULL;

ALTER TABLE TypeCharacter 
    ADD CHECK (typecharacter != ''),
    ADD CHECK (unitmeasure != ''),
    ADD PRIMARY KEY (typecharacterid);
ALTER TABLE DeviceCharacteristics
    ALTER COLUMN deviceID SET NOT NULL,
    ALTER COLUMN valuesChar SET NOT NULL,
    ALTER COLUMN typecharacterid SET NOT NULL;

ALTER TABLE DeviceCharacteristics 
    ADD CHECK (valuesChar > 0),
    ADD PRIMARY KEY (characterid),
    ADD FOREIGN KEY (deviceID) REFERENCES device (deviceID) ON DELETE CASCADE,
    ADD FOREIGN KEY (typecharacterid) REFERENCES TypeCharacter (typecharacterid) ON DELETE CASCADE;
