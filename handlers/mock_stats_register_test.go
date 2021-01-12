package handlers

import (
	"github.com/irenicaa/go-dice-generator/models"
	"github.com/stretchr/testify/mock"
)

type MockStatsRegister struct {
	InnerMock mock.Mock
}

func (mock *MockStatsRegister) Register(dice models.Dice) {
	mock.InnerMock.Called(dice)
}
