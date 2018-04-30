package txsvc

import (
	"fmt"
	"os"
	"time"

	"github.com/WayneShenHH/toolsgo/models"
	"github.com/WayneShenHH/toolsgo/repository"
	"github.com/WayneShenHH/toolsgo/tools/timeutil"
)

// TxService for deal with tx data
type TxService struct {
	repository.Repository
}

// New instence JuService
func New(ctx repository.Repository) *TxService {
	return &TxService{
		Repository: ctx,
	}
}

// GetTxMsg get tx source data for a match
func (service *TxService) GetTxMsg(mid uint) {
	list := service.Repository.TxMessage(mid)
	if len(list) == 0 {
		return
	}
	data := "match_id,offer_lineid,bookmaker,line,home_odds,away_odds,offer_ts\n"
	dataAbnormal := data
	lastItem := list[0]
	cnt := 0
	for _, v := range list {
		current := rowToString(v)
		data += current
		if lastItem.OfferTs.Add(time.Minute * 3).Before(v.OfferTs) {
			dataAbnormal += rowToString(lastItem)
			dataAbnormal += current
			cnt++
		}
		lastItem = v
		fmt.Println(current)
	}
	OutPutCsv(fmt.Sprint(mid, "_result"), data)
	OutPutCsv(fmt.Sprint(mid, "_result_abnormal"), dataAbnormal)
	fmt.Println("count:", len(list), "abnormal:", cnt)
}
func rowToString(row models.TxMessage) string {
	lineid := fmt.Sprint(row.OfferLineid)
	line := fmt.Sprint(row.Line)
	h := fmt.Sprint(row.HomeOdds)
	a := fmt.Sprint(row.AwayOdds)
	t := timeutil.TimeToString(row.OfferTs)
	// out := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t\n", lineid, v.BookmakerName, line, h, a, t[:19])
	out := fmt.Sprintf("%d,%v,%v,%v,%v,%v,%v\n", row.LeaderID, lineid, row.BookmakerName, line, h, a, t[:19])
	return out
}

// OutPutCsv out put csv file
func OutPutCsv(fileName, data string) {
	path := fmt.Sprintf("./assets/csv/%v.csv", fileName)
	os.Remove(path)
	file, _ := os.Create(path)
	defer file.Close()
	file.WriteString(data)
}

// CheckTxSchdule for watching offer is normal
func (service *TxService) CheckTxSchdule(offsetHour time.Duration) {
	s := time.Now().Add(time.Hour * -offsetHour)
	e := s.Add(time.Hour * 2)
	matches := service.Repository.GetMatchesByTime(s, e)
	fmt.Println("[CheckTxSchdule] start_time between :", timeutil.TimeToString(s), timeutil.TimeToString(e))
	fmt.Println("[CheckTxSchdule] matches counts:", len(matches))
	for _, v := range matches {
		service.GetTxMsg(v.ID)
	}
}
