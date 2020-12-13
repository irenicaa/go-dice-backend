package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/irenicaa/go-dice-generator/generator"
	httputils "github.com/irenicaa/go-dice-generator/http-utils"
	"github.com/irenicaa/go-dice-generator/models"
)

func main() {
	http.HandleFunc("/dice", func(writer http.ResponseWriter, request *http.Request) {
		log.Print("received a request at " + request.URL.String())

		tries, err := httputils.GetIntFormValue(request, "tries", 1, 100)
		if err != nil {
			message := fmt.Sprintf("unable to get the tries parameter: %v", err)
			log.Print(message)

			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(message))

			return
		}

		faces, err := httputils.GetIntFormValue(request, "faces", 2, 100)
		if err != nil {
			message := fmt.Sprintf("unable to get the faces parameter: %v", err)
			log.Print(message)

			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(message))

			return
		}

		dice := models.Dice{Tries: tries, Faces: faces}
		values := generator.GenerateDice(dice)
		results := models.NewRollResults(values)
		resultsBytes, err := json.Marshal(results)
		if err != nil {
			message := fmt.Sprintf("unable to marshal the roll results: %v", err)
			log.Print(message)

			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(message))

			return
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.Write(resultsBytes)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
