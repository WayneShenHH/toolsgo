package logsvc

import (
	"encoding/json"
	"fmt"

	"github.com/WayneShenHH/toolsgo/models/entities"
	"github.com/WayneShenHH/toolsgo/repository"
)

// LogService service for log
type LogService struct {
	repository.Repository
}

// New instence JuService
func New(ctx repository.Repository) *LogService {
	return &LogService{
		Repository: ctx,
	}
}

// Read log from redis
func (service *LogService) Read(start int, end int) {
	res := []entities.LogMessage{}
	list := service.Repository.LRange("worker:offer:log", start, end)
	for _, v := range list {
		m := entities.LogMessage{}
		json.Unmarshal(v.([]byte), &m)
		res = append(res, m)
		fmt.Println(m.Log)
	}
}
