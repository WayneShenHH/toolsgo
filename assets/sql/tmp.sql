ALTER TABLE matches ADD INDEX idx_start_time (start_time ASC);

ALTER TABLE match_settlements CHANGE COLUMN event_type_id boxscore_type_id INT(10) UNSIGNED NOT NULL DEFAULT '0';
