package database

import "testing"

func TestFetchPublishers(t *testing.T) {

	params := "7"

	expected := "10"

	DBDef()

	Publisher_Id := FetchPublisher(params)

	if Publisher_Id != expected {
		t.Errorf("Publisher_id : got %s want %s", Publisher_Id, expected)
	}
}
