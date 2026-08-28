package httpapi

import (
	"net/http/httptest"
	"streetlight40/workflow"
	"testing"
)

func TestHTTPHandler(t *testing.T) {
	h := New(&workflow.Service{}).Handler()
	r := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if w.Code != 204 {
		t.Fatal(w.Code)
	}
}
