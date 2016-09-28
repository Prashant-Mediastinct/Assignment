package models

type AdvertiserData struct {
	Advertiser_id  int             `json:"adertiserId,omitempty"`
	Advertise_name string          `json:"advertiserName,omitempty"`
	CreativeData   *[]CreativeData `json:"creativeData,omitempty"`
}
