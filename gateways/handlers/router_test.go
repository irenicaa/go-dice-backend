package handlers

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/irenicaa/go-dice-backend/generator"
	httputils "github.com/irenicaa/go-dice-backend/http-utils"
	"github.com/irenicaa/go-dice-backend/models"
	"github.com/stretchr/testify/assert"
)

func TestRouter_ServeHTTP(t *testing.T) {
	type fields struct {
		BaseURL       string
		StatsRegister StatsRegister
		StatsCopier   StatsCopier
		DiceGenerator DiceGenerator
		Logger        httputils.Logger
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
			name: "success with generating of dice rolls",
			fields: fields{
				BaseURL: "/api/v1",
				StatsRegister: func() StatsRegister {
					dice := models.Dice{Tries: 2, Faces: 6}

					stats := &MockStatsRegister{}
					stats.InnerMock.On("RegisterDice", dice).Return(nil)

					return stats
				}(),
				StatsCopier:   &MockStatsCopier{},
				DiceGenerator: generator.GenerateDice,
				Logger:        &MockLogger{},
			},
			args: args{
				request: httptest.NewRequest(
					http.MethodPost,
					"http://example.com/api/v1/dice?tries=2&faces=6",
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
			name: "success with getting of stats of dice rolls",
			fields: fields{
				BaseURL:       "/api/v1",
				StatsRegister: &MockStatsRegister{},
				StatsCopier: func() StatsCopier {
					data := models.RollStats{"2d3": 5, "4d2": 12}

					stats := &MockStatsCopier{}
					stats.InnerMock.On("CopyRollStats").Return(data, nil)

					return stats
				}(),
				DiceGenerator: nil,
				Logger:        &MockLogger{},
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
				BaseURL:       "/api/v1",
				StatsRegister: &MockStatsRegister{},
				StatsCopier:   &MockStatsCopier{},
				DiceGenerator: nil,
				Logger: func() httputils.Logger {
					logger := &MockLogger{}
					logger.InnerMock.
						On("Print", []interface{}{http.StatusText(http.StatusNotFound)}).
						Return().
						Times(1)

					return logger
				}(),
			},
			args: args{
				request: httptest.NewRequest(
					http.MethodGet,
					"http://example.com/api/v1/incorrect",
					nil,
				),
			},
			wantResponse: &http.Response{
				Status: strconv.Itoa(http.StatusNotFound) + " " +
					http.StatusText(http.StatusNotFound),
				StatusCode: http.StatusNotFound,
				Proto:      "HTTP/1.1",
				ProtoMajor: 1,
				ProtoMinor: 1,
				Header:     http.Header{},
				Body: ioutil.NopCloser(bytes.NewReader([]byte(
					http.StatusText(http.StatusNotFound),
				))),
				ContentLength: -1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rand.Seed(1)

			responseRecorder := httptest.NewRecorder()
			router := Router{
				BaseURL: tt.fields.BaseURL,
				DiceHandler: DiceHandler{
					Stats:         tt.fields.StatsRegister,
					DiceGenerator: tt.fields.DiceGenerator,
					Logger:        tt.fields.Logger,
				},
				StatsHandler: StatsHandler{
					Stats:  tt.fields.StatsCopier,
					Logger: tt.fields.Logger,
				},
				Logger: tt.fields.Logger,
			}
			router.ServeHTTP(responseRecorder, tt.args.request)

			tt.fields.StatsRegister.(*MockStatsRegister).InnerMock.AssertExpectations(t)
			tt.fields.StatsCopier.(*MockStatsCopier).InnerMock.AssertExpectations(t)
			tt.fields.Logger.(*MockLogger).InnerMock.AssertExpectations(t)
			assert.Equal(t, tt.wantResponse, responseRecorder.Result())
		})
	}
}
