package repositoryimpl

import (
	"time"

	"github.com/WayneShenHH/toolsgo/models/entities"
)

func (db *datastore) ClockIn() {
	db.mysql.Create(&entities.ClockIn{ClockInAt: time.Now()})
}
