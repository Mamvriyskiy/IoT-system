package testsdatabase

import (
	"testing"
	"context"
	"os"
	"github.com/jmoiron/sqlx"
	"github.com/Mamvriyskiy/database_course/main/migrations"
	"github.com/Mamvriyskiy/database_course/main/containers"
)

var connDB *sqlx.DB

func TestMain(m *testing.M) {
	dbTestContainers, db, err := containers.SetupTestDataBase()

	if err != nil {
		panic(err)
	}
	defer dbTestContainers.Terminate(context.Background())

	connDB = db
	err = migrations.MigrationsTestDataBase(connDB, "./data/data.sql")
	if err != nil {
		panic(err)
	}

	code := m.Run()

	os.Exit(code)
}
