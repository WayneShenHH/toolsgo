select name from teams where category_id=(select id from categories where name='BANBA')

truncate default_team_maps;
truncate default_group_maps;
truncate default_category_maps;

truncate matches;
truncate match_sources;

delete from category_sources where source_id=2;
delete from group_sources where source_id=2;
delete from team_sources where source_id=2;

SELECT @@global.time_zone, @@session.time_zone;

select * from matches where category_id=(select id from categories where name='BANBA') 



