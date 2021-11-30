package main

import (
	"flag"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/irenicaa/go-dice-backend/gateways/handlers"
	"github.com/irenicaa/go-dice-backend/gateways/storages"
	httputils "github.com/irenicaa/go-dice-backend/http-utils"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	flag.Parse()

	stats := storages.NewRollStats()
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	http.Handle("/api/v1/dice", httputils.LoggingMiddleware(
		handlers.DiceHandler{Stats: stats, Logger: logger},
		logger,
		time.Now,
	))
	http.Handle("/api/v1/stats", httputils.LoggingMiddleware(
		handlers.StatsHandler{Stats: stats, Logger: logger},
		logger,
		time.Now,
	))

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		logger.Fatal(err)
	}
}
