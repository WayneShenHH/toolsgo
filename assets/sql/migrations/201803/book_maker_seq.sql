ALTER TABLE book_maker_sequences DROP COLUMN play_type_code;
ALTER TABLE book_maker_sequences ADD default_book_maker_sequence_id INT(10) UNSIGNED;
ALTER TABLE book_maker_sequences ADD pregame_hours INT(10) UNSIGNED;

ALTER TABLE default_book_maker_sequences ADD sport_id INT(10) UNSIGNED;
ALTER TABLE default_book_maker_sequences ADD category_id INT(10) UNSIGNED;
ALTER TABLE default_book_maker_sequences ADD group_id INT(10) UNSIGNED;
ALTER TABLE default_book_maker_sequences ADD is_running TINYINT(1);
ALTER TABLE default_book_maker_sequences ADD pregame_hours INT(10) UNSIGNED;