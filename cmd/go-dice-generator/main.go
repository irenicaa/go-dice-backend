package main

import (
	"flag"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/irenicaa/go-dice-generator/handlers"
	"github.com/irenicaa/go-dice-generator/middlewares"
	"github.com/irenicaa/go-dice-generator/models"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	port := flag.Int("port", 8080, "")
	flag.Parse()

	stats := models.NewRollStats()
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	http.Handle("/dice", middlewares.LoggingMiddleware(
		handlers.DiceHandler{Stats: stats, Logger: logger},
		logger,
	))
	http.Handle("/stats", middlewares.LoggingMiddleware(
		handlers.StatsHandler{Stats: stats, Logger: logger},
		logger,
	))

	address := ":" + strconv.Itoa(*port)
	if err := http.ListenAndServe(address, nil); err != nil {
		logger.Fatal(err)
	}
}
