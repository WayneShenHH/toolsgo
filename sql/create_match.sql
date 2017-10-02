SET @mid = (select id from matches where category_id = (select id from categories where name = 'BBKOR') and start_time > '2017-09-30' limit 1);
select id, addtime(start_time,'-08:00:00') as name from matches where id = @mid
union
(select category_id as id,name from category_sources where source_id = 2 and category_id = (select id from categories where name = 'BBKOR'))
union
(select group_id as id ,name from group_sources where source_id = 2 and group_id = (select group_id from matches where id = @mid))
union
(select id,name from teams where id = (select hteam_id from matches where id = @mid))
union
(select team_id as id,name_tw as name from team_sources where source_id = 2 and team_id = (select hteam_id from matches where id = @mid))
union
(select id,name from teams where id = (select ateam_id from matches where id = @mid))
union
(select team_id as id,name_tw as name from team_sources where source_id = 2 and team_id = (select ateam_id from matches where id = @mid))