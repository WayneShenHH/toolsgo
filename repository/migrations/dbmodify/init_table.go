package dbmodify

import (
	"github.com/WayneShenHH/toolsgo/models/entities"
	gormigrate "github.com/go-gormigrate/gormigrate"
	"github.com/jinzhu/gorm"
)

var initTable = &gormigrate.Migration{
	ID: "201709132222",
	Migrate: func(db *gorm.DB) error {
		// it's a good pratice to copy the struct inside the function,
		// so side effects are prevented if the original struct changes during the time
		return db.AutoMigrate(
			// &entities.Match{},
			// &entities.CategorySource{},
			// &entities.GroupSource{},
			// &entities.TeamSource{},
			// &entities.User{},
			&entities.ClockIn{},
		).Error
	},
	Rollback: func(db *gorm.DB) error {
		return db.DropTable("people").Error
	},
}
