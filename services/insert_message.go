package services

import (
	"fmt"
	"time"

	"github.com/WayneShenHH/toolsgo/tools"
	"github.com/WayneShenHH/toolsgo/tools/redis"
)

func InsertMessage(key string, jsonfile string) {
	bytes := tools.LoadJson(jsonfile)
	redis.New().Rpush(key, bytes)
	t := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("[", t, "]insert msg")
}
func InsertSpiderOffer(field string, jsonfile string) {
	bytes := tools.LoadJson(jsonfile)
	key := "spider:offers"
	redis.New().Hset(key, field, bytes)
}
