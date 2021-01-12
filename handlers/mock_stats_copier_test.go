package handlers

import (
	"github.com/stretchr/testify/mock"
)

type MockStatsCopier struct {
	InnerMock mock.Mock
}

func (mock *MockStatsCopier) CopyData() map[string]int {
	results := mock.InnerMock.Called()
	return results.Get(0).(map[string]int)
}
