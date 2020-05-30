package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/k-masashi/try-go-clean-arch/app/infrastructure/http/echo"
	"github.com/spf13/viper"
)

func main() {
	setConfig()
	echo.Run()
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
