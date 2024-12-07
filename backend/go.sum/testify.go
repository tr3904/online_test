package handlers_test

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestCreateProduct(t *testing.T) {
	req, _ := http.NewRequest("POST", "/products", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateProduct)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}