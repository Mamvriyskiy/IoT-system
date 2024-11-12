package benchmark

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"time"
	"github.com/pressly/goose"
	// "io"
	// "os"
	"gorm.io/gorm"
	"database/sql"
	"gorm.io/driver/postgres"
	"github.com/google/uuid"
	"math/rand"
	"testing"
)

type Home struct {
    HomeID    string  `gorm:"column:homeid;primaryKey"`
    Longitude float64
    Latitude  float64
    Name      string
}

func migrationsDataBaseUpGorm(connDB *gorm.DB) error {
	// Проверка соединения с базой данных
	sqlDB, err := connDB.DB()
	if err != nil {
		return fmt.Errorf("не удалось получить sql.DB: %w", err)
	}
	
	// Проверка подключения к базе данных
	if err = sqlDB.Ping(); err != nil {
		return fmt.Errorf("ошибка подключения к базе данных: %w", err)
	}

	// Выполнение миграций с использованием goose
	if err = goose.Up(sqlDB, "./migrations/defaultMigrationsSQL/"); err != nil {
		return fmt.Errorf("ошибка применения миграций: %w", err)
	}

	return nil
}

func (Home) TableName() string {
    return "home"
}

func getHomeGorm(db *gorm.DB, homeID uuid.UUID) (error) {
    var id string

    err := db.Model(&Home{}).Select("homeid").Where("homeid = ?", homeID).Scan(&id).Error
    if err != nil {
        if err == sql.ErrNoRows {
        	return nil
        }
        return err
    }

    return nil
}

func getHomeGormTest(db *gorm.DB) func(b *testing.B) {
	return func(b *testing.B) {
		b.N = N
		for i := 0; i < N; i++ {
			id := uuid.New()
			err := getHomeGorm(db, id)
			if err != nil {
				panic(err)
			}
		}
	}
}

func addHomeGorm(db *gorm.DB, longitude, latitude float64, name string) error {
    home := Home{
        HomeID:    uuid.New().String(),
        Longitude: longitude,
        Latitude:  latitude,
        Name:      name,
    }

    if err := db.Create(&home).Error; err != nil {
        return err
    }

    return nil
}

func addHomeGormTest(db *gorm.DB) func(b *testing.B) {
	return func(b *testing.B) {
		b.N = N
		for i := 0; i < N; i++ {
			rand.Seed(time.Now().UnixNano())
			a, b := generateGeographCoords(7)
			err := addHomeGorm(db, a, b, generateName(7))
			if err != nil {
				panic(err)
			}
		}
	}
}

func SetupTestDatabaseGorm() (testcontainers.Container, *gorm.DB, error) {
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

    host, err := dbContainer.Host(context.Background())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get container host: %w", err)
	}

	port, err := dbContainer.MappedPort(context.Background(), "5432")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get container mapped port: %w", err)
	}

    dsn := fmt.Sprintf("host=%s port=%s user=postgres password=postgres dbname=testdb sslmode=disable", host, port.Port())
	
    connDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, nil, fmt.Errorf("gorm open: %w", err)
    }

	// logReader, _ := dbContainer.Logs(context.Background())
	// io.Copy(os.Stdout, logReader)


    return dbContainer, connDB, nil
}
