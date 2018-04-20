package repositoryimpl

import (
	"fmt"
	"time"

	"github.com/WayneShenHH/toolsgo/models"
	"github.com/WayneShenHH/toolsgo/tools/timeutil"
)

func (db *datastore) ClearOdds() {
	st := time.Now().Add(time.Hour * (-48))
	sql := `
	delete from odds 
	where (select start_time from matches where id = odds.match_id) < '%v' limit 10000;`
	fmt.Println("clear data before :", timeutil.TimeToYMD(st))
	db.mysql.Exec(fmt.Sprintf(sql, timeutil.TimeToYMD(st)))
}
func (db *datastore) GetOldData() *[]models.OfferHierarchy {
	st := time.Now().Add(time.Hour * (-48))
	list := &[]models.OfferHierarchy{}
	sql := `select mso.id as match_set_offer_id,mso.match_set_id from match_set_offers mso
	join match_sets s on s.id = mso.match_set_id
	where s.start_time < ? limit 10000;`
	db.mysql.Raw(sql, st).Scan(list)
	return list
}
func (db *datastore) ClearOldData(msoid, msid []uint) {
	sql := `delete from match_set_offers where id in (?);`
	db.mysql.Exec(sql, msoid)
	sql = `delete from odds where match_set_offer_id in (?);`
	db.mysql.Exec(sql, msoid)
	sql = `delete from match_sets where id in (?);`
	db.mysql.Exec(sql, msid)
}
