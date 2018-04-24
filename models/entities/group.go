package entities

import "github.com/jinzhu/gorm"

// GroupSource group_sources
type GroupSource struct {
	gorm.Model
	Name            string `gorm:"index:idx_name;unique_index:idx_name_sport_source_leader"`
	NameTW          string
	NameCN          string
	LeaderCountryID uint
	LeaderSportID   uint `gorm:"unique_index:idx_name_sport_source_leader"`
	SportID         uint `sql:"default:null" gorm:"index:idx_sport_id"`
	CategoryID      uint `sql:"default:null" gorm:"index:idx_category_id"`
	GroupID         uint `gorm:"index:idx_group_id"`
	SourceID        uint `gorm:"not null;unique_index:idx_name_sport_source_leader"`
	LeaderID        uint `gorm:"not null;unique_index:idx_name_sport_source_leader"`
}
