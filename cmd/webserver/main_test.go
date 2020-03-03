package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)
	handler.ServeHTTP(recorder, req)

	status := recorder.Code
	if status != http.StatusOK {
		t.Errorf("hello handler did not return 200: got %d", status)
	}

	body := recorder.Body.String()
	expected := "go-microservice"
	if !strings.Contains(body, expected) {
		t.Errorf("hello handler didn't return expected message: %s",
			expected)
	}

}
