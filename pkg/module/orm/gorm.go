// Package orm database orm
package orm

import (
	"fmt"
	"time"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"

	"github.com/WayneShenHH/toolsgo/pkg/environment"
	"github.com/WayneShenHH/toolsgo/pkg/module/logger"
)

const (
	waitTimeout = 28800 // MySQL預設值
	// maxConn     = 200
	// maxIdleConn = 100
)

// New 建立 gorm 連線
func New(config *environment.DatabaseConfig, appLastVersion *gormigrate.Migration) *gorm.DB {
	var db *gorm.DB
	var err error
	connectionString := mysqlArgs(config)
	db, err = gorm.Open("mysql", mysqlArgs(config))
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
	// 檢查 migrate 版本是否一致
	if !CheckVersion(db, appLastVersion) && !config.Migrate {
		panic("Database migrate version not equal")
	}
	return db
}

func mysqlArgs(config *environment.DatabaseConfig) string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=true&loc=UTC&multiStatements=true",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
		config.Encoding)
}
