
select name from teams where category_id=(select id from categories where name='BANBA')

truncate default_team_maps;
truncate default_group_maps;
truncate default_category_maps;

truncate matches;
truncate match_sources;

truncate match_sets;
truncate match_set_offers;
truncate odds;

truncate team_sources;
truncate category_sources;
truncate group_sources;
truncate categories;
truncate groups;
truncate teams;

delete from category_sources where source_id=2;
delete from group_sources where source_id=2;
delete from team_sources where source_id=2;

SELECT @@global.time_zone, @@session.time_zone;

select * from matches where category_id=(select id from categories where name='BANBA') 

update users set access_token='6s6zXKlB7IGaqt5MLJzGs7xss81FjeYK45jUynRWnVk=' where username='tier1'