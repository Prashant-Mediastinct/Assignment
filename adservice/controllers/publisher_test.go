package controllers

import (
	//"encoding/json"
	//"github.com/Prashant-Mediastinct/Assignment/adservice/controllers"
	//"github.com/Prashant-Mediastinct/Assignment/models"
	"log"
	"net/http"
	"net/http/httptest"
	//"reflect"
	"testing"
)

func TestHandler(t *testing.T) {
	//var ads []models.AdvertiserData
	//var ad models.AdvertiserData
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	log.Println("Testing for Publisher")

	adunit := "4"
	testURL := "/adunit/" + adunit

	req, err := http.NewRequest("GET", testURL, nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetPublisherId)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	//t.Errorf(rr.Body.String())
	//log.Println(ad.CreativeData)
	//log.Println(rr.Body)

	//log.Println(reflect.TypeOf(ad.CreativeData))
	//err = json.NewDecoder(rr.Body).Decode(&ads)
	//log.Printf("Ads: %+v", ads[1].CreativeData)

	if err != nil {
		t.Error(err.Error())
	}

	log.Println("Publisher testing : SUCCESSFUL")
	//t.Error(ads[0].Advertise_name)
	//t.Errorf()

	//t.Errorf(rr.Body.String())
	// Check the status code is what we expect.
	/*if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}*/
}
