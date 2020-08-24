package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rs/zerolog"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	config "github.com/serdarkalayci/configchangergo/configuration"
	"github.com/serdarkalayci/configchangergo/handlers"
)

var bindAddress = ":5200"

func main() {

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	config.SetConfigValues()

	// create the handlers
	apiContext := handlers.NewAPIContext()

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()

	// handlers for API
	getR := sm.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/", apiContext.Index)
	getR.HandleFunc("/health/live", apiContext.Live)
	getR.HandleFunc("/health/ready", apiContext.Live)
	getR.HandleFunc("/log", apiContext.Log)

	// create a new server
	s := http.Server{
		Addr:         bindAddress,       // configure the bind address
		Handler:      sm,                // set the default handler
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		log.Debug().Msgf("Starting server on %s", bindAddress)

		err := s.ListenAndServe()
		if err != nil {
			log.Fatal().
				Err(err).
				Msgf("Cannot start service")
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Debug().Msgf("Got signal %s", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
