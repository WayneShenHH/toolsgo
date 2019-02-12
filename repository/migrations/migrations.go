package migrations

import (
	"fmt"
	"log"

	"github.com/WayneShenHH/toolsgo/repository/migrations/dbviews"

	gormigrate "github.com/go-gormigrate/gormigrate"
	// mysql adapter
	"github.com/WayneShenHH/toolsgo/app"
	"github.com/WayneShenHH/toolsgo/repository"
	"github.com/WayneShenHH/toolsgo/repository/migrations/triggers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // import the mysql driver
)

// DeleteAll db remove
func DeleteAll() error {
	var err error
	db := repository.DBConnect()
	err = DropAllTable(db)
	DropAllView(db)
	return err
}

// Migrate db schema
func Migrate() error {
	var err error
	db := repository.DBConnect()
	if err != nil {
		log.Fatalf("DropAllTable failed: %v", err)
		return err
	}
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		// create persons table
		&initTable,
	})

	if err = m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
		return err
	}

	InitView(db)
	if err != nil {
		log.Fatalf("InitTrigger failed: %v", err)
		return err
	}

	log.Printf("migrations did run successfully")
	return err
}

// DropAllTable 清掉資料庫全部 table 注意使用
func DropAllTable(db *gorm.DB) error {
	q := db.Exec(`
			SET FOREIGN_KEY_CHECKS = 0;
 			SET GROUP_CONCAT_MAX_LEN=32768;
			SET @tables = NULL;
			SELECT GROUP_CONCAT(` + "'`'" + `,table_name,` + "'`'" + `) INTO @tables
			FROM information_schema.tables
			WHERE table_schema = (SELECT DATABASE());
			SELECT IFNULL(@tables,'dummy') INTO @tables;

			SET @tables = CONCAT('DROP TABLE IF EXISTS ', @tables);
			PREPARE stmt FROM @tables;
			EXECUTE stmt;
			DEALLOCATE PREPARE stmt;
			SET FOREIGN_KEY_CHECKS = 1;
			`)
	return q.Error
}

// DropAllView drop all views
func DropAllView(db *gorm.DB) {
	sql := `SET @views = NULL;
	SELECT GROUP_CONCAT(table_schema, '.', table_name) INTO @views
	 FROM information_schema.views
	 WHERE  table_schema = (SELECT DATABASE());
	SET @views = IFNULL(CONCAT('DROP VIEW ', @views), 'SELECT "No Views"');
	PREPARE stmt FROM @views;
	EXECUTE stmt;
	DEALLOCATE PREPARE stmt;`
	db.Exec(sql)
}

// InitTrigger 初始化Trigger
func InitTrigger(db *gorm.DB) {
	for _, v := range triggers.Init() {
		if err := db.Exec(v).Error; err != nil {
			panic(err)
		}
	}
}

// InitView 初始化View
func InitView(db *gorm.DB) {
	for _, v := range dbviews.Init() {
		if err := db.Exec(v).Error; err != nil {
			panic(err)
		}
	}
}

func dropColumnIfExist(db *gorm.DB, tableName string, columnName string) error {
	err := db.Exec(`
		SET @dbname = DATABASE();
		SET @tablename = "` + tableName + `";
		SET @columnname = "` + columnName + `";
		SET @preparedStatement = (SELECT IF(
		(
			SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
			WHERE
			(table_name = @tablename)
			AND (table_schema = @dbname)
			AND (column_name = @columnname)
		) > 0,

		CONCAT("ALTER TABLE ", @tablename, " DROP column ", @columnname),
		"SELECT 1"
		));
		PREPARE alterIfNotExists FROM @preparedStatement;
		EXECUTE alterIfNotExists;
		DEALLOCATE PREPARE alterIfNotExists;
		`).Error
	return err
}
func dropIndexIfExist(db *gorm.DB, tableName string, indexName string) error {
	sql := `
	SET @preparedStatement = (IF(
		(SELECT COUNT(*) FROM INFORMATION_SCHEMA.STATISTICS
		WHERE table_name = '%v' AND table_schema = DATABASE() AND index_name = '%v') > 0,
		'ALTER TABLE %v DROP INDEX %v;',
		'SELECT 1;'
	));
	PREPARE dropIfExist FROM @preparedStatement;
	EXECUTE dropIfExist;
	DEALLOCATE PREPARE dropIfExist;`
	stmt := fmt.Sprintf(sql, tableName, indexName, tableName, indexName)
	err := db.Exec(stmt).Error
	return err
}
func ifColumnExist(db *gorm.DB, tableName string, columnname string) bool {
	dbConfig := app.Setting.Database
	columnCount := 0
	dbErr := db.Table(`INFORMATION_SCHEMA.COLUMNS`).
		Where(`table_schema = ? 
			AND table_name = ? 
			AND column_name = ?`,
			dbConfig.Name, tableName, columnname).Count(&columnCount)

	if dbErr.Error == nil {
		// count > 0 代表欄位存在
		return columnCount > 0
	}
	return false
}
