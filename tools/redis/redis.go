package redis

import (
	"github.com/garyburd/redigo/redis"
)

type Redis struct {
	Conn redis.Conn
}

func New() *Redis {
	r := Redis{}
	r.Conn = connect()
	return &r
}
func connect() redis.Conn {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		// handle error
	}
	return c
}

// Rpush call redis Rpush
func (r *Redis) Rpush(key string, value []byte) {
	_, err := r.Conn.Do("rpush", key, value)
	if err != nil {
		panic(err)
	}
}
