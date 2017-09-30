package services

import (
	"wayne.sdk/wayne/toolsgo/tools"
	"wayne.sdk/wayne/toolsgo/tools/redis"
)

func InsertMessage(key string, jsonfile string) {
	bytes := tools.LoadJson(jsonfile)
	redis.New().Rpush(key, bytes)
}
