package entities

import "github.com/jinzhu/gorm"

// TeamSource team_sources
type TeamSource struct {
	gorm.Model
	Name            string `gorm:"index:idx_team_name;"`
	NameTW          string `gorm:"index:idx_team_name_tw;"`
	NameCN          string
	MasterGroupName string `gorm:"unique_index:idx_team_uni;"`
	LeaderID        uint   `sql:"default:null" gorm:"unique_index:idx_team_uni;index:idx_leader_id"`
	LeaderCountryID uint   `sql:"default:null"`
	LeaderSportID   uint   `sql:"default:null"`
	SportID         uint   `sql:"default:null" gorm:"index:idx_sport_id"`
	CategoryID      uint   `sql:"default:null" gorm:"index:idx_category_id"`
	GroupID         uint   `sql:"default:null" gorm:"index:idx_group_id"`
	TeamID          uint   `gorm:"index:idx_system_id" sql:"default: null"`
	SourceID        uint   `gorm:"not null;unique_index:idx_team_uni;"`
}
