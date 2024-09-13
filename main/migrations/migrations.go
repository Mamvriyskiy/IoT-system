package migrations

import (
	"os"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"
)

func MigrationsTestDataBase(connDB *sqlx.DB, dataFile string) error {
	err := connDB.Ping()
	if err != nil {
		return fmt.Errorf("get db: %w", err)
	}

	if err = goose.Up(connDB.DB, "../migrations/testMigrationsSQL/"); err != nil {
		return fmt.Errorf("up migrations: %w", err)
	}

	text, err := os.ReadFile(dataFile)
	if err != nil {
		return fmt.Errorf("read file: %w", err)
	}

	if _, err := connDB.Exec(string(text)); err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}

func MigrationsDataBaseUp(connDB *sqlx.DB) error {
	err := connDB.Ping()
	if err != nil {
		return fmt.Errorf("get db: %w", err)
	}

	if err = goose.Up(connDB.DB, "./migrations/defaultMigrationsSQL/"); err != nil {
		return fmt.Errorf("up migrations: %w", err)
	}

	return nil
}


func MigrationsDataBaseDown(connDB *sqlx.DB) error {
	err := connDB.Ping()
	if err != nil {
		return fmt.Errorf("get db: %w", err)
	}

	if err = goose.Down(connDB.DB, "./migrations/defaultMigrationsSQL/"); err != nil {
		return fmt.Errorf("up migrations: %w", err)
	}

	return nil
}
