package repositoryimpl

import (
	"fmt"

	"github.com/WayneShenHH/toolsgo/app"
	"github.com/WayneShenHH/toolsgo/models"
	"github.com/WayneShenHH/toolsgo/tools/timeutil"
)

func (db *datastore) TxMessage(mid uint) []models.TxMessage {
	match := models.TxMatch{}
	db.mysql.Raw(`select m.sport_id,m.start_time,h.leader_id as home_id,a.leader_id as away_id from matches m
	join team_sources h on h.team_id = m.hteam_id and h.leader_id <> 0 and h.source_id = 1
	join team_sources a on a.team_id = m.ateam_id and a.leader_id <> 0 and a.source_id = 1
	where m.id = ?`, mid).Scan(&match)
	list := []models.TxMessage{}
	startTs := timeutil.TimeToStamp(match.StartTime)
	env := app.Env()
	if env != "development" {
		env = "_" + env
	} else {
		env = ""
	}
	sportMap := map[uint][]uint{
		1: []uint{33},
		2: []uint{59},
		3: []uint{59},
		4: []uint{59},
		5: []uint{10},
		6: []uint{59},
		7: []uint{59},
	}
	otid := sportMap[match.SportID]
	sql := fmt.Sprintf(`select 
		m.match,m.offer_ot,
		m.offer_lineid,
		m.bookmaker_name,
		m.cls as line,m.price_oh as home_odds,m.price_oa as away_odds,
		FROM_UNIXTIME(m.offer_ts/1000) as offer_ts 
		from txgo%v.messages m where 1
		and hteam_id = ? 
		and ateam_id = ? 
		and match_time = ?
		and offer_inrunning = 1
		and offer_otid in (?)
		order by 
		m.bookmaker_name,m.offer_ts`, env)
	db.mysql.Raw(sql, match.HomeID, match.AwayID, startTs, otid).Scan(&list)
	return list
}
