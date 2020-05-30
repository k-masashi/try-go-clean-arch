package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

func main() {
	setConfig()

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
