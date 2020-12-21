package httputils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
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

func TestHandleError(t *testing.T) {
	type args struct {
		logger    Logger
		status    int
		format    string
		arguments []interface{}
	}

	tests := []struct {
		name         string
		args         args
		wantResponse *http.Response
	}{
		{
			name: "succes",
			args: args{
				logger: func() Logger {
					logger := &MockLogger{}
					logger.InnerMock.
						On("Print", []interface{}{"test: 23 one"}).
						Return().
						Times(1)

					return logger
				}(),
				status:    http.StatusNotFound,
				format:    "test: %d %s",
				arguments: []interface{}{23, "one"},
			},
			wantResponse: &http.Response{
				Status: strconv.Itoa(http.StatusNotFound) + " " +
					http.StatusText(http.StatusNotFound),
				StatusCode:    http.StatusNotFound,
				Proto:         "HTTP/1.1",
				ProtoMajor:    1,
				ProtoMinor:    1,
				Header:        http.Header{},
				Body:          ioutil.NopCloser(bytes.NewReader([]byte("test: 23 one"))),
				ContentLength: -1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			responseRecorder := httptest.NewRecorder()
			HandleError(
				responseRecorder,
				tt.args.logger,
				tt.args.status,
				tt.args.format,
				tt.args.arguments...,
			)

			tt.args.logger.(*MockLogger).InnerMock.AssertExpectations(t)
			assert.Equal(t, tt.wantResponse, responseRecorder.Result())
		})
	}
}

func TestHandleJSON(t *testing.T) {
	type testData struct {
		FieldOne int
		FieldTwo int
	}
	type incorrectTestData struct {
		FieldOne   int
		FieldTwo   int
		FieldThree func()
	}
	type args struct {
		logger Logger
		data   interface{}
	}

	tests := []struct {
		name         string
		args         args
		wantResponse *http.Response
	}{
		{
			name: "success",
			args: args{
				logger: &MockLogger{},
				data:   testData{FieldOne: 23, FieldTwo: 42},
			},
			wantResponse: &http.Response{
				Status: strconv.Itoa(http.StatusOK) + " " +
					http.StatusText(http.StatusOK),
				StatusCode: http.StatusOK,
				Proto:      "HTTP/1.1",
				ProtoMajor: 1,
				ProtoMinor: 1,
				Header:     http.Header{"Content-Type": {"application/json"}},
				Body: ioutil.NopCloser(bytes.NewReader(
					[]byte(`{"FieldOne":23,"FieldTwo":42}`),
				)),
				ContentLength: -1,
			},
		},
		{
			name: "error",
			args: args{
				logger: func() Logger {
					logger := &MockLogger{}
					logger.InnerMock.
						On("Print", []interface{}{
							"unable to marshal the data: json: unsupported type: func()",
						}).
						Return().
						Times(1)

					return logger
				}(),
				data: incorrectTestData{
					FieldOne:   23,
					FieldTwo:   42,
					FieldThree: func() {},
				},
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
					[]byte(`unable to marshal the data: json: unsupported type: func()`),
				)),
				ContentLength: -1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			responseRecorder := httptest.NewRecorder()
			HandleJSON(responseRecorder, tt.args.logger, tt.args.data)

			tt.args.logger.(*MockLogger).InnerMock.AssertExpectations(t)
			assert.Equal(t, tt.wantResponse, responseRecorder.Result())
		})
	}
}
