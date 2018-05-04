package migrations

import (
	"log"

	gormigrate "github.com/go-gormigrate/gormigrate"
	// mysql adapter
	"github.com/WayneShenHH/toolsgo/repository"
	"github.com/WayneShenHH/toolsgo/repository/migrations/triggers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // import the mysql driver
)

// Migrate db schema
func Migrate() error {
	var err error
	db := repository.DBConnect(true)
	err = DropAllTable(db)
	if err != nil {
		log.Fatalf("DropAllTable failed: %v", err)
		return err
	}
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		// create persons table
		&v201709132222,
	})

	if err = m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
		return err
	}

	err = InitTrigger(db)
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

// InitTrigger 初始化Trigger
func InitTrigger(db *gorm.DB) error {
	q := db.Exec(triggers.UpdateMatchSetOffer)
	return q.Error
}
