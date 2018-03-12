package api_test

import (
	"bytes"
	"fmt"
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
		request.Header.Set("X-Access-Token", accessTokenHeader)

		h := &api.AppHandler{}
		h.MainHandler(responseWriter, request)

		expectedCode := http.StatusOK
		if expectedCode != responseWriter.Code {
			t.Errorf("status code didn't match: \n\t%q\n\t%q", expectedCode, responseWriter.Code)
		}

		expectedBody := []byte("authenticated with success.")
		if !bytes.Equal(expectedBody, responseWriter.Body.Bytes()) {
			t.Errorf("status code didn't match: \n\t%q\n\t%q", expectedBody, responseWriter.Body.String())
		}
	})

	t.Run("success", func(t *testing.T) {
		request, err := http.NewRequest("GET", "/api", nil)
		if err != nil {
			t.Fatalf("index request error: %s", err)
		}

		responseWriter := httptest.NewRecorder()
		accessTokenHeader := ""
		request.Header.Set("X-Access-Token", accessTokenHeader)

		h := &api.AppHandler{}
		h.MainHandler(responseWriter, request)

		expectedCode := http.StatusForbidden
		if expectedCode != responseWriter.Code {
			t.Errorf("status code didn't match: \n\t%q\n\t%q", expectedCode, responseWriter.Code)
		}

		expectedBody := []byte("you don't have access.\n")
		if !bytes.Equal(expectedBody, responseWriter.Body.Bytes()) {
			t.Errorf("status code didn't match: \n\t%q\n\t%q", expectedBody, responseWriter.Body.String())
		}
	})
}

// MockHandler ...
type MockHandler struct {
	MainHandlerInvoked bool
}

// MainHandler ...
func (m *MockHandler) MainHandler(w http.ResponseWriter, r *http.Request) {
	m.MainHandlerInvoked = true
	fmt.Fprint(w, "authenticated with success.")
}

func TestRouter(t *testing.T) {
	t.Run("get api", func(t *testing.T) {
		h := &MockHandler{}
		srv := httptest.NewServer(api.Server(h))
		defer srv.Close()

		client := &http.Client{}
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/api", srv.URL), nil)
		if err != nil {
			t.Fatalf("could not create new GET request: %v", err)
		}

		// Set token
		res, err := client.Do(req)
		if err != nil {
			t.Fatalf("could not send GET request: %v", err)
		}

		if res.StatusCode != http.StatusOK {
			t.Errorf("expected status code 200, got %v", res.Status)
		}

		if !h.MainHandlerInvoked {
			t.Errorf("mock handler expected to be called")
		}
	})
}
