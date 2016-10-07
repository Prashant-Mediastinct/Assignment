package models

type AdvertiserData struct {
	Advertiser_id   string         `json:"adertiserId,omitempty"`
	Advertiser_name string         `json:"advertiserName,omitempty"`
	CreativeData    []CreativeData `json:"creativeData,omitempty"`
}
