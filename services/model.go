package services

// Message struct serves to represent data retrieved from push/offer queue in redis.
type Message struct {
	// OldMessage

	// 時間相關
	MessageTime
	SourceType string `json:"source_type"`
	// Bid        int         `json:"bid"`
	Offer SourceOffer `json:"offer"`
	Match SourceMatch `json:"match"`
}

// SourceMatch 資料來源的 match
type SourceMatch struct {
	ID        uint   `json:"id"`
	HteamID   uint   `json:"hteam_id"`
	Hteam     string `json:"hteam"`
	HteamCH   string `json:"hteam_ch"`
	AteamID   uint   `json:"ateam_id"`
	Ateam     string `json:"ateam"`
	AteamCH   string `json:"ateam_ch"`
	StartDate string `json:"start_date"`
	StartTS   int64  `json:"start_ts"`
	StartTime string `json:"start_time"`
	// MatchLive      bool   `json:"match_live"`

	// group 相關
	GroupID        uint     `json:"group_id"`
	GroupNameCh    string   `json:"group_name_ch"`
	CategoryID     uint     `json:"category_id"`
	CategoryName   string   `json:"category_name"`
	SportID        uint     `json:"sport_id"`
	OfferIDs       []string `json:"offer_id"`
	TieResult      bool     `json:"tie_result"`
	EnableAsianNew bool     `json:"enable_asian_new"`
}

// SourceOffer 資料來源的 Offer
type SourceOffer struct {
	ID string `json:"Id"`
	// MatchID uint    `json:"match_id"`
	Halves     string  `json:"halves"`
	Bid        uint    `json:"bid"`
	OtID       uint    `json:"ot_id"`
	OtName     string  `json:"otname"`
	HalvesType string  `json:"halves_type"`
	PlayType   string  `json:"play_type"`
	Head       float64 `json:"head"`
	Hodd       float64 `json:"h_odd"`
	Aodd       float64 `json:"a_odd"`
	Dodd       float64 `json:"d_odd"`
	Hoppo      float64 `json:"h_oppo"`
	Aoppo      float64 `json:"a_oppo"`
	Doppo      float64 `json:"d_oppo"`
	IsRunning  bool    `json:"is_running"`
	PushID     string  `json:"push_id"`
	IsOTB      bool    `json:"is_otb"`
	IsAsians   bool    `json:"is_asians"`
	Proportion int     `json:"proportion"`
	MessageTime
}

// TxData data form source
type TxData struct {
	Type  string // Either "offers" or "matches"
	Match SourceMatch
	Offer SourceOffer
}

// MessageTime 訊息時間戳記
type MessageTime struct {
	// LastUpdated string `json:"last_updated"`
	Ts       int64 `json:"ts"`
	AdpterTs int64 `json:"adpter_ts"`
	OfferTs  int64 `json:"offer_ts"`
}
