package handlers

import (
	"net/http"

	httputils "github.com/irenicaa/go-dice-generator/http-utils"
)

// StatsCopier ...
type StatsCopier interface {
	CopyData() map[string]int
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
//   @success 200 {object} map[string]int
func (statsHandler StatsHandler) ServeHTTP(
	writer http.ResponseWriter,
	request *http.Request,
) {
	statsCopy := statsHandler.Stats.CopyData()
	httputils.HandleJSON(writer, statsHandler.Logger, statsCopy)
}
