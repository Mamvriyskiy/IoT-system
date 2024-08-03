package main

import (
	"os"

	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/logger"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/handler"
	//"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/repository"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/repositoryCH"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		logger.Log("Error", "initCongig", "Error config DB:", err, "")
		return
	}
	logger.Log("Info", "", "InitConfig", nil)

	if err := godotenv.Load(); err != nil {
		logger.Log("Error", "Load", "Load env file:", err, "")
		return
	}
	logger.Log("Info", "", "Load env", nil)

	// db, err := repository.NewPostgresDB(&repository.Config{
	// 	Host:     viper.GetString("db.host"),
	// 	Port:     viper.GetString("db.port"),
	// 	Username: viper.GetString("db.username"),
	// 	Password: os.Getenv("DB_PASSWORD"),
	// 	DBName:   viper.GetString("db.dbname"),
	// 	SSLMode:  viper.GetString("db.sslmode"),
	// })

	db, err := repositoryCH.NewClickHouseDB(&repositoryCH.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logger.Log("Error", "initCongig", "Error config DB:", err, "")
		return
	}
	
	//repos := repository.NewRepository(db)

	repos := repositoryCH.NewRepositoryCH(db)
	services := service.NewServicesCH(repos)
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
