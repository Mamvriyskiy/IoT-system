package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"
)

func MigrationsResearchDataBaseUp(connDB *sqlx.DB, dataFile, migrationFile string) error {
	err := connDB.Ping()
	if err != nil {
		return fmt.Errorf("get db: %w", err)
	}

	if err = goose.Up(connDB.DB, migrationFile); err != nil {
		return fmt.Errorf("up migrations: %w", err)
	}

	requestText := fmt.Sprintf("COPY client(password,login,email) FROM '%s' DELIMITER ',' CSV HEADER;", dataFile)
	if _, err := connDB.Exec(string(requestText)); err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}

func MigrationsResearchDataBaseDown(connDB *sqlx.DB) error {
	err := connDB.Ping()
	if err != nil {
		return fmt.Errorf("get db: %w", err)
	}

	if err = goose.Down(connDB.DB, "./migrations/indx"); err != nil {
		return fmt.Errorf("down migrations: %w", err)
	}
	
	if err = goose.Down(connDB.DB, "./migrations"); err != nil {
		return fmt.Errorf("down migrations: %w", err)
	}

	return nil
}
