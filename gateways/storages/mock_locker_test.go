package storages

import "github.com/stretchr/testify/mock"

type MockLocker struct {
	InnerMock mock.Mock
}

func (mock *MockLocker) Lock() {
	mock.InnerMock.Called()
}

func (mock *MockLocker) Unlock() {
	mock.InnerMock.Called()
}

func (mock *MockLocker) RLock() {
	mock.InnerMock.Called()
}

func (mock *MockLocker) RUnlock() {
	mock.InnerMock.Called()
}
