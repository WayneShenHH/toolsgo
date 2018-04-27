package schedulesvc

import (
	"fmt"
	"time"

	"github.com/WayneShenHH/toolsgo/models"
	"github.com/WayneShenHH/toolsgo/repository"
	"github.com/WayneShenHH/toolsgo/services/txsvc"
	"github.com/WayneShenHH/toolsgo/tools"
	"github.com/WayneShenHH/toolsgo/tools/timeutil"
	"github.com/robfig/cron"
)

// CronService for schedule
type CronService struct {
	repository.Repository
}

// New instence CronService
func New(ctx repository.Repository) *CronService {
	return &CronService{
		Repository: ctx,
	}
}

// Start worker
func (service *CronService) Start() {
	scheduler := cron.New()
	service.timerSchedule(scheduler)
	service.CheckTxTask(scheduler)
	scheduler.Start()
	entries := scheduler.Entries()
	for _, en := range entries {
		fmt.Println(en.Job, en.Next)
	}
	select {} //hang on main process
}
func (service *CronService) timerSchedule(scheduler *cron.Cron) {
	fmt.Println("[scheduletask] schedule a task on the every quarter hour.")
	//run at each hour on 00,15,30,45
	spec := "0 0/15 * * * ?"
	//run at each hour on 01,16,31,46
	// spec = "* 1-59/15 * * * *"
	// spec = "1-59/15 * * * * *"
	scheduler.AddFunc(spec, func() {
		fmt.Println("[timer]running at", time.Now())
	})
}

// ClearDataTask for clear old data
func (service *CronService) ClearDataTask() {
	i := 1
	for {
		fmt.Print(i, " -> ")
		list := service.Repository.GetOldData()
		mso, ms, _ := MergeID(*list)
		cnt := timeutil.GetTimer()
		if len(mso) > 0 {
			service.Repository.ClearOldData(mso, ms)
			cnt.Counting("ClearDataTask")
		} else {
			fmt.Println("[ClearDataTask] No data, sleep 10 seconds.")
			time.Sleep(time.Second * 10)
		}
		i++
	}
}

// MergeID OfferHierarchy into id lists
func MergeID(list []models.OfferHierarchy) ([]uint, []uint, []uint) {
	mso, ms, m := []uint{}, []uint{}, []uint{}
	for _, v := range list {
		if v.MatchSetOfferID > 0 {
			mso = tools.UniAppend(mso, v.MatchSetOfferID)
		}
		if v.MatchSetID > 0 {
			ms = tools.UniAppend(ms, v.MatchSetID)
		}
		if v.MatchID > 0 {
			m = tools.UniAppend(m, v.MatchID)
		}
	}
	return mso, ms, m
}

// CheckTxTask for watching offer is normal
func (service *CronService) CheckTxTask(scheduler *cron.Cron) {
	spec := "0 0 0/2 * * ?"
	fmt.Println("[scheduletask] schedule a task for watching tx on every two hours.")
	scheduler.AddFunc(spec, func() {
		fmt.Println("[CheckTxTask] running at", time.Now().Format(time.Kitchen))
		txSvc := txsvc.New(service.Repository)
		txSvc.CheckTxSchdule(10)
	})
}
