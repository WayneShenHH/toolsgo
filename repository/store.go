package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/WayneShenHH/toolsgo/app"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
)

// Key is the key name of the store in the Gin context.
const Key = "store"
const (
	waitTimeout = 28800 // MySQL預設值
	maxConn     = 200
	maxIdleConn = 100
)

// DBConnect 取得 gorm DB 連線池
func DBConnect(logMode bool) *gorm.DB {
	config := app.Configuration()
	db, err := gorm.Open("mysql", config.Mysql)
	if err != nil {
		panic("Database connection failed.")
	}
	db.DB().SetConnMaxLifetime(time.Second * waitTimeout / 2)
	db.DB().SetMaxIdleConns(maxIdleConn)
	db.DB().SetMaxOpenConns(maxConn)
	db.LogMode(logMode)
	fmt.Printf("Connect to MySQL %s successful\n", config.Mysql)
	return db
}

// RedisConnect 取得 redis 連線池
func RedisConnect() *redis.Pool {
	config := app.Configuration()
	return &redis.Pool{
		MaxIdle:   30,
		MaxActive: 1000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.Redis, redis.DialDatabase(config.RedisDatabaseIndex))
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

// Setter defines a context that enables setting values.
type Setter interface {
	Set(string, interface{})
}

// FromContext returns the Store associated with this context.
func FromContext(c context.Context) Repository {
	return c.Value(Key).(Repository)
}

// ToContext adds the Store to this context if it supports
// the Setter interface.
func ToContext(c Setter, repo Repository) {
	c.Set(Key, repo)
}
