//go:generate go run github.com/go-swagger/go-swagger/cmd/swagger@latest generate spec --scan-models -o /home/maaw/work/Golang/crmsrvg/docs/swagger.yaml
//go:generate go run github.com/go-swagger/go-swagger/cmd/swagger@latest validate /home/maaw/work/Golang/crmsrvg/docs/swagger.yaml

package crm

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	"github.com/maaw77/crmsrvg/config"
	"github.com/maaw77/crmsrvg/internal/database"
	"github.com/maaw77/crmsrvg/internal/handlers/v1"
)

// Run runs the server with the specified parameters.
func Run(pathConfig string) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// log.Println(pathConfig)

	// Connection to the database.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	crmDB, err := database.NewCrmDatabase(ctx, config.InitConnString(""))
	if err != nil {
		log.Fatalln(err)
	}
	if err := crmDB.DBpool.Ping(ctx); err != nil {
		log.Fatal(err)
	}
	defer crmDB.DBpool.Close()

	// Creating a router and registering handlers.
	r := mux.NewRouter()
	handlers.RegGsmHanlders(r)

	srv, wait := config.InitConfigServer(pathConfig)
	srv.Handler = r
	// log.Printf("srv = %v", srv)
	// log.Printf("wait = %s\n", wait)

	// Starting the server.
	log.Printf("starting the server with the TCP address = %s\n", srv.Addr)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx1, cancel1 := context.WithTimeout(context.Background(), wait)
	defer cancel1()
	srv.Shutdown(ctx1)

	log.Println("the server is shutting down")
	// os.Exit(0)
}
