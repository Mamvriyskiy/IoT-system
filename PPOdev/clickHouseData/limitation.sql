ALTER TABLE client MODIFY COLUMN login String NOT NULL;
ALTER TABLE client MODIFY COLUMN email String NOT NULL;


ALTER TABLE client MODIFY COLUMN login String NOT NULL;
ALTER TABLE client MODIFY COLUMN email String NOT NULL;

ALTER TABLE client ADD CONSTRAINT check_login CHECK (login != '');
ALTER TABLE client ADD CONSTRAINT check_email CHECK (email != '');
