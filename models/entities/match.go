package entities

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Match matches
type Match struct {
	gorm.Model
	StartTime         time.Time `gorm:"unique_index:idx_teams_start_time"`
	HteamID           uint      `gorm:"required;not null;unique_index:idx_teams_start_time"`
	AteamID           uint      `gorm:"required;not null;unique_index:idx_teams_start_time"`
	CategoryID        uint      `gorm:"required;"`
	GroupID           uint      `gorm:"required"`
	SportID           uint      `gorm:"required;"`
	MatchAmountLimit  uint      // 下注限額限制
	Available         bool
	StartingPitcherID uint // 棒球先發投手資訊
	IsClosed          bool // 比賽是否已結束，結束條件參考./store/datastore/match.go GetMatchSetForClose()
	// 即時比分
	MatchState uint //1:First Half,2:running,3:Second Half
	HomeScore  uint
	AwayScore  uint
	GameMinute uint //比賽不含中場休息＆暫停，開打經過的時間
}

/*MatchSource 資料源 match 資料
比賽資料參考的資料來源
*/
type MatchSource struct {
	gorm.Model
	LeaderID   uint      `gorm:"index:idx_leader_id"`
	StartTime  time.Time `gorm:"unique_index:idx_teams_start_time"`
	HteamID    uint      `gorm:"not null;unique_index:idx_teams_start_time"`
	AteamID    uint      `gorm:"not null;unique_index:idx_teams_start_time"`
	CategoryID uint      `gorm:"not null;"`
	GroupID    uint      `gorm:"not null;unique_index:idx_teams_start_time"`
	SportID    uint      `gorm:"not null;"`
	MatchID    uint
	SourceID   uint
}
