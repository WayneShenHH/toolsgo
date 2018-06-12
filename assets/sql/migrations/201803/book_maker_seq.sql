ALTER TABLE book_maker_sequences DROP COLUMN play_type_code;
ALTER TABLE book_maker_sequences ADD default_book_maker_sequence_id INT(10) UNSIGNED;
ALTER TABLE book_maker_sequences ADD pregame_hours INT(10) UNSIGNED;
ALTER TABLE book_maker_sequences DROP COLUMN pregame_hours;

ALTER TABLE default_book_maker_sequences ADD sport_id INT(10) UNSIGNED;
ALTER TABLE default_book_maker_sequences ADD category_id INT(10) UNSIGNED;
ALTER TABLE default_book_maker_sequences ADD group_id INT(10) UNSIGNED;
ALTER TABLE default_book_maker_sequences ADD is_running TINYINT(1);
ALTER TABLE default_book_maker_sequences ADD pregame_hours INT(10) UNSIGNED;
ALTER TABLE default_book_maker_sequences ADD forbidden_asians TINYINT(1);

ALTER TABLE default_book_maker_sequences DROP COLUMN pregame_hours ;
ALTER TABLE default_book_maker_sequences DROP COLUMN forbidden_asians;

-- add column if exist
SET @col = (SELECT column_name FROM INFORMATION_SCHEMA.COLUMNS WHERE  TABLE_NAME='match_set_offers' AND column_name='offer_ts');
SET @q = (IF(ISNULL(@col),'ALTER TABLE `match_set_offers` ADD `offer_ts` bigint NOT NULL  DEFAULT 0;','select 1;'));
PREPARE stmt FROM @q;
EXECUTE stmt;