package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWebhook(t *testing.T) {
	successBodey := `{
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
		{method: http.MethodPost, expectedCode: http.StatusOK, expectedBody: successBodey},
	}

	for _, testCase := range testCases {
		t.Run(testCase.method, func(t *testing.T) {
			r := httptest.NewRequest(testCase.method, "/", nil)
			w := httptest.NewRecorder()

			webhook(w, r)
			assert.Equal(t, testCase.expectedCode, w.Code, "The response code doesn't equal to expected")

			if testCase.expectedBody != "" {
				assert.JSONEq(t, testCase.expectedBody, w.Body.String(), "The response body doesn't equal to expected")
			}
		})
	}
}
