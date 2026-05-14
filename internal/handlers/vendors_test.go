package handlers

import (
	"net/http/httptest"
	"testing"
)

func TestListVendors(t *testing.T) {
	rec := httptest.NewRecorder()
	ListVendors(rec, httptest.NewRequest("GET", "/api/vendors", nil))
	if rec.Code != 200 {
		t.Fatalf("status %d", rec.Code)
	}
}
