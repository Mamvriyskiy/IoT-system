package benchmark

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"time"
	//"os"
	"database/sql"
	"github.com/google/uuid"
	"math/rand"
	"testing"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
)

func getHomeSqlx(db *sqlx.DB, homeID uuid.UUID) error {
	var id string
	query := fmt.Sprintf("SELECT homeid from %s where homeid = $1", "home")
	err := db.Get(&id, query, homeID)
	if err != nil {
        if err == sql.ErrNoRows {
            return nil
        }

        panic(err)
    }

	return err
}

func getHomeSqlxTest(db *sqlx.DB) func(b *testing.B) {
	return func(b *testing.B) {
		b.N = N
		for i := 0; i < N; i++ {
			id := uuid.New()
			err := getHomeSqlx(db, id)
			if err != nil {
				panic(err)
			}
		}
	}
}

func addHomeSqlx(db *sqlx.DB, longitude, latitude float64, name string) error {
	id := uuid.New()
	var homeID string
	query := fmt.Sprintf("INSERT INTO %s (longitude, latitude, name, homeID) values ($1, $2, $3, $4) RETURNING homeid", "home")
	row := db.QueryRow(query, longitude, latitude, name, id)
	var err error
	if err := row.Scan(&homeID); err != nil {
		panic(err)
	}

	_ = homeID

	return err
}

func addHomeSqlxTest(db *sqlx.DB) func(b *testing.B) {
	return func(b *testing.B) {
		b.N = N
		for i := 0; i < N; i++ {
			rand.Seed(time.Now().UnixNano())
			a, b := generateGeographCoords(7)
			err := addHomeSqlx(db, a, b, generateName(7))
			if err != nil {
				panic(err)
			}
		}
	}
}

func SetupTestDatabaseSqlx() (testcontainers.Container, *sqlx.DB, error) {
	containerReq := testcontainers.ContainerRequest{
        Image:        "postgres:16",
        ExposedPorts: []string{"5432/tcp"}, 
		WaitingFor: wait.ForListeningPort("5432/tcp"),
		//.WithStartupTimeout(3 * time.Minute),  // Увеличено время ожидания
        Env: map[string]string{
            "POSTGRES_DB":       "testdb",
            "POSTGRES_PASSWORD": "postgres",
            "POSTGRES_USER":     "postgres",
        },
		//Networks:     []string{"bench"},
    }

    dbContainer, err := testcontainers.GenericContainer(
        context.Background(),
        testcontainers.GenericContainerRequest{
            ContainerRequest: containerReq,
            Started:          true,
        })
    if err != nil {
        return nil, nil, fmt.Errorf("failed to start container: %w", err)
    }

    host, _ := dbContainer.Host(context.Background())
    port, _ := dbContainer.MappedPort(context.Background(), "5432")  // Правильный порт

    dsn := fmt.Sprintf("host=%s port=%d user=postgres password=postgres dbname=testdb sslmode=disable", host, port.Int())
    connDB, err := sqlx.Open("postgres", dsn)
    if err != nil {
        return nil, nil, fmt.Errorf("gorm open: %w", err)
    }

    return dbContainer, connDB, nil
}
