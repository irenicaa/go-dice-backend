package main

import (
	"flag"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/irenicaa/go-dice-generator/handlers"
	httputils "github.com/irenicaa/go-dice-generator/http-utils"
	"github.com/irenicaa/go-dice-generator/models"
)

func main() {
	rand.Seed(time.Now().UnixNano())

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

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	address := ":" + port
	if err := http.ListenAndServe(address, nil); err != nil {
		logger.Fatal(err)
	}
}
