alter table category_sources drop column leader_sport_id;
alter table category_sources drop column leader_country_id;
alter table group_sources drop column leader_sport_id;
alter table group_sources drop column leader_country_id;
alter table team_sources drop column leader_sport_id;
alter table team_sources drop column group_id;
alter table team_sources drop column leader_country_id;