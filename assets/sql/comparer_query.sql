-- find out all group_sources which has the same leader_id 
select * from group_sources where leader_id in (
    select leader_id from (
        select count(leader_id) as cnt,leader_id from group_sources where source_id = 1 group by leader_id
    )q where cnt > 1 
)
-- restore abnormal groups
update group_sources s set s.group_id = (select g.id from groups g where g.name = s.name limit 1)
where s.group_id = 125855
-- query for abnormal groups
select s.id as group_source_id,s.group_id,g.id as gid,s.name,s.leader_id from group_sources s 
join groups g on s.name = g.name
where s.group_id <> g.id
-- remove dup groups
delete from groups where id in (
    select group_id from group_sources where leader_id in (
        select leader_id from (
            select count(leader_id) as cnt,leader_id from group_sources where source_id = 1 group by leader_id
        )q where cnt > 1 
    ) and sport_id is null
)
-- remove dup group-soueces
delete from group_sources where id in (
54958,56841,56843
)
select max(id) from group_sources where leader_id in (
    select leader_id from (
        select count(leader_id) as cnt,leader_id from group_sources where source_id = 1 group by leader_id
    )q where cnt > 1 
) group by leader_id
-- query abnormal match which has abnormal group_id
select m.id as mid,g.id as gid,m.group_id from matches m
left join groups g on m.group_id = g.id
where g.id is null
order by m.id desc 
-- check team
select s.* from team_sources s left join teams t on s.team_id = t.id
where s.team_id > 0 and t.id is null
----------------------------------------------
GO_ENV=production go/bin/libgo tx:sync team