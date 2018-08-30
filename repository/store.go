package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/WayneShenHH/toolsgo/app"
	"github.com/WayneShenHH/toolsgo/module/logger"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
)

const (
	waitTimeout = 28800 // MySQL預設值
	// maxConn     = 200
	// maxIdleConn = 100
)

// Context 外部資源
type Context struct {
	// DbContext *gorm.DB
	// Redis      *redis.Redis
	Repository Repository
}

// DBConnect 取得 gorm DB 連線池
func DBConnect() *gorm.DB {
	config := app.Setting.Database
	connectionString := mysqlArgs()
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		logger.Debug(connectionString)
		logger.Error(err.Error())
		panic("Database connection failed.")
	}
	db.DB().SetConnMaxLifetime(time.Second * waitTimeout / 2)
	db.DB().SetMaxIdleConns(config.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.MaxConns)
	db.LogMode(config.LogMode)
	logger.Debug(fmt.Sprintf("Connect to MySQL %s successful", connectionString))
	return db
}
func mysqlArgs() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=true&loc=UTC&multiStatements=true",
		app.Setting.Database.Username,
		app.Setting.Database.Password,
		app.Setting.Database.Host,
		app.Setting.Database.Port,
		app.Setting.Database.Name,
		app.Setting.Database.Encoding)
}

// RedisConnect 取得 redis 連線池
func RedisConnect() *redis.Pool {
	config := app.Setting.Redis
	return &redis.Pool{
		MaxIdle:   config.MaxIdleConns,
		MaxActive: config.MaxConns, // max number of connections
		Dial: func() (redis.Conn, error) {
			hap := fmt.Sprintf("%v:%v", config.Host, config.Port)
			c, err := redis.Dial("tcp", hap, redis.DialDatabase(config.Index))
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

// Key is the key name of the store in the Gin context.
const Key = "store"

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
func ToContext(c Setter, store Repository) {
	c.Set(Key, store)
}
