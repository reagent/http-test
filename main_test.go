package main

import (
	// "fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func GET(a *App, path string) (*httptest.ResponseRecorder, *http.Request) {
	recorder   := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", path, nil)

	a.Router.ServeHTTP(recorder, request)

	return recorder, request
}

func TestGetReturnsSuccess(t *testing.T) {
	app := newApp()

	r, _ := GET(app, "/")

	if code := r.Code; code != 200 {
		t.Errorf("Expected code of 200 but was %d instead", code)
	}
}

func TestGetReturnsFailure(t *testing.T) {
  app := newApp()

  r, _ := GET(app, "/missing")

  if code := r.Code; code != 404 {
    t.Errorf("Expected code of 404 but was %d instead", code)
  }
}
