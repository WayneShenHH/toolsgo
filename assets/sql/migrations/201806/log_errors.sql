SET @dbname = DATABASE();
SET @tablename = "log_errors";
SET @columnname = "target_id";
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
ALTER TABLE log_errors ADD target_id int unsigned NOT NULL  DEFAULT 0; 
SET @preparedStatement = (IF(
    (SELECT COUNT(*) FROM INFORMATION_SCHEMA.STATISTICS
    WHERE table_name = 'log_errors' AND table_schema = DATABASE() AND index_name = 'idx_err_uni') > 0,
    'ALTER TABLE log_errors DROP INDEX idx_err_uni;',
    'SELECT 1;'
));
PREPARE dropIfExist FROM @preparedStatement;
EXECUTE dropIfExist;
DEALLOCATE PREPARE dropIfExist;
CREATE UNIQUE INDEX idx_err_uni ON log_errors(name, target_id); 