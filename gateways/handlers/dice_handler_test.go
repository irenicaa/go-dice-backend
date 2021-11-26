package handlers

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	httputils "github.com/irenicaa/go-dice-backend/http-utils"
	"github.com/irenicaa/go-dice-backend/models"
	"github.com/stretchr/testify/assert"
)

func TestDiceHandler_ServeHTTP(t *testing.T) {
	type fields struct {
		Stats  StatsRegister
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
				Stats: func() StatsRegister {
					dice := models.Dice{Tries: 2, Faces: 6}

					stats := &MockStatsRegister{}
					stats.InnerMock.On("RegisterDice", dice).Return()

					return stats
				}(),
				Logger: &MockLogger{},
			},
			args: args{
				request: httptest.NewRequest(
					http.MethodGet,
					"http://example.com/dice?tries=2&faces=6",
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
					[]byte(`{"Values":[6,4],"Sum":10,"Min":4,"Max":6}`),
				)),
				ContentLength: -1,
			},
		},
		{
			name: "error with the tries parameter",
			fields: fields{
				Stats: &MockStatsRegister{},
				Logger: func() httputils.Logger {
					logger := &MockLogger{}
					logger.InnerMock.
						On("Print", []interface{}{
							"unable to get the tries parameter: value is incorrect: " +
								"strconv.Atoi: parsing \"incorrect\": invalid syntax",
						}).
						Return().
						Times(1)

					return logger
				}(),
			},
			args: args{
				request: httptest.NewRequest(
					http.MethodGet,
					"http://example.com/dice?tries=incorrect&faces=6",
					nil,
				),
			},
			wantResponse: &http.Response{
				Status: strconv.Itoa(http.StatusBadRequest) + " " +
					http.StatusText(http.StatusBadRequest),
				StatusCode: http.StatusBadRequest,
				Proto:      "HTTP/1.1",
				ProtoMajor: 1,
				ProtoMinor: 1,
				Header:     http.Header{},
				Body: ioutil.NopCloser(bytes.NewReader([]byte(
					"unable to get the tries parameter: value is incorrect: " +
						"strconv.Atoi: parsing \"incorrect\": invalid syntax",
				))),
				ContentLength: -1,
			},
		},
		{
			name: "error with the faces parameter",
			fields: fields{
				Stats: &MockStatsRegister{},
				Logger: func() httputils.Logger {
					logger := &MockLogger{}
					logger.InnerMock.
						On("Print", []interface{}{
							"unable to get the faces parameter: value is incorrect: " +
								"strconv.Atoi: parsing \"incorrect\": invalid syntax",
						}).
						Return().
						Times(1)

					return logger
				}(),
			},
			args: args{
				request: httptest.NewRequest(
					http.MethodGet,
					"http://example.com/dice?tries=2&faces=incorrect",
					nil,
				),
			},
			wantResponse: &http.Response{
				Status: strconv.Itoa(http.StatusBadRequest) + " " +
					http.StatusText(http.StatusBadRequest),
				StatusCode: http.StatusBadRequest,
				Proto:      "HTTP/1.1",
				ProtoMajor: 1,
				ProtoMinor: 1,
				Header:     http.Header{},
				Body: ioutil.NopCloser(bytes.NewReader([]byte(
					"unable to get the faces parameter: value is incorrect: " +
						"strconv.Atoi: parsing \"incorrect\": invalid syntax",
				))),
				ContentLength: -1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rand.Seed(1)

			responseRecorder := httptest.NewRecorder()
			diceHandler := DiceHandler{
				Stats:  tt.fields.Stats,
				Logger: tt.fields.Logger,
			}
			diceHandler.ServeHTTP(responseRecorder, tt.args.request)

			tt.fields.Stats.(*MockStatsRegister).InnerMock.AssertExpectations(t)
			tt.fields.Logger.(*MockLogger).InnerMock.AssertExpectations(t)
			assert.Equal(t, tt.wantResponse, responseRecorder.Result())
		})
	}
}
