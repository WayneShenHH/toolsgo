package services

import (
	"github.com/WayneShenHH/toolsgo/tools"
	"github.com/WayneShenHH/toolsgo/tools/redis"
)

func InsertMessage(key string, jsonfile string) {
	bytes := tools.LoadJson(jsonfile)
	redis.New().Rpush(key, bytes)
}
func InsertSpiderOffer(field string, jsonfile string) {
	bytes := tools.LoadJson(jsonfile)
	key := "spider:offers"
	redis.New().Hset(key, field, bytes)
}
