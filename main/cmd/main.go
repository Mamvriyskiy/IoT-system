package main

import (
	"os"
	"github.com/Mamvriyskiy/database_course/main/logger"
	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/Mamvriyskiy/database_course/main/pkg/handler"
	"github.com/Mamvriyskiy/database_course/main/pkg/repository"
	"github.com/Mamvriyskiy/database_course/main/pkg/service"
	"github.com/Mamvriyskiy/database_course/main/migrations"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"fmt"
)

func main() {
	fmt.Println("Server start...")
	if err := initConfig(); err != nil {
		logger.Log("Error", "initCongig", "Error config DB:", err, "")
		fmt.Println(err)
		return
	}
	logger.Log("Info", "", "InitConfig", nil)

	if err := godotenv.Load(); err != nil {
		logger.Log("Error", "Load", "Load env file:", err, "")
		fmt.Println(err)
		return
	}
	logger.Log("Info", "", "Load env", nil)

	db, err := repository.NewPostgresDB(&repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		fmt.Println(err)
		logger.Log("Error", "initCongig", "Error config DB:", err, "")
		return
	}

	err = migrations.MigrationsDataBaseUp(db)

	if err != nil {
		fmt.Println(err)
		logger.Log("Error", "MigrationsDataBaseUp", "Error migrations:", err, "")
		return
	}

	// defer func() {
	// 	err = migrations.MigrationsDataBaseDown(db)
	// 	if err != nil {
	// 		logger.Log("Error", "MigrationsDataBaseDown", "Error migrations:", err, "")
	// 	}
	// }()

	repos := repository.NewRepository(db)
	services := service.NewServicesPsql(repos)
	handlers := handler.NewHandler(services)

	logger.Log("Info", "", "The connection to the database is established", nil)

	srv := new(pkg.Server)
	if err := srv.Run("8000", handlers.InitRouters()); err != nil {
		logger.Log("Error", "Run", "Error occurred while running http server:", err, "")
		return
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
