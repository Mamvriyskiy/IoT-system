package repositoryCH

import (
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	//"github.com/jmoiron/sqlx"
	//"github.com/Mamvriyskiy/DBCourse/main/logger"
	//"github.com/ClickHouse/clickhouse-go"
	// Импорт драйвера PostgreSQL для его регистрации.
	_ "github.com/lib/pq"
	//"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	//"context"
	"database/sql"

)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewClickHouseDB(cfg *Config) (*sql.DB, error) {
	var (
        //ctx       = context.Background()
		conn = clickhouse.OpenDB(&clickhouse.Options{
			Addr: []string{"127.0.0.1:8123"},
			Auth: clickhouse.Auth{
				Database: "default",
				Username: "default",
				Password: "",
			},
			Protocol:  clickhouse.HTTP,
		})
    )

	var err error
    if err != nil {
		fmt.Println("1")
        return nil, err
    }

    if err := conn.Ping(); err != nil {
        if exception, ok := err.(*clickhouse.Exception); ok {
            fmt.Printf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
        }
		fmt.Println("2", err)
        // return nil, err
    }
	fmt.Println("3", conn, err)
    return conn, nil
}
