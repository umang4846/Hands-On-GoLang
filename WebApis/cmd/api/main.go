package api

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

//The application version number
const version = "1.0.0"

//The configuration settings
type config struct {
	port int
	env  string //development, staging, production
}

//Dependency Injection
type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config
	//Read in the flags that are needed to populate our cfg
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development | staging | production)")
	flag.Parse()
	//create a logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	//Create an instance of our application struct
	app := &application{
		config: cfg,
		logger: logger,
	}
	//create our mux servemux
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)
	//create our HTTP server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	//Start our server
	logger.Printf("starting %s on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
