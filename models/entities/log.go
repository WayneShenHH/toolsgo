package entities

import "github.com/jinzhu/gorm"

type LogMessage struct {
	gorm.Model
	Data         string  `json:"data"`
	Log          string  `json:"log"`
	TxTimestamp  int64   `json:"tx_timestamp"`
	TxMatchID    uint    `json:"tx_match_id"`
	TxMatch      string  `json:"tx_match"`
	TxOfferID    string  `json:"tx_offer_id"`
	HOppo        float64 `json:"h_oppo"`
	AOppo        float64 `json:"a_oppo"`
	DOppo        float64 `json:"d_oppo"`
	BookMakerID  uint    `json:"book_maker_id"`
	TxOffer      string  `json:"tx_offer"`
	OtName       string  `json:"ot_name"`
	OtType       string  `json:"ot_type"`
	Ot           int     `json:"ot"`
	Head         float64 `json:"head"`
	Action       string  `json:"action"`
	HasError     int     `json:"has_error"`
	Event        string  `json:"event"`
	UUID         string  `json:"uuid"`
	AfuTimestamp string  `json:"afu_timestamp"`
}
