package handlers

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"testing/iotest"

	"github.com/irenicaa/go-dice-backend/models"
	httputils "github.com/irenicaa/go-http-utils"
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
					data := models.RollStats{"2d3": 5, "4d2": 12}

					stats := &MockStatsCopier{}
					stats.InnerMock.On("CopyRollStats").Return(data, nil)

					return stats
				}(),
				Logger: &MockLogger{},
			},
			args: args{
				request: httptest.NewRequest(
					http.MethodGet,
					"http://example.com/api/v1/stats",
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
		{
			name: "error",
			fields: fields{
				Stats: func() StatsCopier {
					stats := &MockStatsCopier{}
					stats.InnerMock.
						On("CopyRollStats").
						Return(models.RollStats(nil), iotest.ErrTimeout)

					return stats
				}(),
				Logger: func() httputils.Logger {
					logger := &MockLogger{}
					logger.InnerMock.
						On("Print", []interface{}{"unable to copy the roll stats: timeout"}).
						Return().
						Times(1)

					return logger
				}(),
			},
			args: args{
				request: httptest.NewRequest(
					http.MethodGet,
					"http://example.com/api/v1/stats",
					nil,
				),
			},
			wantResponse: &http.Response{
				Status: strconv.Itoa(http.StatusInternalServerError) + " " +
					http.StatusText(http.StatusInternalServerError),
				StatusCode: http.StatusInternalServerError,
				Proto:      "HTTP/1.1",
				ProtoMajor: 1,
				ProtoMinor: 1,
				Header:     http.Header{},
				Body: ioutil.NopCloser(bytes.NewReader(
					[]byte("unable to copy the roll stats: timeout"),
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
