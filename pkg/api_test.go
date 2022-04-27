package deus_cc

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetEvent(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		Description      string
		Method           string
		Payload          interface{}
		ExpectedResponse int
		ExpectedSucceed  bool
	}{
		{
			Description:      "when HTTP method is not allowed, return status != 200",
			Method:           http.MethodGet,
			ExpectedResponse: http.StatusMethodNotAllowed,
			ExpectedSucceed:  false,
		},
		{
			Description:      "when HTTP method is allowed, but there is no payload",
			Method:           http.MethodPost,
			ExpectedResponse: http.StatusBadRequest,
			ExpectedSucceed:  false,
		},
		{
			Description: "when HTTP method is allowed, but mandatory params are not present",
			Method:      http.MethodPost,
			Payload: map[string]string{
				"foobar": "foobar",
			},
			ExpectedResponse: http.StatusBadRequest,
			ExpectedSucceed:  false,
		},
		{
			Description:      "when HTTP method is allowed, but data is not valid",
			Method:           http.MethodPost,
			Payload:          "foobar",
			ExpectedResponse: http.StatusUnprocessableEntity,
			ExpectedSucceed:  false,
		},
		{
			Description: "when HTTP method is allowed, and mandatory params are present",
			Method:      http.MethodPost,
			Payload: map[string]string{
				"url":  "foobar",
				"uuid": "foobar",
			},
			ExpectedResponse: http.StatusOK,
			ExpectedSucceed:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {

			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			}))
			defer func() { testServer.Close() }()

			jsonStr, err := json.Marshal(test.Payload)
			assert.NoError(err)

			req := httptest.NewRequest(test.Method, testServer.URL, bytes.NewBuffer(jsonStr))
			res := httptest.NewRecorder()

			handler := SetEvent()
			handler.ServeHTTP(res, req)
			assert.Equal(res.Result().StatusCode, test.ExpectedResponse)
		})
	}
}

func TestGetDistinctVisitors(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		Description      string
		Method           string
		QueryParam       string
		QueryValue       string
		ExpectedResponse int
		ExpectedSucceed  bool
	}{
		{
			Description:      "when HTTP method is not allowed, return status != 200",
			Method:           http.MethodPost,
			ExpectedResponse: http.StatusMethodNotAllowed,
			ExpectedSucceed:  false,
		},
		{
			Description:      "when HTTP method is allowed, but there is no querystring",
			Method:           http.MethodGet,
			ExpectedResponse: http.StatusInternalServerError,
			ExpectedSucceed:  false,
		},
		{
			Description:      "when HTTP method is allowed, but there is no `url` param",
			Method:           http.MethodGet,
			QueryParam:       "foobar",
			QueryValue:       "foobar",
			ExpectedResponse: http.StatusInternalServerError,
			ExpectedSucceed:  false,
		},
		{
			Description:      "when HTTP method is allowed, and there is `url` param",
			Method:           http.MethodGet,
			QueryParam:       "url",
			QueryValue:       "foobar",
			ExpectedResponse: http.StatusOK,
			ExpectedSucceed:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {

			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			}))
			defer func() { testServer.Close() }()

			req := httptest.NewRequest(test.Method, testServer.URL, nil)
			res := httptest.NewRecorder()

			if test.QueryParam != "" {
				q := req.URL.Query()
				q.Add(test.QueryParam, test.QueryValue)
				req.URL.RawQuery = q.Encode()
			}

			handler := GetDistinctVisitors()
			handler.ServeHTTP(res, req)
			assert.Equal(res.Result().StatusCode, test.ExpectedResponse)
		})
	}
}
