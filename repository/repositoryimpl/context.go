package repositoryimpl

import (
	"sync"

	"github.com/WayneShenHH/toolsgo/repository"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	// mysql adapter
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	waitTimeout = 28800 // MySQL預設值
	maxConn     = 200
	maxIdleConn = 100
)

type datastore struct {
	mysql *gorm.DB
	cache *redis.Pool
}

var (
	dbInstance    *gorm.DB
	redisInstance *redis.Pool
	mutex         sync.Mutex
)

// New implement the Store interface with the database connection.
// connect to database and storeage
func New(logMode bool) repository.Repository {
	if dbInstance == nil {
		// Use lock to prove only create one dbinstance
		mutex.Lock()
		if dbInstance == nil {
			dbInstance = repository.DBConnect(logMode)
		}
		mutex.Unlock()
	}
	if redisInstance == nil {
		mutex.Lock()
		if redisInstance == nil {
			redisInstance = repository.RedisConnect()
		}
		mutex.Unlock()
	}
	return &datastore{
		mysql: dbInstance,
		cache: redisInstance,
	}
}

type redisConnext struct {
	Db redis.Conn
}

// DB Get database connection
func DB(logMode bool) *gorm.DB {
	if dbInstance == nil {
		// Use lock to prove only create one dbinstance
		mutex.Lock()
		if dbInstance == nil {
			dbInstance = repository.DBConnect(logMode)
		}
		mutex.Unlock()
	}
	return dbInstance
}

// Redis Get redis connection
func Redis() redis.Conn {
	if redisInstance == nil {
		// Use lock to prove only create one dbinstance
		mutex.Lock()
		if redisInstance == nil {
			redisInstance = repository.RedisConnect()
		}
		mutex.Unlock()
	}
	return redisInstance.Get()
}
