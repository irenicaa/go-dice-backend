package models

import "sync"

// RollStatsData ...
type RollStatsData map[string]int

type locker interface {
	sync.Locker

	RLock()
	RUnlock()
}

// RollStats ...
type RollStats struct {
	data  RollStatsData
	mutex locker
}

// NewRollStats ...
func NewRollStats() RollStats {
	return RollStats{data: RollStatsData{}, mutex: &sync.RWMutex{}}
}

// Register ...
func (rollStats RollStats) Register(dice Dice) {
	rollStats.mutex.Lock()
	defer rollStats.mutex.Unlock()

	rollStats.data[dice.String()]++
}

// CopyData ...
func (rollStats RollStats) CopyData() RollStatsData {
	rollStats.mutex.RLock()
	defer rollStats.mutex.RUnlock()

	dataCopy := RollStatsData{}
	for dice, count := range rollStats.data {
		dataCopy[dice] = count
	}

	return dataCopy
}
