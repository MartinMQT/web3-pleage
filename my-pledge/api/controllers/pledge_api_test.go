package controllers


import (
	"net/http"
	"net/http/httptest"
	"testing"
)
func TestPledgeHandler(t *testing.T) {
	req, err := http.NewRequest("POST", "/api/pledge", nil)
	if err != nil {
			t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PledgeHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Pledge successful"
	if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}