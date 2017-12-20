ALTER TABLE default_category_maps ADD source_id INT UNSIGNED;
ALTER TABLE default_group_maps ADD source_id INT UNSIGNED;
ALTER TABLE default_team_maps ADD source_id INT UNSIGNED;
update default_category_maps set source_id = 2;
update default_group_maps set source_id = 2;
update default_team_maps set source_id = 2;

ALTER TABLE default_category_maps CHANGE ju_name source_category_name VARCHAR(255);
ALTER TABLE default_category_maps CHANGE ju_name_tw source_category_name_tw VARCHAR(255);
ALTER TABLE default_group_maps CHANGE ju_name source_group_name VARCHAR(255);
ALTER TABLE default_group_maps CHANGE ju_name_tw source_group_name_tw VARCHAR(255);
ALTER TABLE default_team_maps CHANGE ju_name source_team_name VARCHAR(255);
ALTER TABLE default_team_maps CHANGE ju_name_tw source_team_name_tw VARCHAR(255);
ALTER TABLE default_team_maps CHANGE ju_category_name source_category_name VARCHAR(255);

ALTER TABLE default_category_maps
  ADD UNIQUE INDEX idx_cat_map (tx_name,source_category_name,sport_name,source_id);
ALTER TABLE default_group_maps
  ADD UNIQUE INDEX idx_cat_grp (tx_name,source_group_name,sport_name,source_id);
ALTER TABLE default_team_maps
  ADD UNIQUE INDEX idx_team_map (tx_name,tx_category_name,source_id);