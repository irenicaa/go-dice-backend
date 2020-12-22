package handlers

import (
	"net/http"

	"github.com/irenicaa/go-dice-generator/generator"
	httputils "github.com/irenicaa/go-dice-generator/http-utils"
	"github.com/irenicaa/go-dice-generator/models"
)

// DiceHandler ...
type DiceHandler struct {
	Stats  models.RollStats
	Logger httputils.Logger
}

func (diceHandler DiceHandler) ServeHTTP(
	writer http.ResponseWriter,
	request *http.Request,
) {
	tries, err := httputils.GetIntFormValue(request, "tries", 1, 100)
	if err != nil {
		httputils.HandleError(
			writer,
			diceHandler.Logger,
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
			diceHandler.Logger,
			http.StatusBadRequest,
			"unable to get the faces parameter: %v",
			err,
		)

		return
	}

	dice := models.Dice{Tries: tries, Faces: faces}
	diceHandler.Stats.Register(dice)

	values := generator.GenerateDice(dice)
	results := models.NewRollResults(values)
	httputils.HandleJSON(writer, diceHandler.Logger, results)
}
