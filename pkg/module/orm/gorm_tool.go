package orm

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"

	"github.com/WayneShenHH/toolsgo/pkg/module/logger"
)

// DropColumnIfExist check before drop column
func DropColumnIfExist(db *gorm.DB, tableName string, columnName string) error {
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

// IfColumnExists check is column exist
func IfColumnExists(db *gorm.DB, tableName string, columnName string) (bool, error) {
	sql := fmt.Sprintf(`
SET @dbname = DATABASE();
SELECT COUNT(*) AS count FROM INFORMATION_SCHEMA.COLUMNS
WHERE
(table_name = '%v')
AND (table_schema = @dbname)
AND (column_name = '%v')`, tableName, columnName)

	var cnt int
	rows, err := db.Raw(sql).Rows()
	for rows.Next() {
		rows.Scan(&cnt)
		return cnt > 0, nil
	}
	return false, err
}

// CreateViews add sql view
func CreateViews(db *gorm.DB, views []string) {
	for _, v := range views {
		if err := db.Exec(v).Error; err != nil {
			logger.Error(err)
			os.Exit(1)
		}
	}
}

// CheckVersion check database migration version
func CheckVersion(db *gorm.DB, appLastVersion *gormigrate.Migration) bool {
	var dbLast gormigrate.Migration
	db.Last(&dbLast)
	if dbLast.ID == appLastVersion.ID {
		return true
	}
	return false
}
