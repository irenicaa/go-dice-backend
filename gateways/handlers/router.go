package handlers

import (
	"net/http"

	httputils "github.com/irenicaa/go-http-utils"
)

// Router ...
type Router struct {
	BaseURL      string
	DiceHandler  DiceHandler
	StatsHandler StatsHandler
	Logger       httputils.Logger
}

// ServeHTTP ...
func (router Router) ServeHTTP(
	writer http.ResponseWriter,
	request *http.Request,
) {
	method, urlPath := request.Method, request.URL.Path
	if urlPath == router.BaseURL+"/dice" && method == http.MethodPost {
		router.DiceHandler.ServeHTTP(writer, request)
	} else if urlPath == router.BaseURL+"/stats" && method == http.MethodGet {
		router.StatsHandler.ServeHTTP(writer, request)
	} else {
		status, message := http.StatusNotFound, http.StatusText(http.StatusNotFound)
		httputils.HandleError(writer, router.Logger, status, message)
	}
}
