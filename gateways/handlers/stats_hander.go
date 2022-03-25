package handlers

import (
	"net/http"

	"github.com/irenicaa/go-dice-backend/v2/models"
	httputils "github.com/irenicaa/go-http-utils"
)

// StatsCopier ...
type StatsCopier interface {
	CopyRollStats() (models.RollStats, error)
}

// StatsHandler ...
type StatsHandler struct {
	Stats  StatsCopier
	Logger httputils.Logger
}

// ServeHTTP ...
//   @router /stats [GET]
//   @summary get stats of dice rolls
//   @produce json
//   @success 200 {object} models.RollStats
//   @failure 500 {string} string
func (statsHandler StatsHandler) ServeHTTP(
	writer http.ResponseWriter,
	request *http.Request,
) {
	statsCopy, err := statsHandler.Stats.CopyRollStats()
	if err != nil {
		status, message :=
			http.StatusInternalServerError, "unable to copy the roll stats: %v"
		httputils.HandleError(writer, statsHandler.Logger, status, message, err)

		return
	}

	httputils.HandleJSON(writer, statsHandler.Logger, statsCopy)
}
