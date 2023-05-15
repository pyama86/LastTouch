package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AccessEndpoint(api func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(api)
}
func TestGetUpdateTimeHandlerSuccess(t *testing.T) {
	os.Setenv("LASTTOUCH_USER", "user")
	os.Setenv("LASTTOUCH_PASSWORD", "password")

	req, err := http.NewRequest("GET", "http://localhost:8080/getUpdateTime", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.SetBasicAuth("user", "password")

	q := req.URL.Query()
	q.Add("db", "lasttouch")
	q.Add("table", "example")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()

	handler := AccessEndpoint(getUpdateTimeHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, rr.Body.String())

	var resp Response
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetUpdateTimeHandlerMissingToken(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/getUpdateTime", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := AccessEndpoint(getUpdateTimeHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusForbidden, rr.Code)
}

func TestGetUpdateTimeHandlerInvalidToken(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/getUpdateTime", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "invalid_token")

	rr := httptest.NewRecorder()

	handler := AccessEndpoint(getUpdateTimeHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusForbidden, rr.Code)
}

func TestGetUpdateTime(t *testing.T) {
	t.Run("success_case", func(t *testing.T) {
		updateTime, err := getUpdateTime("lasttouch", "example")
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if updateTime == "" {
			t.Error("expected non-empty update time")
		}
	})

}
