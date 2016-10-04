package controllers

import (
	"github.com/Prashant-Mediastinct/Assignment/database"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAdvertiser(t *testing.T) {

	log.Println("Testing for Advertiser")
	Publisher_id := "4"
	testURL := "/publisher_id/" + Publisher_id

	database.DBDef()

	req, err := http.NewRequest("GET", testURL, nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAdvertisers)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if err != nil {
		t.Error(err.Error())
	}

	log.Println("Advertiser testing : SUCCESSFUL!")
}
