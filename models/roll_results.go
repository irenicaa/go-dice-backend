package models

// RollResults ...
type RollResults struct {
	Values []int
	Sum    int
	Min    int
	Max    int
}

// NewRollResults ...
func NewRollResults(values []int) RollResults {
	var sum, min, max int
	for index, value := range values {
		sum += value
		if index == 0 || value < min {
			min = value
		}
		if index == 0 || value > max {
			max = value
		}
	}

	return RollResults{Values: values, Sum: sum, Min: min, Max: max}
}
