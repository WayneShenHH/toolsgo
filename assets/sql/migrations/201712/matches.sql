ALTER TABLE groups ADD is_special TINYINT(1);
ALTER TABLE group_sources ADD is_special TINYINT(1);
ALTER TABLE match_sources ADD group_id INT(10) UNSIGNED;

ALTER TABLE matches
DROP INDEX idx_teams_start_time;
ALTER TABLE matches
ADD UNIQUE INDEX idx_teams_start_time (start_time,hteam_id,ateam_id,group_id);

ALTER TABLE match_sources
DROP INDEX idx_teams_start_time;
ALTER TABLE match_sources
ADD UNIQUE INDEX idx_teams_start_time (start_time,hteam_id,ateam_id,group_id);
