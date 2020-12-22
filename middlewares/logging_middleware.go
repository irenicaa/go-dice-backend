package middlewares

import (
	"fmt"
	"net/http"
	"time"

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
		startTime := time.Now()
		handler.ServeHTTP(writer, request)

		elapsedTime := time.Now().Sub(startTime)
		message := fmt.Sprintf("%s %s %s", request.Method, request.URL, elapsedTime)
		logger.Print(message)
	})
}
