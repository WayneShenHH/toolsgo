package repositoryimpl

import (
	"fmt"
	"time"

	"github.com/WayneShenHH/toolsgo/app"
	"github.com/WayneShenHH/toolsgo/repository"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	// mysql adapter
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type datastore struct {
	mysql *gorm.DB
	cache *redisConnext
}

var dbInstance *gorm.DB

// New implement the Store interface with the database connection.
// connect to database and storeage
func New() repository.Repository {
	if dbInstance == nil {
		dbInstance = Connect()
	}
	r := redisConnext{}
	r.Db = redisConnect()
	return &datastore{
		mysql: dbInstance,
		cache: &r,
	}
}

type redisConnext struct {
	Db redis.Conn
}

// Connect 建立 gorm DB 連線
func Connect() *gorm.DB {
	config := app.Configuration()
	db, err := gorm.Open("mysql", config.Mysql)
	if err != nil {
		fmt.Println("Database connection failed.", err.Error())
	}
	fmt.Println("db connect to :" + config.Mysql)
	db.DB().SetMaxIdleConns(200)
	// db.LogMode(true)
	return db
}

// redisConnect 建立 Redis 連線
func redisConnect() redis.Conn {
	config := app.Configuration()
	c, err := redis.Dial("tcp", config.Redis)
	if err != nil {
		fmt.Println("連線失敗，1秒後重新連線。")
		time.Sleep(1000 * time.Millisecond)
		redisConnect()
	}
	return c
}
