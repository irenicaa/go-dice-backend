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
	"github.com/irenicaa/go-dice-backend/generator"
	"github.com/irenicaa/go-http-utils/middlewares"
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
	handler := middlewares.LoggingMiddleware(
		handlers.Router{
			BaseURL: "/api/v1",
			DiceHandler: handlers.DiceHandler{
				Stats:         stats,
				DiceGenerator: generator.GenerateDice,
				Logger:        logger,
			},
			StatsHandler: handlers.StatsHandler{
				Stats:  stats,
				Logger: logger,
			},
			Logger: logger,
		},
		logger,
		time.Now,
	)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		logger.Fatal(err)
	}
}
