SET @cid = (select id from categories where name = 'BANBA');
select m.id,m.start_time,h.name as home,a.name as away from matches m
join teams h on m.hteam_id=h.id
join teams a on m.ateam_id=a.id
where m.category_id = @cid and m.start_time > NOW()
order by id desc;