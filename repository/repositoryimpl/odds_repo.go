package repositoryimpl

import (
	"fmt"
	"time"

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
