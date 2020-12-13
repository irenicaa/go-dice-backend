package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRollResults(t *testing.T) {
	type args struct {
		values []int
	}

	tests := []struct {
		name string
		args args
		want RollResults
	}{
		{
			name: "empty",
			args: args{values: []int{}},
			want: RollResults{Values: []int{}},
		},
		{
			name: "nonempty",
			args: args{values: []int{0, 9, 1, 8, 2, 7, 3, 6, 4, 5}},
			want: RollResults{
				Values: []int{0, 9, 1, 8, 2, 7, 3, 6, 4, 5},
				Sum:    45,
				Min:    0,
				Max:    9,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRollResults(tt.args.values)

			assert.Equal(t, tt.want, got)
		})
	}
}
