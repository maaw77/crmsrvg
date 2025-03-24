package crm

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/maaw77/crmsrvg/config"
	_ "github.com/maaw77/crmsrvg/docs"
	"github.com/maaw77/crmsrvg/internal/database"
	"github.com/maaw77/crmsrvg/internal/handlers/v1"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title        CRM server
// @version         1.0
// @description     Fuel and Lubricants Accounting Service.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Maaw
// @contact.email  maaw@mail.ru

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securitydefinitions.bearerauth BearerAuth

// @query.collection.format multi

// Run runs the server with the specified parameters.
func Run(pathConfig string) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// log.Println(pathConfig)

	// Connection to the database.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	crmDB, err := database.NewCrmDatabase(ctx, config.InitConnString(pathConfig))
	if err != nil {
		log.Fatalln(err)
	}
	if err := crmDB.DBpool.Ping(ctx); err != nil {
		log.Fatal(err)
	}
	defer crmDB.DBpool.Close()

	// Creating a router and registering handlers.
	r := mux.NewRouter()

	r.NotFoundHandler = http.HandlerFunc(handlers.DefaultHandler)
	r.MethodNotAllowedHandler = http.HandlerFunc(handlers.MethodNotAllowed)

	apiR := r.PathPrefix("/api/v1").Subrouter()
	// Маршрут для Swagger UI
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), // URL для документации JSON
	))

	// docR := apiR.Methods(http.MethodGet).Subrouter()
	// opts := middleware.RedocOpts{BasePath: "/api/v1", SpecURL: "/api/v1/docs/swagger.yaml"}
	// sh := middleware.Redoc(opts, nil)
	// docR.Handle("/docs/", sh)
	// docR.Handle("/docs/swagger.yaml", http.StripPrefix("/api/v1", http.FileServer(http.Dir("."))))

	// Uesrs
	handlers.RegUsersHanlders(apiR, crmDB)
	// GSM
	handlers.RegGsmHanlders(apiR, crmDB)

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
