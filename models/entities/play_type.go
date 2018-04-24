package entities

import "github.com/jinzhu/gorm"

// PlayType play_types
type PlayType struct {
	gorm.Model
	Name        string `sql:"default:null"`
	Description string `sql:"default:null"`
	Code        string `gorm:"not null;required;unique_index:idx_name_code"`
	IsRunning   bool   `gorm:"not null;required;unique_index:idx_name_code"`
	IsParlay    bool   `gorm:"not null;required;unique_index:idx_name_code"`
}
