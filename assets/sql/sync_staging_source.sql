-- truncate source maps on staging
truncate sbodds_staging.default_team_maps;
truncate sbodds_staging.default_group_maps;
truncate sbodds_staging.default_category_maps;

-- remove old source data on staging
delete from sbodds_staging.team_sources where source_id <> 1;
delete from sbodds_staging.group_sources where source_id <> 1;
delete from sbodds_staging.category_sources where source_id <> 1;
-- assign data
insert into default_team_maps select * from sbodds_staging.default_team_maps;
insert into default_group_maps select * from sbodds_staging.default_group_maps;
insert into default_category_maps select * from sbodds_staging.default_category_maps;
-- sync team_source from production into staging
insert ignore into sbodds_staging.default_team_maps (created_at,updated_at,tx_name,source_team_name,source_team_name_tw,sport_name,tx_category_name,source_category_name,source_id)
select now(),now(),t.name,s.name,s.name_tw,sp.name,c.name,s.master_group_name,s.source_id 
from sbodds.team_sources s
join sbodds.sports sp on sp.id=s.sport_id
join sbodds.teams t on t.id=s.team_id
join sbodds.categories c on c.id=s.category_id
where s.source_id <> 1 and s.team_id is not null;
-- insert teams by team_sources
insert ignore into teams (created_at, updated_at, name, sport_id, category_id)
select now(),now(),s.name,s.sport_id,(select id from categories where name = s.master_group_name) as cid
from team_sources s
where source_id = 1
-- check team maps
select s.id,s.name,tx.id,tx.name from team_sources s 
join default_team_maps d on s.name = d.source_team_name and s.master_group_name = d.tx_category_name
join teams tx on tx.name = d.tx_name and d.tx_category_name = (select name from categories where id = tx.category_id)
where s.source_id > 1 and s.team_id is null
-- update team_sources by maps
update team_sources s set s.team_id = (
    select tx.id from default_team_maps d
    join teams tx on d.tx_name = tx.name and d.tx_category_name = (select name from categories where id = tx.category_id)
    where d.source_team_name = s.name and d.tx_category_name = s.master_group_name
)
where s.source_id > 1 and s.team_id is null;
-- insert team_sources by maps
insert ignore into team_sources (created_at, updated_at, name, master_group_name, leader_id, leader_country_id, leader_sport_id, source_id, sport_id, team_id)
select now(),now(),d.source_team_name,d.source_category_name,0,0,sp.id,d.source_id,sp.id,tx.id
from default_team_maps d
join sports sp on sp.name = d.sport_name
join teams tx on d.tx_name = tx.name and d.tx_category_name = (select name from categories where id = tx.category_id)

-- sync category_source from production into staging
insert ignore into sbodds_staging.default_category_maps (created_at,updated_at,tx_name,source_category_name,source_category_name_tw,sport_name,source_id)
select now(),now(),c.name,s.name,s.name_tw,sp.name,s.source_id
from sbodds.category_sources s
join sbodds.sports sp on sp.id=s.sport_id
join sbodds.categories c on c.id=s.category_id
where s.source_id <> 1 and s.category_id is not null;
-- check category maps
select s.id,s.name,tx.id,tx.name from category_sources s 
join default_category_maps d on s.name = d.source_category_name
join categories tx on tx.name = d.tx_name
where s.source_id > 1 and s.category_id is null
-- update category_sources by maps
update category_sources s set s.category_id = (
    select id from categories tx where tx.name = (
        select d.tx_name from default_category_maps d where d.source_category_name = s.name
    )
)
where s.source_id > 1 and s.category_id is null;

-- sync group_source from production into staging
insert ignore into sbodds_staging.default_group_maps (created_at,updated_at,tx_name,source_group_name,source_group_name_tw,sport_name,source_id)
select now(),now(),g.name,s.name,s.name_tw,sp.name,s.source_id
from sbodds.group_sources s
join sbodds.sports sp on sp.id=s.sport_id
join sbodds.groups g on s.group_id = g.id
where s.source_id <> 1 and s.group_id is not null;
-- check group maps
select s.id,s.name,tx.id,tx.name from group_sources s 
join default_group_maps d on s.name = d.source_group_name
join groups tx on tx.name = d.tx_name
where s.source_id > 1 and s.group_id is null
-- update group_sources by maps
update group_sources s set s.group_id = (
    select id from groups tx where tx.name = (
        select d.tx_name from default_group_maps d where d.source_group_name = s.name
    )
)
where s.source_id > 1 and s.group_id is null;