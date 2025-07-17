package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterUserEndpoint(t *testing.T) {
	req := httptest.NewRequest("POST", "/api/register", nil)
	w := httptest.NewRecorder()
	http.NotImplemented(w, req)
	if w.Code != http.StatusNotImplemented {
		t.Fatalf("expected 501 Not Implemented, got %d", w.Code)
	}
}
