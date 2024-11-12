package benchmark

import (
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"github.com/Mamvriyskiy/database_course/main/migrations"
	"crypto/rand"
	"math/big"
	"testing"
)

const (
	N = 1
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func ClientBench() []string {

	// ========================== GORM ==========================

	dbContainerGorm, dbGorm, errGorm := SetupTestDatabaseGorm()
	defer func(dbContainerGorm testcontainers.Container, ctx context.Context) {
		errGorm := dbContainerGorm.Terminate(ctx)
		if errGorm != nil {
			return
		}
	}(dbContainerGorm, context.Background())

	if errGorm != nil {
		fmt.Println(errGorm)
		return nil
	}

	err := migrationsDataBaseUpGorm(dbGorm)
	if err != nil {
		panic(err)
	}

	var res []string
	addHome := addHomeGormTest(dbGorm)
	resultsAddHome := testing.Benchmark(addHome)
	// res = append(res, fmt.Sprintf("gorm.AddHome -- runs %5d times\tCPU: %5d ns/op\tRAM: %5d allocs/op %5d bytes/op\n",
	// resultsAddHome.N, resultsAddHome.NsPerOp(), resultsAddHome.AllocsPerOp(), resultsAddHome.AllocedBytesPerOp()))
	res = append(res, fmt.Sprintf("%5d %5d %5d",
		resultsAddHome.NsPerOp(), resultsAddHome.AllocsPerOp(), resultsAddHome.AllocedBytesPerOp()))

	getHome := getHomeGormTest(dbGorm)
	resultsGetHome := testing.Benchmark(getHome)
	// res = append(res, fmt.Sprintf("gorm.GetHome -- runs %5d times\tCPU: %5d ns/op\tRAM: %5d allocs/op %5d bytes/op\n",
	// resultsGetHome.N, resultsGetHome.NsPerOp(),resultsGetHome.AllocsPerOp(), resultsGetHome.AllocedBytesPerOp()))

	res = append(res, fmt.Sprintf("%5d %5d %5d",
		resultsGetHome.NsPerOp(), resultsGetHome.AllocsPerOp(), resultsGetHome.AllocedBytesPerOp()))


	// ========================== SQLX ==========================

	// ======================== Container =======================

	dbContainerSqlx, dbSqlx, errSqlx := SetupTestDatabaseSqlx()
	defer func(dbContainerSqlx testcontainers.Container, ctx context.Context) {
		errSqlx := dbContainerSqlx.Terminate(ctx)
		if errSqlx != nil {
			return
		}
	}(dbContainerSqlx, context.Background())

	if errSqlx != nil {
		fmt.Println(errSqlx)
		return nil
	}

	// ======================== Migration =======================

	err = migrations.MigrationsDataBaseUp(dbSqlx)
	if err != nil {
		panic(err)
	}

	// ======================== BenchTest =======================

	addHomeSqlx := addHomeSqlxTest(dbSqlx)
	resultsAddHomeSqlx := testing.Benchmark(addHomeSqlx)
	// res = append(res, fmt.Sprintf("sqlx.AddHome -- runs %5d times\tCPU: %5d ns/op\tRAM: %5d allocs/op %5d bytes/op\n",
	// resultsAddHomeSqlx.N, resultsAddHomeSqlx.NsPerOp(), resultsAddHomeSqlx.AllocsPerOp(), resultsAddHomeSqlx.AllocedBytesPerOp()))

	res = append(res, fmt.Sprintf("%5d %5d %5d",
		resultsAddHomeSqlx.NsPerOp(), resultsAddHomeSqlx.AllocsPerOp(), resultsAddHomeSqlx.AllocedBytesPerOp()))

	getHomeSqlx := getHomeSqlxTest(dbSqlx)
	resultsGetHomeSqlx := testing.Benchmark(getHomeSqlx)
	// res = append(res, fmt.Sprintf("sqlx.GetHome -- runs %5d times\tCPU: %5d ns/op\tRAM: %5d allocs/op %5d bytes/op\n",
	// resultsGetHomeSqlx.N, resultsGetHomeSqlx.NsPerOp(), resultsGetHomeSqlx.AllocsPerOp(), resultsGetHomeSqlx.AllocedBytesPerOp()))

	res = append(res, fmt.Sprintf("%5d %5d %5d",
		resultsGetHomeSqlx.NsPerOp(), resultsGetHomeSqlx.AllocsPerOp(), resultsGetHomeSqlx.AllocedBytesPerOp()))

	fmt.Println(res)

	return res
}

func generateName(lengthName int) string {
	name := make([]byte, lengthName)
	for j := 0; j < lengthName; j++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		name[j] = charset[n.Int64()]
	}

	return string(name)
}

func generateGeographCoords(lengthGeographCoords int) (float64, float64){
	n, err := rand.Int(rand.Reader, big.NewInt(int64(lengthGeographCoords)))
	if err != nil {
		return 11111111, 11111111
	}

	return float64(n.Int64()), float64(n.Int64())
}
