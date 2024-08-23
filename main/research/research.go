package main

import (
	"context"
	//"os"
	//"github.com/jmoiron/sqlx"
	"github.com/Mamvriyskiy/database_course/main/research/migration"
	"github.com/Mamvriyskiy/database_course/main/containers"
)

const dataFile = "./data/researchdata.sql"

func main() {
	dbTestContainers, connDB, err := containers.SetupTestDataBase()

	if err != nil {
		panic(err)
	}
	defer dbTestContainers.Terminate(context.Background())

	err = migration.MigrationsResearchDataBase(connDB, dataFile)
	if err != nil {
		panic(err)
	}

	createRandomData()
}


