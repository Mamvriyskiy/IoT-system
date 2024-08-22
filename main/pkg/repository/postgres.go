package repository

import (
	"fmt"

	"github.com/Mamvriyskiy/DBCourse/main/logger"
	"github.com/jmoiron/sqlx"

	// Импорт драйвера PostgreSQL для его регистрации.
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg *Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		logger.Log("Error", " sqlx.Open", "Error connect DB:", err, "postgres", fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password='' sslmode=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.SSLMode))
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		logger.Log("Error", "Ping()", "Error check connection:", err, "")
		logger.Log("Error", " sqlx.Open", "Error connect DB:", err, "postgres", fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password='' sslmode=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.DBName, cfg.SSLMode))
		return nil, err
	}

	return db, nil
}
