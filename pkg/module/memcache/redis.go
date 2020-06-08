package memcache

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	json "github.com/json-iterator/go"

	"github.com/WayneShenHH/toolsgo/pkg/environment"
	"github.com/WayneShenHH/toolsgo/pkg/errors"
)

type redisCache struct {
	Pool        *redis.Pool
	tokenExpire uint
}

// New 初始化 Redis Pool 並封裝成 Cache
func New(config *environment.RedisConfig) MemCache {
	return &redisCache{
		Pool:        redisConnect(config),
		tokenExpire: config.JWTExpire * 60,
	}
}

// redisConnect 取得 redis 連線池
func redisConnect(config *environment.RedisConfig) *redis.Pool {
	timeout := redis.DialConnectTimeout(time.Duration(config.Timeout) * time.Millisecond)
	writeTimeout := redis.DialWriteTimeout(time.Duration(config.WriteTimeout) * time.Millisecond)
	readTimeout := redis.DialReadTimeout(time.Duration(config.ReadTimeout) * time.Millisecond)
	return &redis.Pool{
		Wait:        config.Block,
		IdleTimeout: time.Duration(config.IdleTimeout) * time.Millisecond,
		MaxIdle:     config.MaxIdleConns,
		MaxActive:   config.MaxConns, // max number of connections
		Dial: func() (redis.Conn, error) {
			hap := fmt.Sprintf("%v:%v", config.Host, config.Port)
			c, err := redis.Dial("tcp", hap, redis.DialDatabase(config.Index),
				timeout,
				writeTimeout,
				readTimeout,
			)
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

func (r *redisCache) do(command string, args ...interface{}) (interface{}, error) {
	conn := r.Pool.Get()
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()

	return conn.Do(command, args...)
}

// KEYS 取得所有符合 pattern 的鍵
func (r *redisCache) KEYS(pattern string) ([]string, error) {
	res := []string{}
	inter, err := r.do("keys", pattern)
	if err != nil {
		return res, err
	}
	list := inter.([]interface{})
	for _, v := range list {
		res = append(res, string(v.([]uint8)))
	}
	return res, nil
}

// SET 新建一筆 Key:Value
func (r *redisCache) SET(key interface{}, value interface{}, options ...interface{}) (interface{}, error) {
	body, err := json.Marshal(value)
	if err != nil {
		return nil, errors.Str(fmt.Sprintf("memcache/SET json.Marshal error: %v", err))
	}
	if len(options) > 0 {
		return r.do("SET", append([]interface{}{key, string(body)}, options...)...)
	}
	return r.do("SET", key, string(body))
}

// SETToken 透過 SET 新增一筆 Token ，並會為其加上有效時間
func (r *redisCache) SETToken(key interface{}, token interface{}) error {
	_, err := r.do("SET", key, token, "EX", r.tokenExpire)
	return err
}

// GETToken 透過 SET 新增一筆 Token ，並會為其加上有效時間
func (r *redisCache) GETSTR(key interface{}) (string, error) {
	inte, err := r.do("GET", key)
	if inte == nil {
		return "", err
	}
	return string(inte.([]byte)), err
}

func (r *redisCache) SETEX(key interface{}, value interface{}, expireSeconds uint) error {
	j, _ := json.Marshal(value)
	_, err := r.do("SET", key, j, "EX", expireSeconds)
	return err
}

// GET 透過 Key 找出對應之 Value，傳入物件的指標
func (r *redisCache) GET(key, model interface{}) error {
	inte, err := r.do("GET", key)
	if inte == nil {
		return err
	}
	return json.Unmarshal(inte.([]byte), model)
}

//MGET call redis mget
func (r *redisCache) MGET(keys []string) ([][]byte, error) {
	res := make([][]byte, 0)
	if len(keys) == 0 {
		return res, nil
	}
	args := make([]interface{}, len(keys))
	for i, v := range keys {
		args[i] = v
	}
	inter, err := r.do("mget", args...)
	if err != nil {
		return res, err
	}
	values := inter.([]interface{})
	for _, v := range values {
		if v == nil {
			continue
		}
		res = append(res, v.([]byte))
	}
	return res, nil
}

// DEL 刪除指定的 Key 與對應的 Value
func (r *redisCache) DEL(key interface{}) error {
	_, err := r.do("DEL", key)
	return err
}

// DELs 刪除指定的 Key 與對應的 Value
func (r *redisCache) DELs(keys ...interface{}) error {
	_, err := r.do("DEL", keys...)
	return err
}

// HSET 透過 Hash 在指定 Key 中新建一筆 Field:Value
func (r *redisCache) HSET(key, field, value interface{}) (err error) {
	_, err = r.do("HSET", key, field, value)
	return err
}

// HGET 透過 Hash 來取得指定 Key 中 Field 的 Value
func (r *redisCache) HGET(key, field interface{}) (interface{}, error) {
	return r.do("HGET", key, field)
}

// HDEL 透過 Hash 來刪除指定 Key 的 Field 與對應的 Value
func (r *redisCache) HDEL(key, field interface{}) (err error) {
	_, err = r.do("HDEL", key, field)
	return err
}

// EVAL 執行給予的 Script ，並回傳執行完的資料
func (r *redisCache) EVAL(script string, options ...interface{}) (interface{}, error) {
	return r.do("EVAL", append([]interface{}{script}, options...)...)
}

// EXISTS 確認指定的 Key 是否存在
func (r *redisCache) EXISTS(key interface{}) (bool, error) {
	const exist int64 = 1

	reply, err := r.do("EXISTS", key)
	if err != nil {
		return false, err
	}

	return reply == exist, nil
}

// Flush 清空整個快取
func (r *redisCache) Flush() error {
	reply, err := r.do("FLUSHALL")
	if err != nil {
		return err
	}
	if reply != "OK" {
		return errors.Str("cache flush failed")
	}
	return nil
}

// RPush redis RPush
func (r *redisCache) RPush(key string, content interface{}) (err error) {
	bytes, err := json.Marshal(content)
	if err != nil {
		return
	}
	_, err = r.do("RPUSH", key, bytes)
	return
}
