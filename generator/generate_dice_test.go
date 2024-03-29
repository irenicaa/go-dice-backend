package generator

import (
	"math/rand"
	"testing"

	"github.com/irenicaa/go-dice-backend/v2/models"
	"github.com/stretchr/testify/assert"
)

func TestGenerateDice(t *testing.T) {
	rand.Seed(1)

	dice := models.Dice{Tries: 2, Faces: 6}
	values, err := GenerateDice(dice)

	wantedValues := []int{6, 4}
	assert.Equal(t, wantedValues, values)
	assert.NoError(t, err)
}
