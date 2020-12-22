package middlewares

import (
	"net/http"

	httputils "github.com/irenicaa/go-dice-generator/http-utils"
)

// LoggingMiddleware ...
func LoggingMiddleware(
	handler http.Handler,
	logger httputils.Logger,
) http.Handler {
	return http.HandlerFunc(func(
		writer http.ResponseWriter,
		request *http.Request,
	) {
		logger.Print("received a request at " + request.URL.String())

		handler.ServeHTTP(writer, request)
	})
}
