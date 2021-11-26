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

// CopyRollStats ...
func (rollStats RollStats) CopyRollStats() RollStatsData {
	rollStats.mutex.RLock()
	defer rollStats.mutex.RUnlock()

	dataCopy := RollStatsData{}
	for dice, count := range rollStats.data {
		dataCopy[dice] = count
	}

	return dataCopy
}

// RegisterDice ...
func (rollStats RollStats) RegisterDice(dice Dice) {
	rollStats.mutex.Lock()
	defer rollStats.mutex.Unlock()

	rollStats.data[dice.String()]++
}
