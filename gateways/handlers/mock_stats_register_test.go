package handlers

import (
	"github.com/irenicaa/go-dice-backend/v2/models"
	"github.com/stretchr/testify/mock"
)

type MockStatsRegister struct {
	InnerMock mock.Mock
}

func (mock *MockStatsRegister) RegisterDice(dice models.Dice) error {
	results := mock.InnerMock.Called(dice)
	return results.Error(0)
}
