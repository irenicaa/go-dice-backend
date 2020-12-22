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
	httputils "github.com/irenicaa/go-dice-generator/http-utils"
	"github.com/irenicaa/go-dice-generator/models"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	port := flag.Int("port", 8080, "")
	flag.Parse()

	stats := models.NewRollStats()

	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	http.Handle("/dice", handlers.DiceHandler{Stats: stats, Logger: logger})

	http.HandleFunc("/stats", func(writer http.ResponseWriter, request *http.Request) {
		logger.Print("received a request at " + request.URL.String())

		statsCopy := stats.CopyData()
		httputils.HandleJSON(writer, logger, statsCopy)
	})

	address := ":" + strconv.Itoa(*port)
	if err := http.ListenAndServe(address, nil); err != nil {
		logger.Fatal(err)
	}
}
