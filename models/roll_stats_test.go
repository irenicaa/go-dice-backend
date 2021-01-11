package models

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRollStats(t *testing.T) {
	rollStats := NewRollStats()

	assert.Equal(t, map[string]int{}, rollStats.data)
	assert.Equal(t, &sync.RWMutex{}, rollStats.mutex)
}

func TestRollStats_Register(t *testing.T) {
	type fields struct {
		data  map[string]int
		mutex Locker
	}
	type args struct {
		dice Dice
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		wantData map[string]int
	}{
		{
			name: "existing key",
			fields: fields{
				data: map[string]int{"2d3": 5, "4d2": 12},
				mutex: func() Locker {
					locker := &MockLocker{}
					locker.InnerMock.On("Lock").Return()
					locker.InnerMock.On("Unlock").Return()

					return locker
				}(),
			},
			args: args{
				dice: Dice{Tries: 4, Faces: 2},
			},
			wantData: map[string]int{"2d3": 5, "4d2": 13},
		},
		{
			name: "not existing key",
			fields: fields{
				data: map[string]int{"2d3": 5, "4d2": 12},
				mutex: func() Locker {
					locker := &MockLocker{}
					locker.InnerMock.On("Lock").Return()
					locker.InnerMock.On("Unlock").Return()

					return locker
				}(),
			},
			args: args{
				dice: Dice{Tries: 10, Faces: 100},
			},
			wantData: map[string]int{"2d3": 5, "4d2": 12, "10d100": 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rollStats := RollStats{
				data:  tt.fields.data,
				mutex: tt.fields.mutex,
			}
			rollStats.Register(tt.args.dice)

			tt.fields.mutex.(*MockLocker).InnerMock.AssertExpectations(t)
			assert.Equal(t, tt.wantData, rollStats.data)
		})
	}
}

func TestRollStats_CopyData(t *testing.T) {
	type fields struct {
		data  map[string]int
		mutex Locker
	}

	tests := []struct {
		name   string
		fields fields
		want   map[string]int
	}{
		{
			name: "success",
			fields: fields{
				data: map[string]int{"2d3": 5, "4d2": 12},
				mutex: func() Locker {
					locker := &MockLocker{}
					locker.InnerMock.On("RLock").Return()
					locker.InnerMock.On("RUnlock").Return()

					return locker
				}(),
			},
			want: map[string]int{"2d3": 5, "4d2": 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rollStats := RollStats{
				data:  tt.fields.data,
				mutex: tt.fields.mutex,
			}
			got := rollStats.CopyData()
			rollStats.data["10d100"] = 23

			tt.fields.mutex.(*MockLocker).InnerMock.AssertExpectations(t)
			assert.Equal(t, tt.want, got)
		})
	}
}
