SET @cid = (select id from categories where name = 'BANBA');
SET @mid = (select id from matches where category_id = @cid and start_time > NOW() limit 1);
select 'match' as title, id, start_time as name from matches where id = @mid
union
(select 'category' as title,category_id as id,name from category_sources where source_id = 2 and category_id = @cid)
union
(select 'group' as title, group_id as id ,name from group_sources where source_id = 2 and group_id = (select group_id from matches where id = @mid))
union
(select 'tx_hteam' as title, id,name from teams where id = (select hteam_id from matches where id = @mid))
union
(select 'ju_hteam' as title, team_id as id,name from team_sources where source_id = 2 and team_id = (select hteam_id from matches where id = @mid))
union
(select 'tx_ateam' as title,id,name from teams where id = (select ateam_id from matches where id = @mid))
union
(select 'ju_ateam' as title, team_id as id,name from team_sources where source_id = 2 and team_id = (select ateam_id from matches where id = @mid))