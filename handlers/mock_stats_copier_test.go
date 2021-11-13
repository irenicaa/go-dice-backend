package handlers

import (
	"github.com/irenicaa/go-dice-generator/models"
	"github.com/stretchr/testify/mock"
)

type MockStatsCopier struct {
	InnerMock mock.Mock
}

func (mock *MockStatsCopier) CopyData() models.RollStatsData {
	results := mock.InnerMock.Called()
	return results.Get(0).(models.RollStatsData)
}
