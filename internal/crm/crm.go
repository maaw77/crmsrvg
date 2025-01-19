package crm

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	"github.com/maaw77/crmsrvg/config"
)

func HelloHandeler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>%s</h1>", "Hello Hugo boss!")

}

// Run runs the server with the specified parameters.
func Run() {
	var wait time.Duration
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	flag.DurationVar(&wait, "gt", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	log.Println("Graceful timeout", wait)

	r := mux.NewRouter()
	// Add your routes as needed
	r.HandleFunc("/", HelloHandeler)

	srv := config.InitConfigServer()
	srv.Handler = r
	log.Println(srv)

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

	log.Println("shutdown")
	// os.Exit(0)
}
