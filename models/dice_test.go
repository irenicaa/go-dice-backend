package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDice_String(t *testing.T) {
	type fields struct {
		Tries int
		Faces int
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "success",
			fields: fields{Tries: 2, Faces: 3},
			want:   "2d3",
		},
		{
			name:   "success (with a single try)",
			fields: fields{Tries: 1, Faces: 2},
			want:   "d2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dice := Dice{
				Tries: tt.fields.Tries,
				Faces: tt.fields.Faces,
			}
			got := dice.String()

			assert.Equal(t, tt.want, got)
		})
	}
}
