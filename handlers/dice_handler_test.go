package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	httputils "github.com/irenicaa/go-dice-generator/http-utils"
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
