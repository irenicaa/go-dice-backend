package main

import (
	"flag"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/irenicaa/go-dice-backend/handlers"
	httputils "github.com/irenicaa/go-dice-backend/http-utils"
	"github.com/irenicaa/go-dice-backend/models"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	flag.Parse()

	stats := models.NewRollStats()
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	http.Handle("/dice", httputils.LoggingMiddleware(
		handlers.DiceHandler{Stats: stats, Logger: logger},
		logger,
		time.Now,
	))
	http.Handle("/stats", httputils.LoggingMiddleware(
		handlers.StatsHandler{Stats: stats, Logger: logger},
		logger,
		time.Now,
	))

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		logger.Fatal(err)
	}
}
