package main

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	app := initializeApi()

	tests := []struct {
		route               string
		expectedStatusCode  int
		expectedBodyContent string
	}{
		{
			route:               "/",
			expectedStatusCode:  200,
			expectedBodyContent: "Hello, world! 👋 (ID: ",
		},
	}

	for _, testCase := range tests {
		req := httptest.NewRequest("GET", testCase.route, nil)
		resp, _ := app.Test(req, 1)

		bodyContent, err := io.ReadAll(resp.Body)
		if err != nil {
			assert.Fail(t, "cannot read response body")
		}

		err = resp.Body.Close()
		if err != nil {
			assert.Fail(t, "cannot close response body resource")
		}

		assert.Equal(t, testCase.expectedStatusCode, resp.StatusCode)
		assert.Contains(t, string(bodyContent), testCase.expectedBodyContent)
	}
}
