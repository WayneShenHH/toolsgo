package services

import (
	"fmt"
	"testing"
	"time"

	"github.com/WayneShenHH/toolsgo/tools/timeutil"

	"github.com/WayneShenHH/toolsgo/repository/repositoryimpl"
	"github.com/WayneShenHH/toolsgo/services/txsvc"
)

func txSvc() *txsvc.TxService {
	ctx := repositoryimpl.New(true)
	return txsvc.New(ctx)
}
func TestArray(t *testing.T) {
	input := 8
	for i := 0; i < input; i++ {
		for j := 0; j < 2*input; j++ {
			if j < (input-i) || j > (input+i) {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Print("\n")
	}
}
func Test_CheckTxSchdule(t *testing.T) {
	service := txSvc()
	n := timeutil.StringToTime("2018-04-30T06:00:00Z")
	s := n.Add(time.Hour * -10)
	e := s.Add(time.Hour * 2)
	matches := service.Repository.GetMatchesByTime(s, e)
	fmt.Println("[CheckTxSchdule] start_time between :", timeutil.TimeToString(s), timeutil.TimeToString(e))
	fmt.Println("[CheckTxSchdule] matches counts:", len(matches))
	for _, v := range matches {
		service.GetTxMsg(v.ID)
	}
}
