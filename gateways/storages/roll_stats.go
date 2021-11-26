package storages

import (
	"sync"

	"github.com/irenicaa/go-dice-backend/models"
)

type locker interface {
	sync.Locker

	RLock()
	RUnlock()
}

// RollStats ...
type RollStats struct {
	data  models.RollStatsData
	mutex locker
}

// NewRollStats ...
func NewRollStats() RollStats {
	return RollStats{data: models.RollStatsData{}, mutex: &sync.RWMutex{}}
}

// CopyRollStats ...
func (rollStats RollStats) CopyRollStats() models.RollStatsData {
	rollStats.mutex.RLock()
	defer rollStats.mutex.RUnlock()

	dataCopy := models.RollStatsData{}
	for dice, count := range rollStats.data {
		dataCopy[dice] = count
	}

	return dataCopy
}

// RegisterDice ...
func (rollStats RollStats) RegisterDice(dice models.Dice) {
	rollStats.mutex.Lock()
	defer rollStats.mutex.Unlock()

	rollStats.data[dice.String()]++
}
