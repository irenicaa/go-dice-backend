package handlers

import (
	"net/http"

	httputils "github.com/irenicaa/go-dice-backend/http-utils"
	"github.com/irenicaa/go-dice-backend/models"
)

// StatsCopier ...
type StatsCopier interface {
	CopyRollStats() models.RollStatsData
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
//   @success 200 {object} models.RollStatsData
func (statsHandler StatsHandler) ServeHTTP(
	writer http.ResponseWriter,
	request *http.Request,
) {
	statsCopy := statsHandler.Stats.CopyRollStats()
	httputils.HandleJSON(writer, statsHandler.Logger, statsCopy)
}
