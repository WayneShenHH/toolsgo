package repositoryimpl

import (
	"time"

	"github.com/WayneShenHH/toolsgo/models/entities"
)

func (db *datastore) GetMatchByID(id uint) *entities.Match {
	m := &entities.Match{}
	query := db.mysql.Model(m)
	query.Find(m, id)
	return m
}
func (db *datastore) ClearWorkerData() {
	db.mysql.Exec(`truncate match_sets;
		truncate match_set_offers;
		truncate odds;
		truncate log_closes;
		update matches set available = 0, is_closed = 0, available_time = null;`)
}
func (db *datastore) GetMatchesByTime(start, end time.Time) []entities.Match {
	matches := []entities.Match{}
	db.mysql.Model(&entities.Match{}).Where("start_time > ? AND start_time < ?", start, end).Find(&matches)
	return matches
}
