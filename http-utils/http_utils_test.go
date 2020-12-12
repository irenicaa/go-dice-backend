package httputils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIntFormValue(t *testing.T) {
	type args struct {
		request *http.Request
		key     string
		min     int
		max     int
	}

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				request: httptest.NewRequest(http.MethodGet, "/test?key=23", nil),
				key:     "key",
				min:     0,
				max:     100,
			},
			want:    23,
			wantErr: assert.NoError,
		},
		{
			name: "error with a missed key",
			args: args{
				request: httptest.NewRequest(http.MethodGet, "/test", nil),
				key:     "key",
				min:     0,
				max:     100,
			},
			want:    0,
			wantErr: assert.Error,
		},
		{
			name: "error with an incorrect key",
			args: args{
				request: httptest.NewRequest(http.MethodGet, "/test?key=value", nil),
				key:     "key",
				min:     0,
				max:     100,
			},
			want:    0,
			wantErr: assert.Error,
		},
		{
			name: "error with a too less value",
			args: args{
				request: httptest.NewRequest(http.MethodGet, "/test?key=23", nil),
				key:     "key",
				min:     50,
				max:     100,
			},
			want:    0,
			wantErr: assert.Error,
		},
		{
			name: "error with a too greater value",
			args: args{
				request: httptest.NewRequest(http.MethodGet, "/test?key=23", nil),
				key:     "key",
				min:     0,
				max:     10,
			},
			want:    0,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err :=
				GetIntFormValue(tt.args.request, tt.args.key, tt.args.min, tt.args.max)

			assert.Equal(t, tt.want, got)
			tt.wantErr(t, err)
		})
	}
}
