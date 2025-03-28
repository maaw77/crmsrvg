//go:build migrate

package crm

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func init() {
	log.Println("the beginning of migrations")
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	// postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${HOST_DB}:${PORT_DB}/${POSTGRES_DB}?sslmode=disable
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("HOST_DB"), os.Getenv("PORT_DB"), os.Getenv("POSTGRES_DB"))
	// dbUrl := "postgres://postgres:crmpassword@localhost:5433/postgres?sslmode=disable"
	m, err := migrate.New(
		"file://./migrations",
		dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Println("migrations: ", err)
	}
}
