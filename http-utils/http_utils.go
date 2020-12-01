package httputils

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

// GetIntFormValue ...
func GetIntFormValue(request *http.Request, key string) (int, error) {
	value := request.FormValue(key)
	if value == "" {
		return 0, errors.New("key is missed")
	}

	valueAsInt, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("value is incorrect: %v", err)
	}

	return valueAsInt, nil
}
