package entities

import (
	"time"

	"github.com/jinzhu/gorm"
)

// ClockIn log for clock in
type ClockIn struct {
	gorm.Model
	ClockInAt time.Time
}
