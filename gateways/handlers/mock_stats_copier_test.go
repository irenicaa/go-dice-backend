package handlers

import (
	"github.com/irenicaa/go-dice-backend/models"
	"github.com/stretchr/testify/mock"
)

type MockStatsCopier struct {
	InnerMock mock.Mock
}

func (mock *MockStatsCopier) CopyRollStats() models.RollStats {
	results := mock.InnerMock.Called()
	return results.Get(0).(models.RollStats)
}
