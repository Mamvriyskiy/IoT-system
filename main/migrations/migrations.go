package migrations

import (
	"github.com/pressly/goose"
)

func MigrationsDataBase() {
	sqlDB, err := pureDB.DB()
	if err != nil {
		return nil, nil, fmt.Errorf("get db: %w", err)
	}

	if err = goose.Up(sqlDB, "../../../deployments/migrations/test_migrations"); err != nil {
		return nil, nil, fmt.Errorf("up migrations: %w", err)
	}

	text, err := os.ReadFile("../../containers/data.sql")
	if err != nil {
		return nil, nil, fmt.Errorf("read file: %w", err)
	}

	if err := pureDB.Exec(string(text)).Error; err != nil {
		return nil, nil, fmt.Errorf("exec: %w", err)
	}
}

