package storages

import (
	"sync"

	"github.com/irenicaa/go-dice-backend/v2/models"
)

type locker interface {
	sync.Locker

	RLock()
	RUnlock()
}

// RollStats ...
type RollStats struct {
	data  models.RollStats
	mutex locker
}

// NewRollStats ...
func NewRollStats() RollStats {
	return RollStats{data: models.RollStats{}, mutex: &sync.RWMutex{}}
}

// CopyRollStats ...
func (rollStats RollStats) CopyRollStats() (models.RollStats, error) {
	rollStats.mutex.RLock()
	defer rollStats.mutex.RUnlock()

	dataCopy := models.RollStats{}
	for dice, count := range rollStats.data {
		dataCopy[dice] = count
	}

	return dataCopy, nil
}

// RegisterDice ...
func (rollStats RollStats) RegisterDice(dice models.Dice) error {
	rollStats.mutex.Lock()
	defer rollStats.mutex.Unlock()

	rollStats.data[dice.String()]++

	return nil
}
