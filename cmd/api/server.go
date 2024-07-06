package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (app *application) serve() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(app.logger.Handler(), slog.LevelError),
	}

	shutdownError := make(chan error)
	// start a background goroutine
	go func() {
		// quit channel(os.Signal)
		// this is buffered channel
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		// read the signal from the quit channel
		// This code will block until a signal is received
		s := <-quit

		// log when caught
		app.logger.Info("caught signal", "signal", s.String())

	}()

	// Likewise log a "starting server" message.
	app.logger.Info("starting server", "addr", srv.Addr, "env", app.config.env)

	// Start the server as normal, returning any error.
	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	app.logger.Info("stopped server", "addr", srv.Addr)

	return nil
}
