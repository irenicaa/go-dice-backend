package main

import (
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/irenicaa/go-dice-generator/generator"
	httputils "github.com/irenicaa/go-dice-generator/http-utils"
	"github.com/irenicaa/go-dice-generator/models"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	port := flag.Int("port", 8080, "")
	flag.Parse()

	stats := map[string]int{}
	statsMutex := sync.RWMutex{}

	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	http.HandleFunc("/dice", func(writer http.ResponseWriter, request *http.Request) {
		logger.Print("received a request at " + request.URL.String())

		tries, err := httputils.GetIntFormValue(request, "tries", 1, 100)
		if err != nil {
			httputils.HandleError(
				writer,
				logger,
				http.StatusBadRequest,
				"unable to get the tries parameter: %v",
				err,
			)

			return
		}

		faces, err := httputils.GetIntFormValue(request, "faces", 2, 100)
		if err != nil {
			httputils.HandleError(
				writer,
				logger,
				http.StatusBadRequest,
				"unable to get the faces parameter: %v",
				err,
			)

			return
		}

		dice := models.Dice{Tries: tries, Faces: faces}
		statsMutex.Lock()
		stats[dice.String()]++
		statsMutex.Unlock()

		values := generator.GenerateDice(dice)
		results := models.NewRollResults(values)
		resultsBytes, err := json.Marshal(results)
		if err != nil {
			httputils.HandleError(
				writer,
				logger,
				http.StatusInternalServerError,
				"unable to marshal the roll results: %v",
				err,
			)

			return
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.Write(resultsBytes)
	})

	http.HandleFunc("/stats", func(writer http.ResponseWriter, request *http.Request) {
		logger.Print("received a request at " + request.URL.String())

		statsCopy := map[string]int{}
		statsMutex.RLock()
		for dice, count := range stats {
			statsCopy[dice] = count
		}
		statsMutex.RUnlock()

		statsBytes, err := json.Marshal(statsCopy)
		if err != nil {
			httputils.HandleError(
				writer,
				logger,
				http.StatusInternalServerError,
				"unable to marshal the stats: %v",
				err,
			)

			return
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.Write(statsBytes)
	})

	address := ":" + strconv.Itoa(*port)
	if err := http.ListenAndServe(address, nil); err != nil {
		logger.Fatal(err)
	}
}
