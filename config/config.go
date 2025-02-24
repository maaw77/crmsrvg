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

// InitConnString returns a connection string for parsing using pgxpool.ParseConfig (github.com/jackc/pgx/v5/pgconn)
// in datebase.NewCrmDatabase.
func InitConnString(pathConfig string) (connString string) {
	dir, file := filepath.Split(pathConfig)
	log.Println(dir, file)
	fileName := strings.Split(file, ".")[0]

	// Setting default values for the connection string.
	viper.SetDefault("db.DB", "postgres")
	viper.SetDefault("db.User", "postgres")
	viper.SetDefault("db.Password", "crmpassword")
	viper.SetDefault("db.Host", "localhost")
	viper.SetDefault("db.Port", "5433")
	viper.SetDefault("db.PoolMaxConns", "10")

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

	connString = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&pool_max_conns=%s",
		viper.GetString("db.User"),
		viper.GetString("db.Password"),
		viper.GetString("db.Host"),
		viper.GetString("db.Port"),
		viper.GetString("db.DB"),
		viper.GetString("db.PoolMaxConns"),
	)
	return connString
}
