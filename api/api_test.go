package api_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aymone/sample/api"
)

func TestAuth(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		if !api.Auth("password") {
			t.Error("authService expected to be true")
		}
	})

	t.Run("fail", func(t *testing.T) {
		if api.Auth("") {
			t.Error("authService expected to be false")
		}
	})
}

func TestMainHandler(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		request, err := http.NewRequest("GET", "/api", nil)
		if err != nil {
			t.Fatalf("index request error: %s", err)
		}

		responseWriter := httptest.NewRecorder()
		accessTokenHeader := "password"
		expectedCode := http.StatusOK
		expectedBody := []byte("authenticated with success.\n")

		request.Header.Set("X-Access-Token", accessTokenHeader)
		api.MainHandler(responseWriter, request)

		if expectedCode != responseWriter.Code {
			t.Errorf("status code didn't match: \n\t%q\n\t%q", expectedCode, responseWriter.Code)
		}

		if !bytes.Equal(expectedBody, responseWriter.Body.Bytes()) {
			t.Errorf("status code didn't match: \n\t%q\n\t%q", expectedBody, responseWriter.Body.String())
		}
	})
}
