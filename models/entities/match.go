package entities

import (
	"time"

	"github.com/jinzhu/gorm"
)

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
}
