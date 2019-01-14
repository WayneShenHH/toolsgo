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
	db.mysql.Exec(`
		truncate match_set_offers;
		truncate log_pushes;
		truncate log_closes;
		truncate log_operators;
		truncate log_errors;
		update matches set available=0, status=2, is_closed=0,is_running=0, available_time = null, auto_enable_time = null, enable=0;
		update matches set start_time=date_format(date_add(now(),interval 1 day),'%Y-%m-%d 01:30:00'), 
		account_time=date_format(date_add(now(),interval 1 day),'%Y-%m-%d') where id=1;`)
}
func (db *datastore) GetMatchesByTime(start, end time.Time) []entities.Match {
	matches := []entities.Match{}
	db.mysql.Model(&entities.Match{}).Where("start_time > ? AND start_time < ?", start, end).Find(&matches)
	return matches
}
func (db *datastore) GetSourceMatchByStruct(filter entities.MatchSource) *entities.MatchSource {
	m := &entities.MatchSource{}
	db.mysql.Where(filter).Find(m)
	return m
}
