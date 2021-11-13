package handlers

import (
	"net/http"

	httputils "github.com/irenicaa/go-dice-generator/http-utils"
	"github.com/irenicaa/go-dice-generator/models"
)

// StatsCopier ...
type StatsCopier interface {
	CopyData() models.RollStatsData
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
	statsCopy := statsHandler.Stats.CopyData()
	httputils.HandleJSON(writer, statsHandler.Logger, statsCopy)
}
