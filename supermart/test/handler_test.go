package test

import (
	"bytes"
	gohttp "net/http"
	"net/http/httptest"
	"testing"

	"github.com/alka/supermart"
	"github.com/alka/supermart/http"
	"github.com/alka/supermart/store"
)

func TestCreateNewMart(t *testing.T) {
	martStore := store.NewMartStore()
	martService := supermart.NewSuperMartService(store.NewItemStore(), martStore)
	martHandler := http.NewMartHandler(martService, martStore)

	var jsonStr = []byte(`{"name":"FoodWorld"}`)

	req, err := gohttp.NewRequest("POST", "/supermarts", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := gohttp.HandlerFunc(martHandler.CreateNewMart)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != gohttp.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, gohttp.StatusCreated)
	}

	// Check the response body is what we expect.
	//It will fail  just because every time the UUID is creating unique id
	expected := `[{"id": "0322016b-60c9-4d06-a9d3-0b8c34fa3408","name":"FoodWorld"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
