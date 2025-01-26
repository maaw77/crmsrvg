package crm

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/gorilla/mux"

	"github.com/maaw77/crmsrvg/config"
	"github.com/maaw77/crmsrvg/internal/handlers/v1"
)

// Run runs the server with the specified parameters.
func Run(pathConfig string) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println(pathConfig)

	r := mux.NewRouter()
	// Add your routes as needed
	// r.HandleFunc("/", HelloHandeler)
	handlers.RegHanlders(r)

	srv, wait := config.InitConfigServer(pathConfig)
	srv.Handler = r
	log.Printf("srv = %v", srv)
	log.Printf("wait = %s\n", wait)

	log.Printf("Starting the server with the TCP address = %s\n", srv.Addr)
	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("The server is shutting down")
	// os.Exit(0)
}
