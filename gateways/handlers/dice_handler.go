package handlers

import (
	"net/http"

	httputils "github.com/irenicaa/go-dice-backend/http-utils"
	"github.com/irenicaa/go-dice-backend/models"
)

// StatsRegister ...
type StatsRegister interface {
	RegisterDice(dice models.Dice) error
}

// DiceGenerator ...
type DiceGenerator func(dice models.Dice) ([]int, error)

// DiceHandler ...
type DiceHandler struct {
	Stats         StatsRegister
	DiceGenerator DiceGenerator
	Logger        httputils.Logger
}

// ServeHTTP ...
//   @router /dice [GET]
//   @summary generate dice rolls
//   @param tries query integer true "amount of roll tries" minimum(1) maximum(100)
//   @param faces query integer true "amount of dice faces" minimum(2) maximum(100)
//   @produce json
//   @success 200 {object} models.RollResults
//   @failure 400 {string} string
//   @failure 500 {string} string
func (diceHandler DiceHandler) ServeHTTP(
	writer http.ResponseWriter,
	request *http.Request,
) {
	tries, err := httputils.GetIntFormValue(request, "tries", 1, 100)
	if err != nil {
		status, message :=
			http.StatusBadRequest, "unable to get the tries parameter: %v"
		httputils.HandleError(writer, diceHandler.Logger, status, message, err)

		return
	}

	faces, err := httputils.GetIntFormValue(request, "faces", 2, 100)
	if err != nil {
		status, message :=
			http.StatusBadRequest, "unable to get the faces parameter: %v"
		httputils.HandleError(writer, diceHandler.Logger, status, message, err)

		return
	}

	dice := models.Dice{Tries: tries, Faces: faces}
	if err := diceHandler.Stats.RegisterDice(dice); err != nil {
		status, message :=
			http.StatusInternalServerError, "unable to register the dice: %v"
		httputils.HandleError(writer, diceHandler.Logger, status, message, err)

		return
	}

	values, err := diceHandler.DiceGenerator(dice)
	if err != nil {
		status, message :=
			http.StatusInternalServerError, "unable to generate dice rolls: %v"
		httputils.HandleError(writer, diceHandler.Logger, status, message, err)

		return
	}

	results := models.NewRollResults(values)
	httputils.HandleJSON(writer, diceHandler.Logger, results)
}
