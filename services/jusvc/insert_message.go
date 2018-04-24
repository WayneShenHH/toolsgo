package jusvc

import (
	"fmt"
	"time"

	"github.com/WayneShenHH/toolsgo/tools"
)

// InsertMessage insert a message to redis from json file
func (service *JuService) InsertMessage(key string, jsonfile string) {
	bytes := tools.LoadJSON(jsonfile)
	service.Repository.Rpush(key, bytes)
	t := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(fmt.Sprint("[", t, "]insert msg to ", key))
}

// InsertSpiderOffer insert a message to redis from json file
func (service *JuService) InsertSpiderOffer(field string, jsonfile string) {
	bytes := tools.LoadJSON(jsonfile)
	key := "spider:offers"
	service.Repository.Hset(key, field, bytes)
}
