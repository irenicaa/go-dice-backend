package generator

import (
	"math/rand"

	"github.com/irenicaa/go-dice-backend/models"
)

// GenerateDice ...
func GenerateDice(dice models.Dice) ([]int, error) {
	var values []int
	for try := 0; try < dice.Tries; try++ {
		value := rand.Intn(dice.Faces) + 1
		values = append(values, value)
	}

	return values, nil
}
