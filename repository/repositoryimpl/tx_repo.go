package repositoryimpl

import (
	"fmt"

	"github.com/WayneShenHH/toolsgo/app"
	"github.com/WayneShenHH/toolsgo/models"
)

func (db *datastore) TxMessage(mid uint) []models.TxMessage {
	db.mysql.LogMode(true)
	match := models.TxMatch{}
	db.mysql.Raw(`select s.leader_id,s.match_id from match_sources s
		where s.match_id = ? and s.source_id = 1`, mid).Scan(&match)
	list := []models.TxMessage{}
	env := app.Env()
	if env != "development" {
		env = "_" + env
	} else {
		env = ""
	}
	// sportMap := map[uint][]uint{
	// 	1: []uint{33},
	// 	2: []uint{59},
	// 	3: []uint{59},
	// 	4: []uint{59},
	// 	5: []uint{10},
	// 	6: []uint{59},
	// 	7: []uint{59},
	// }
	// otid := sportMap[match.SportID]
	sql := fmt.Sprintf(`select 
		m.match,m.offer_ot,
		m.offer_lineid,
		m.bookmaker_name,
		m.cls as line,m.price_oh as home_odds,m.price_oa as away_odds,
		FROM_UNIXTIME(m.offer_ts/1000) as offer_ts 
		from txgo%v.price_updates m where 1
		and match_txid = ? 
		and offer_inrunning = 1
		#and offer_otid in (?)
		order by 
		m.bookmaker_name,m.cls,m.offer_ts`, env)
	db.mysql.Raw(sql, match.LeaderID).Scan(&list)
	return list
}
