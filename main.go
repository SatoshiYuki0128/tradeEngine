package main

import (
	"fmt"
	"net/http"
	"tradeEngine/router"

	"github.com/spf13/viper"
)

func main() {
	initConfig()
	routers := router.NewRouter()
	http.ListenAndServe(":80", routers)
}

func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
