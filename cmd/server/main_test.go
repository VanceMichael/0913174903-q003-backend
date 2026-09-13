package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/health", nil)
	router().ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK { t.Fatalf("状态码为 %d", recorder.Code) }
}
