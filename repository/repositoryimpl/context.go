package repositoryimpl

import (
	"sync"

	"github.com/WayneShenHH/toolsgo/repository"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"

	// mysql adapter
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
func New() repository.Repository {
	if dbInstance == nil {
		// Use lock to prove only create one dbinstance
		mutex.Lock()
		if dbInstance == nil {
			dbInstance = repository.DBConnect()
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
func DB() *gorm.DB {
	if dbInstance == nil {
		// Use lock to prove only create one dbinstance
		mutex.Lock()
		if dbInstance == nil {
			dbInstance = repository.DBConnect()
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
