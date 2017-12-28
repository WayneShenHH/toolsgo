-- truncate source maps on staging
truncate sbodds_staging.default_team_maps;
truncate sbodds_staging.default_group_maps;
truncate sbodds_staging.default_category_maps;

-- remove old source data on staging
delete from sbodds_staging.team_sources where source_id <> 1;
delete from sbodds_staging.group_sources where source_id <> 1;
delete from sbodds_staging.category_sources where source_id <> 1;

-- sync team_source from production into staging
insert ignore into sbodds_staging.default_team_maps (created_at,updated_at,tx_name,source_team_name,source_team_name_tw,sport_name,tx_category_name,source_category_name,source_id)
select now(),now(),t.name,s.name,s.name_tw,sp.name,c.name,s.master_group_name,s.source_id 
from sbodds.team_sources s
join sbodds.sports sp on sp.id=s.sport_id
join sbodds.teams t on t.id=s.team_id
join sbodds.categories c on c.id=s.category_id
where s.source_id <> 1 and s.team_id is not null;

-- sync category_source from production into staging
insert ignore into sbodds_staging.default_category_maps (created_at,updated_at,tx_name,source_category_name,source_category_name_tw,sport_name,source_id)
select now(),now(),c.name,s.name,s.name_tw,sp.name,s.source_id
from sbodds.category_sources s
join sbodds.sports sp on sp.id=s.sport_id
join sbodds.categories c on c.id=s.category_id
where s.source_id <> 1 and s.category_id is not null;

-- sync group_source from production into staging
insert ignore into sbodds_staging.default_group_maps (created_at,updated_at,tx_name,source_group_name,source_group_name_tw,sport_name,source_id)
select now(),now(),g.name,s.name,s.name_tw,sp.name,s.source_id
from sbodds.group_sources s
join sbodds.sports sp on sp.id=s.sport_id
join sbodds.groups g on s.group_id = g.id
where s.source_id <> 1 and s.group_id is not null;