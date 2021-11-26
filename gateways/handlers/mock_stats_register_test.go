package handlers

import (
	"github.com/irenicaa/go-dice-backend/models"
	"github.com/stretchr/testify/mock"
)

type MockStatsRegister struct {
	InnerMock mock.Mock
}

func (mock *MockStatsRegister) RegisterDice(dice models.Dice) {
	mock.InnerMock.Called(dice)
}
