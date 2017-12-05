package jusvc

import (
	"fmt"
	"time"

	"github.com/WayneShenHH/toolsgo/tools"
)

func (service *JuService) InsertMessage(key string, jsonfile string) {
	bytes := tools.LoadJson(jsonfile)
	service.Repository.Rpush(key, bytes)
	t := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(fmt.Sprint("[", t, "]insert msg to ", key))
}
func (service *JuService) InsertSpiderOffer(field string, jsonfile string) {
	bytes := tools.LoadJson(jsonfile)
	key := "spider:offers"
	service.Repository.Hset(key, field, bytes)
}
