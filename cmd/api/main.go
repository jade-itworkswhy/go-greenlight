package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

// todo: will be generated at build time
const version = "1.0.0"

// todo: read config when the app starts
type config struct {
	port int
	env  string
}

// todo: add handlers, helpers, middlewares
type application struct {
	config config
	logger *slog.Logger
}

func main() {
	var cfg config

	// read values
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// init logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// init app
	app := &application{
		config: cfg,
		logger: logger,
	}

	// init server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	// start server
	logger.Info("starting server", "addr", srv.Addr, "env", cfg.env)

	err := srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
