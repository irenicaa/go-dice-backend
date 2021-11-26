package handlers

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	httputils "github.com/irenicaa/go-dice-backend/http-utils"
	"github.com/irenicaa/go-dice-backend/models"
	"github.com/stretchr/testify/assert"
)

func TestStatsHandler_ServeHTTP(t *testing.T) {
	type fields struct {
		Stats  StatsCopier
		Logger httputils.Logger
	}
	type args struct {
		request *http.Request
	}

	tests := []struct {
		name         string
		fields       fields
		args         args
		wantResponse *http.Response
	}{
		{
			name: "success",
			fields: fields{
				Stats: func() StatsCopier {
					data := models.RollStatsData{"2d3": 5, "4d2": 12}

					stats := &MockStatsCopier{}
					stats.InnerMock.On("CopyData").Return(data)

					return stats
				}(),
				Logger: &MockLogger{},
			},
			args: args{
				request: httptest.NewRequest(
					http.MethodGet,
					"http://example.com/stats",
					nil,
				),
			},
			wantResponse: &http.Response{
				Status: strconv.Itoa(http.StatusOK) + " " +
					http.StatusText(http.StatusOK),
				StatusCode: http.StatusOK,
				Proto:      "HTTP/1.1",
				ProtoMajor: 1,
				ProtoMinor: 1,
				Header:     http.Header{"Content-Type": []string{"application/json"}},
				Body: ioutil.NopCloser(bytes.NewReader(
					[]byte(`{"2d3":5,"4d2":12}`),
				)),
				ContentLength: -1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			responseRecorder := httptest.NewRecorder()
			statsHandler := StatsHandler{
				Stats:  tt.fields.Stats,
				Logger: tt.fields.Logger,
			}
			statsHandler.ServeHTTP(responseRecorder, tt.args.request)

			tt.fields.Stats.(*MockStatsCopier).InnerMock.AssertExpectations(t)
			tt.fields.Logger.(*MockLogger).InnerMock.AssertExpectations(t)
			assert.Equal(t, tt.wantResponse, responseRecorder.Result())
		})
	}
}
