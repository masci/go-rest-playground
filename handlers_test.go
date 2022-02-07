package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	s "github.com/masci/go-rest-playground/storage"
)

/*
	NOTE: I haven't tested all the handlers because that would be mostly repeating
	boilerplate. Here you can find two use cases from which all the other tests
	can be easily derived: one handler requiring access to the router context (TestGetClass)
	and another one that doesn't (TestListClasses)
*/

func TestMain(m *testing.M) {
	storage = s.NewSqliteStorage(":memory:")
	os.Exit(m.Run())
}

func TestListClasses(t *testing.T) {
	req, err := http.NewRequest("GET", "/classes", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ListClasses)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	want := `[{"ID":"PI0001","name":"Pilates","start_date":"2020-01-29T00:00:00Z","end_date":"2020-02-28T00:00:00Z","capacity":20},{"ID":"DA0001","name":"Dance+","start_date":"2020-01-29T00:00:00Z","end_date":"2020-02-28T00:00:00Z","capacity":20},{"ID":"FB0001","name":"Full Body","start_date":"2020-01-29T00:00:00Z","end_date":"2020-02-28T00:00:00Z","capacity":20},{"ID":"YO0001","name":"Yoga","start_date":"2020-01-29T00:00:00Z","end_date":"2020-02-28T00:00:00Z","capacity":20}]`
	result := strings.TrimSpace(rr.Body.String())
	if result != want {
		t.Errorf("body: got %v want %v", result, want)
	}
}

func TestGetClass(t *testing.T) {
	req, err := http.NewRequest("GET", "/classes", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Since we use a router context to get the param from the URL,
	// we need to replicate that behaviour  in the tests, we can't
	// just do req, err := http.NewRequest("GET", "/classes/PI0001", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("classID", "PI0001")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	// From here, it's like any other httptest logic
	rr := httptest.NewRecorder()
	handler := ClassCtx(http.HandlerFunc(GetClass))
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	want := `{"ID":"PI0001","name":"Pilates","start_date":"2020-01-29T00:00:00Z","end_date":"2020-02-28T00:00:00Z","capacity":20}`
	result := strings.TrimSpace(rr.Body.String())
	if result != want {
		t.Errorf("body: got %v want %v", rr.Body.String(), want)
	}
}
