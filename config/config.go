package config

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// InitCongigServer returns a pointer to http.Server
// with certain parameters for running an HTTP server and
// a deadline for waiting for the server to shut down.
func InitConfigServer(pathConfig string) (*http.Server, time.Duration) {
	dir, file := filepath.Split(pathConfig)
	log.Println(dir, file)
	fileName := strings.Split(file, ".")[0]

	// Setting default values for the server and the deadline.
	viper.SetDefault("server.Addr", "0.0.0.0:8080")
	viper.SetDefault("server.WriteTimeout", 15)
	viper.SetDefault("server.ReadTimeout", 15)
	viper.SetDefault("server.IdleTimeout", 60)
	viper.SetDefault("server.ShutdownTimeout", 15)

	viper.SetConfigName(fileName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(dir)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("The configuration file was not found. The default parameters are used.")
		} else {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}

	return &http.Server{
			Addr:         viper.GetString("server.Addr"),
			WriteTimeout: time.Second * viper.GetDuration("server.WriteTimeout"),
			ReadTimeout:  time.Second * viper.GetDuration("server.ReadTimeout"),
			IdleTimeout:  time.Second * viper.GetDuration("server.IdleTimeout"),
		},
		time.Second * viper.GetDuration("server.ShutdownTimeout")
}

func InitConnString(pathConfig string) (connString string, err error) {
	dir, file := filepath.Split(pathConfig)
	log.Println(dir, file)
	fileName := strings.Split(file, ".")[0]

	// Setting default values for the server and the deadline.
	viper.SetDefault("server.Addr", "0.0.0.0:8080")
	viper.SetDefault("server.WriteTimeout", 15)
	viper.SetDefault("server.ReadTimeout", 15)
	viper.SetDefault("server.IdleTimeout", 60)
	viper.SetDefault("server.ShutdownTimeout", 15)

	viper.SetConfigName(fileName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(dir)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("The configuration file was not found. The default parameters are used.")
		} else {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}

	return
}
