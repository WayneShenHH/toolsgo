package schedulesvc

import (
	"fmt"
	"time"

	"github.com/WayneShenHH/toolsgo/repository"
	"github.com/robfig/cron"
)

type CronService struct {
	repository.Repository
}

// New instence CronService
func New(ctx repository.Repository) *CronService {
	return &CronService{
		Repository: ctx,
	}
}

func (service *CronService) Start() {
	scheduler := cron.New()
	service.timerSchedule(scheduler)
	scheduler.Start()
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
func (service *CronService) ClearDataTask() {
	for {
		<-time.After(time.Second * 1)
		service.Repository.ClearOdds()
	}
}
