package repositoryimpl

import (
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
		truncate book_maker_sequences;
		truncate auto_available_messages;
		truncate offer_stats;
		update matches set available = 0, available_time = null;`)
}
