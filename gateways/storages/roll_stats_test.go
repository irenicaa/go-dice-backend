package storages

import (
	"sync"
	"testing"

	"github.com/irenicaa/go-dice-backend/models"
	"github.com/stretchr/testify/assert"
)

func TestNewRollStats(t *testing.T) {
	rollStats := NewRollStats()

	assert.Equal(t, models.RollStats{}, rollStats.data)
	assert.Equal(t, &sync.RWMutex{}, rollStats.mutex)
}

func TestRollStats_CopyRollStats(t *testing.T) {
	type fields struct {
		data  models.RollStats
		mutex locker
	}

	tests := []struct {
		name    string
		fields  fields
		want    models.RollStats
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			fields: fields{
				data: models.RollStats{"2d3": 5, "4d2": 12},
				mutex: func() locker {
					locker := &MockLocker{}
					locker.InnerMock.On("RLock").Return()
					locker.InnerMock.On("RUnlock").Return()

					return locker
				}(),
			},
			want:    models.RollStats{"2d3": 5, "4d2": 12},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rollStats := RollStats{
				data:  tt.fields.data,
				mutex: tt.fields.mutex,
			}
			got, err := rollStats.CopyRollStats()
			rollStats.data["10d100"] = 23

			tt.fields.mutex.(*MockLocker).InnerMock.AssertExpectations(t)
			assert.Equal(t, tt.want, got)
			tt.wantErr(t, err)
		})
	}
}

func TestRollStats_RegisterDice(t *testing.T) {
	type fields struct {
		data  models.RollStats
		mutex locker
	}
	type args struct {
		dice models.Dice
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		wantData models.RollStats
	}{
		{
			name: "existing key",
			fields: fields{
				data: models.RollStats{"2d3": 5, "4d2": 12},
				mutex: func() locker {
					locker := &MockLocker{}
					locker.InnerMock.On("Lock").Return()
					locker.InnerMock.On("Unlock").Return()

					return locker
				}(),
			},
			args: args{
				dice: models.Dice{Tries: 4, Faces: 2},
			},
			wantData: models.RollStats{"2d3": 5, "4d2": 13},
		},
		{
			name: "not existing key",
			fields: fields{
				data: models.RollStats{"2d3": 5, "4d2": 12},
				mutex: func() locker {
					locker := &MockLocker{}
					locker.InnerMock.On("Lock").Return()
					locker.InnerMock.On("Unlock").Return()

					return locker
				}(),
			},
			args: args{
				dice: models.Dice{Tries: 10, Faces: 100},
			},
			wantData: models.RollStats{"2d3": 5, "4d2": 12, "10d100": 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rollStats := RollStats{
				data:  tt.fields.data,
				mutex: tt.fields.mutex,
			}
			rollStats.RegisterDice(tt.args.dice)

			tt.fields.mutex.(*MockLocker).InnerMock.AssertExpectations(t)
			assert.Equal(t, tt.wantData, rollStats.data)
		})
	}
}
