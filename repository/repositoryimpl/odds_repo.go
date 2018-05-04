package repositoryimpl

import (
	"time"

	"github.com/WayneShenHH/toolsgo/models"
)

func (db *datastore) GetOldData() *[]models.OfferHierarchy {
	st := time.Now().Add(time.Hour * (-48))
	list := &[]models.OfferHierarchy{}
	sql := `select mso.id as match_set_offer_id,mso.match_set_id from match_set_offers mso
	join match_sets s on s.id = mso.match_set_id
	where s.start_time < ? limit 10000;`
	db.mysql.Raw(sql, st).Scan(list)
	return list
}
func (db *datastore) ClearOldData() {
	st := time.Now().Add(time.Hour * (-48))
	sql := `delete from match_set_offers 
	where (select start_time from match_sets where id=match_set_offers.match_set_id) < ? 
	or (select id from match_sets where id=match_set_offers.match_set_id) is null limit 10000;`
	db.mysql.Exec(sql, st)
	sql = `delete from odds
	where (select start_time from matches where id = odds.match_id) < ? 
	and (select odds_id from order_items i where i.odds_id = odds.id) is null limit 10000;`
	db.mysql.Exec(sql, st)
	sql = `select * from match_sets where start_time < ? limit 10000;`
	db.mysql.Exec(sql, st)
}
