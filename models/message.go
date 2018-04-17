package models

import "time"

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
	HteamCH   string `json:"hteam_ch"` // 英文隊名
	HteamTW   string `json:"hteam_tw"` // 繁中隊名
	HteamCN   string `json:"hteam_cn"` // 簡中隊名
	AteamID   uint   `json:"ateam_id"`
	AteamCH   string `json:"ateam_ch"` // 英文隊名
	AteamTW   string `json:"ateam_tw"` // 繁中隊名
	AteamCN   string `json:"ateam_cn"` // 簡中隊名
	StartDate string `json:"start_date"`
	StartTS   int64  `json:"start_ts"`
	StartTime string `json:"start_time"`
	// MatchLive      bool   `json:"match_live"`

	// group 相關
	GroupID        uint     `json:"group_id"`
	GroupNameCh    string   `json:"group_name_ch"` // 英文group名稱
	GroupNameTW    string   `json:"group_name_tw"` // 繁中group名稱
	GroupNameCN    string   `json:"group_name_cn"` // 簡中group名稱
	CategoryID     uint     `json:"category_id"`
	CategoryName   string   `json:"category_name"`
	SportID        uint     `json:"sport_id"`
	OfferIDs       []string `json:"offer_id"`
	TieResult      bool     `json:"tie_result"`
	EnableAsianNew bool     `json:"enable_asian_new"`

	//即時比分
	MatchState  uint   `json:"MatchState,omitempty"` //1:First Half,2:running,3:Second Half
	StateString string `json:"state_string"`
	HomeScore   uint   `json:"home_score,omitempty"`
	AwayScore   uint   `json:"away_score,omitempty"`
	HomeRedcard uint   `json:"HomeRedcard,omitempty"`
	AwayRedcard uint   `json:"AwayRedcard,omitempty"`
	Gametime    uint   `json:"Gametime,omitempty"`
	GameMinute  uint   //比賽不含中場休息＆暫停，開打的時間
	Minute      string //tx 給的型別是字串，比對成功後轉入 GameMinute
}

// SourceOffer 資料來源的 Offer
type SourceOffer struct {
	ID string `json:"Id"`
	// MatchID uint    `json:"match_id"`
	Halves      string  `json:"halves"`
	Bid         uint    `json:"bid"`
	OtID        uint    `json:"ot_id"`
	OtName      string  `json:"otname"`
	HalvesType  string  `json:"halves_type"`
	PlayType    string  `json:"play_type"`
	Head        float64 `json:"head"`
	Hodd        float64 `json:"h_odd"`
	Aodd        float64 `json:"a_odd"`
	Dodd        float64 `json:"d_odd"`
	Hoppo       float64 `json:"h_oppo"`
	Aoppo       float64 `json:"a_oppo"`
	Doppo       float64 `json:"d_oppo"`
	IsRunning   bool    `json:"is_running"`
	PushID      string  `json:"push_id"`
	IsOTB       bool    `json:"is_otb"`
	IsAsians    bool    `json:"is_asians"`
	Proportion  int     `json:"proportion"`
	OfferLineID uint    `json:"offer_lineid"`
	MessageTime
}
type MessageTime struct {
	LastUpdated string `json:"last_updated"`
	Ts          int64  `json:"ts"`
	AdpterTs    int64  `json:"adpter_ts"`
	OfferTs     int64  `json:"offer_ts"`
}
type TxMatch struct {
	StartTime time.Time
	SportID   uint
	HomeID    uint
	AwayID    uint
}
type TxMessage struct {
	Match         string
	OfferOt       string
	OfferLineid   uint
	BookmakerName string
	Line          float64
	HomeOdds      float64
	AwayOdds      float64
	OfferTs       time.Time
}
