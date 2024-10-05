package containers

import (
	"context"
	"fmt"
	//"github.com/pressly/goose"
	"github.com/jmoiron/sqlx"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	//"gorm.io/driver/postgres"
	//"gorm.io/gorm"
	// "github.com/Mamvriyskiy/database_course/main/migrations"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"

)

func SetupTestDataBase() (testcontainers.Container, *sqlx.DB, error) {
	containerReq := testcontainers.ContainerRequest{
		Image:        "postgres:15.4-alpine",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
		Env: map[string]string{
			"POSTGRES_DB":       "testdb",
			"POSTGRES_PASSWORD": "postgres",
			"POSTGRES_USER":     "postgres",
		},
		Binds: []string{
			"/Users/ivanmamvriyskiy/Desktop/DBCourse/main/research/mnt:/mnt",
		},
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

    // Получение порта контейнера
    port, _ := dbContainer.MappedPort(context.Background(), "5432")

	dsn := fmt.Sprintf("host=%s port=%d user=postgres password=postgres dbname=testdb sslmode=disable", host, port.Int())
	connDB, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, nil, fmt.Errorf("gorm open: %w", err)
	}

	return dbContainer, connDB, nil
}
