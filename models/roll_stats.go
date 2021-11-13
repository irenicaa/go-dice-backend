package models

import "sync"

type locker interface {
	sync.Locker

	RLock()
	RUnlock()
}

// RollStats ...
type RollStats struct {
	data  map[string]int
	mutex locker
}

// NewRollStats ...
func NewRollStats() RollStats {
	return RollStats{data: map[string]int{}, mutex: &sync.RWMutex{}}
}

// Register ...
func (rollStats RollStats) Register(dice Dice) {
	rollStats.mutex.Lock()
	defer rollStats.mutex.Unlock()

	rollStats.data[dice.String()]++
}

// CopyData ...
func (rollStats RollStats) CopyData() map[string]int {
	rollStats.mutex.RLock()
	defer rollStats.mutex.RUnlock()

	dataCopy := map[string]int{}
	for dice, count := range rollStats.data {
		dataCopy[dice] = count
	}

	return dataCopy
}
