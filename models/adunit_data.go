package models

type AdunitData struct {
	Adunit_id   string     `json:"adunitId,omitempty"`
	Adunit_info AdunitInfo `json:"adunitInfo,omitempty"`
}

type AdunitInfo struct {
	Advertiser []AdvertiserData `json:"advertiserData,omitempty"`
}
