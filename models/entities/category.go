package entities

import "github.com/jinzhu/gorm"

type CategorySource struct {
	gorm.Model
	Name            string `gorm:"not null;index:idx_name;unique_index:idx_name_source_leader"`
	NameTW          string `sql:"default:null"`
	NameCN          string `sql:"default:null"`
	LeaderCountryID uint   `sql:"default:null"`
	LeaderSportID   uint   `sql:"default:null"`
	SportID         uint   `sql:"default:null" gorm:"index:idx_sport_id"`
	CategoryID      uint   `sql:"default:null" gorm:"index:idx_category_id"`
	SourceID        uint   `gorm:"not null;unique_index:idx_name_source_leader"`
	LeaderID        uint   `gorm:"not null;unique_index:idx_name_source_leader"`
}
