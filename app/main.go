package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

func main() {
	setConfig()

	// DB Settings
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	dbConnection, error := sql.Open("mysql", connectionString)

	if error != nil {
		log.Fatal(error)
	}

	error = dbConnection.Ping()
	if error != nil {
		log.Fatal(error)
	}

	defer func() {
		error = dbConnection.Close()
		if error != nil {
			log.Fatal(error)
		}
	}()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Start(viper.GetString("server.address"))
}

func setConfig() {
	viper.SetConfigFile(`config.json`)
	error := viper.ReadInConfig()

	if error != nil {
		log.Println("Failed to read config")
		panic(error)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service is Running on Debug mode !")
	}
}
