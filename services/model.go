package services

// Message struct serves to represent data retrieved from push/offer queue in redis.
type Message struct {
	Content string
	// 時間相關
	MessageTime
	SourceType string `json:"source_type"`
}

// MessageTime 訊息時間戳記
type MessageTime struct {
	// LastUpdated string `json:"last_updated"`
	Ts       int64 `json:"ts"`
	AdpterTs int64 `json:"adpter_ts"`
	OfferTs  int64 `json:"offer_ts"`
}
