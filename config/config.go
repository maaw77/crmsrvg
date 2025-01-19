package config

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

// InitCongigServer returns a pointer to http.Server
// with certain parameters for running an HTTP server.
func InitConfigServer() *http.Server {

	// Setting default values for the server.
	viper.SetDefault("server.Addr", "0.0.0.0:8080")
	viper.SetDefault("server.WriteTimeout", 15)
	viper.SetDefault("server.ReadTimeout", 15)
	viper.SetDefault("server.IdleTimeout", 60)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("The configuration file was not found. Default parameters are used.")
		} else {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}
	return &http.Server{
		Addr:         viper.GetString("server.Addr"),
		WriteTimeout: time.Second * viper.GetDuration("server.WriteTimeout"),
		ReadTimeout:  time.Second * viper.GetDuration("server.ReadTimeout"),
		IdleTimeout:  time.Second * viper.GetDuration("server.IdleTimeout"),
	}
}
