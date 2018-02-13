-- team
select s.* from team_sources s left join teams t on s.team_id = t.id
where s.team_id > 0 and t.id is null
-- group
select s.* from group_sources s left join groups t on s.group_id = t.id
where s.group_id > 0 and t.id is null
-- category
select s.* from category_sources s left join categories t on s.category_id = t.id
where s.category_id > 0 and t.id is null
-- matches
select s.* from match_sources s left join matches t on s.match_id = t.id
where s.match_id > 0 and t.id is null
----------------------------------------------
GO_ENV=production go/bin/libgo tx:sync team