package main

import (
	"os"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"
)

func MigrationsResearchDataBase(connDB *sqlx.DB, dataFile string) error {
	err := connDB.Ping()
	if err != nil {
		return fmt.Errorf("get db: %w", err)
	}

	if err = goose.Up(connDB.DB, "./migration"); err != nil {
		return fmt.Errorf("up migrations: %w", err)
	}
	
	requestText := fmt.Sprintf("COPY %s FROM client (FORMAT csv)", dataFile)
	if _, err := connDB.Exec(string(text)); err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}
