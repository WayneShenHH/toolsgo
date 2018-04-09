package txsvc

import (
	"fmt"
	"os"

	"github.com/WayneShenHH/toolsgo/repository"
	"github.com/WayneShenHH/toolsgo/tools/timeutil"
)

type TxService struct {
	repository.Repository
}

// New instence JuService
func New(ctx repository.Repository) *TxService {
	return &TxService{
		Repository: ctx,
	}
}

func (service *TxService) GetTxMsg(mid uint) {
	list := service.Repository.TxMessage(mid)
	data := "match_id,offer_lineid,bookmaker,line,home_odds,away_odds,offer_ts\n"
	for _, v := range list {
		lineid := fmt.Sprint(v.OfferLineid)
		line := fmt.Sprint(v.Line)
		h := fmt.Sprint(v.HomeOdds)
		a := fmt.Sprint(v.AwayOdds)
		t := timeutil.TimeToString(v.OfferTs)
		// out := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t\n", lineid, v.BookmakerName, line, h, a, t[:19])
		out := fmt.Sprintf("%d,%v,%v,%v,%v,%v,%v\n", mid, lineid, v.BookmakerName, line, h, a, t[:19])
		data += out
		fmt.Println(out)
	}
	OutPutCsv(data)
	fmt.Println("count:", len(list))
}
func OutPutCsv(data string) {
	path := "./result.csv"
	os.Remove(path)
	file, _ := os.Create(path)
	defer file.Close()
	file.WriteString(data)
}
