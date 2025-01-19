package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestWebhook(t *testing.T) {
	handler := http.HandlerFunc(webhook)
	server := httptest.NewServer(handler)
	defer server.Close()
	successBody := `{
		"response": {
			"text": "Sorry, i don't do anyting"
		},
		"version": "1.0"
	}`

	testCases := []struct {
		method       string
		expectedCode int
		expectedBody string
	}{
		{method: http.MethodGet, expectedCode: http.StatusMethodNotAllowed, expectedBody: ""},
		{method: http.MethodPut, expectedCode: http.StatusMethodNotAllowed, expectedBody: ""},
		{method: http.MethodDelete, expectedCode: http.StatusMethodNotAllowed, expectedBody: ""},
		{method: http.MethodPost, expectedCode: http.StatusOK, expectedBody: successBody},
	}

	for _, testCase := range testCases {
		t.Run(testCase.method, func(t *testing.T) {
			req := resty.New().R()
			req.Method = testCase.method
			req.URL = server.URL

			resp, err := req.Send()
			assert.NoError(t, err, "error making HTTP request")
			assert.Equal(t, testCase.expectedCode, resp.StatusCode(), "Response code didn't match expected")

			if testCase.expectedBody != "" {
				assert.JSONEq(t, testCase.expectedBody, string(resp.Body()))
			}
		})
	}

}
