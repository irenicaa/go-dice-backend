// +build integration

package tests

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"testing"
	"time"

	"github.com/irenicaa/go-dice-backend/v2/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var port = flag.Int("port", 8080, "server port")
var requestCount = flag.Int("requestCount", 10, "test request count")

func TestDiceBackend(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	statsURL := fmt.Sprintf("http://localhost:%d/api/v1/stats", *port)
	stats, err := loadStats(statsURL)
	require.NoError(t, err)

	for i := 0; i < *requestCount; i++ {
		tries := rand.Intn(100) + 1
		faces := rand.Intn(99) + 2

		dice := models.Dice{Tries: tries, Faces: faces}
		stats[dice.String()]++

		url := fmt.Sprintf(
			"http://localhost:%d/api/v1/dice?tries=%d&faces=%d",
			*port,
			tries,
			faces,
		)
		response, err := http.Post(url, "text/plain", nil)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, response.StatusCode)
	}

	gotStats, err := loadStats(statsURL)
	require.NoError(t, err)

	assert.Equal(t, stats, gotStats)
}

func loadStats(url string) (models.RollStats, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("unable to send the request: %v", err)
	}
	if response.StatusCode != http.StatusOK {
		return nil,
			fmt.Errorf("request was failed with the status %d", response.StatusCode)
	}

	var stats models.RollStats
	if err := json.NewDecoder(response.Body).Decode(&stats); err != nil {
		return nil, fmt.Errorf("unable to decode the stats: %v", err)
	}

	return stats, nil
}
