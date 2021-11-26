package handlers

import (
	"net/http"

	"github.com/irenicaa/go-dice-backend/generator"
	httputils "github.com/irenicaa/go-dice-backend/http-utils"
	"github.com/irenicaa/go-dice-backend/models"
)

// StatsRegister ...
type StatsRegister interface {
	RegisterDice(dice models.Dice)
}

// DiceHandler ...
type DiceHandler struct {
	Stats  StatsRegister
	Logger httputils.Logger
}

// ServeHTTP ...
//   @router /dice [GET]
//   @summary generate dice rolls
//   @param tries query integer true "amount of roll tries" minimum(1) maximum(100)
//   @param faces query integer true "amount of dice faces" minimum(2) maximum(100)
//   @produce json
//   @success 200 {object} models.RollResults
//   @failure 400 {string} string
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
	diceHandler.Stats.RegisterDice(dice)

	values := generator.GenerateDice(dice)
	results := models.NewRollResults(values)
	httputils.HandleJSON(writer, diceHandler.Logger, results)
}
