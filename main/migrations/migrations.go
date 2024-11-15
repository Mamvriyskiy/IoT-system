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

	if err = goose.Up(connDB.DB, "../../migrations/testMigrationsSQL/"); err != nil {
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

func isPrimaryDB(connDB *sqlx.DB) (bool, error) {
    var inRecovery bool
    err := connDB.QueryRow("SELECT pg_is_in_recovery()").Scan(&inRecovery)
    if err != nil {
        return false, fmt.Errorf("checking recovery status: %w", err)
    }
    return !inRecovery, nil // true, если это основной сервер
}


func MigrationsDataBaseUp(connDB *sqlx.DB) error {
    err := connDB.Ping()
    if err != nil {
        return fmt.Errorf("get db: %w", err)
    }

    // Проверка, является ли база данных основной
    isPrimary, err := isPrimaryDB(connDB)
    if err != nil {
        return err
    }

    if !isPrimary {
        fmt.Println("Database is in recovery mode (replica). Skipping migrations.")
        return nil // Завершить с успешным кодом
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
