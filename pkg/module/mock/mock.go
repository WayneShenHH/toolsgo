// Package mock for sql, redis and other modules
package mock

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mssql connect adapter
	"github.com/rafaeljusto/redigomock"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

// NewMock implement the database connection with mock
func NewMock() (*gorm.DB, *redis.Pool, sqlmock.Sqlmock, *redigomock.Conn) {
	db, mysqlMock := sqlMock()
	redis, redisMock := redisMock()
	return db,
		redis,
		mysqlMock,
		redisMock
}

func sqlMock() (*gorm.DB, sqlmock.Sqlmock) {
	mockDB, mysqlMock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open("sqlmock", mockDB)
	if err != nil {
		panic(err)
	}
	return db, mysqlMock
}

func redisMock() (*redis.Pool, *redigomock.Conn) {
	redisMock := redigomock.NewConn()
	redisPool := redis.Pool{
		MaxIdle:   3,
		MaxActive: 3,
		Dial:      func() (redis.Conn, error) { return redisMock, nil },
	}
	return &redisPool, redisMock
}
