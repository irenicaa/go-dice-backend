package httputils

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

// Logger ...
type Logger interface {
	Print(arguments ...interface{})
}

// GetIntFormValue ...
func GetIntFormValue(
	request *http.Request,
	key string,
	min int,
	max int,
) (int, error) {
	value := request.FormValue(key)
	if value == "" {
		return 0, errors.New("key is missed")
	}

	valueAsInt, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("value is incorrect: %v", err)
	}
	if valueAsInt < min {
		return 0, errors.New("value too less")
	}
	if valueAsInt > max {
		return 0, errors.New("value too greater")
	}

	return valueAsInt, nil
}

// HandleError ...
func HandleError(
	writer http.ResponseWriter,
	logger Logger,
	status int,
	format string,
	arguments ...interface{},
) {
	message := fmt.Sprintf(format, arguments...)
	logger.Print(message)

	writer.WriteHeader(status)
	writer.Write([]byte(message))
}
