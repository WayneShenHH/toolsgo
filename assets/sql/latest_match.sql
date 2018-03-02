SET @cid = (select id from categories where name = 'BANBA');
select m.id,m.start_time,h.name as home,a.name as away from matches m
join teams h on m.hteam_id=h.id
join teams a on m.ateam_id=a.id
where m.category_id = @cid
order by start_time desc;
-- check set
select distinct match_id from match_sets where match_id  in (
	select id from matches where start_time between CONCAT(DATE_FORMAT(DATE_ADD(now(), INTERVAL -1 DAY),'%Y-%m-%d'), ' 16:00:00') and CONCAT(DATE_FORMAT(now(),'%Y-%m-%d'), ' 16:00:00') and sport_id = 1
)
select id from match_sets where match_id  in (
	select id from matches where start_time between CONCAT(DATE_FORMAT(DATE_ADD(now(), INTERVAL -1 DAY),'%Y-%m-%d'),' 16:00:00') and CONCAT(DATE_FORMAT(now(),'%Y-%m-%d'), ' 16:00:00') and sport_id = 1
)
-- check soccer odds everyday
select matches.id,match_sets.id,match_set_offers.id,odds.id,matches.available from matches 
left join match_sets on matches.id = match_sets.match_id
left join match_set_offers on match_sets.id = match_set_offers.match_set_id 
left join odds on odds.match_set_offer_id = match_set_offers.id 
where 1=1 and matches.sport_id = 1 and odds.is_book_maker_flat = 1 
and matches.start_time between CONCAT(DATE_FORMAT(DATE_ADD(now(), INTERVAL -1 DAY),'%Y-%m-%d'),' 16:00:00') and CONCAT(DATE_FORMAT(now(),'%Y-%m-%d'),' 16:00:00') 
and match_set_offers.selected_odds_id is not null
-- check odds-protal 
select d.id,d.leader_id,d.name,d.set,d.is_asians,d.origin_line,d.origin_home_odds,d.origin_away_odds,d.origin_draw_odds,b.name,FROM_UNIXTIME(d.offer_ts/1000) from odds d
join book_makers b on b.id = d.book_maker_id
where d.match_leader_id = (
select match_leader_id from odds where id = 17428015
)
and d.is_parlay = 0 
and d.set = 'full'
and d.is_book_maker_flat = 1 
and d.book_maker_id = 42 -- 42 126 34
-- check last odds always zero on each line of offers
select d.leader_id,d.id,d.is_book_maker_flat,d.is_running,d.name,d.origin_line,d.origin_home_odds,d.origin_away_odds,b.name,FROM_UNIXTIME(d.offer_ts/1000) as offer_ts
from odds d
join book_makers b on b.id = d.book_maker_id
where 1
and d.id in (
select max(id) as id 
from odds d
where d.match_set_id = 2 
and d.is_parlay = 0
and d.is_asians = 0
#and d.book_maker_id = 34
#and d.name = 'point'
#and origin_line = 1.7500
group by d.name,d.origin_line,d.book_maker_id
)
order by d.offer_ts
-- find match
select * from categories where name = 'BAIRI'
select * from matches where category_id = 1062 and start_time > '2018-02-13'
select * from match_sets where match_id = 795714
-- find odds info
select FROM_UNIXTIME(d.offer_ts/1000) as ts,d.name,d.set,d.origin_line,d.origin_home_odds,d.origin_away_odds,b.name,c.name,g.name,h.name,a.name
from odds d 
join matches m on m.id = d.match_leader_id
join book_makers b on b.id = d.book_maker_id
join teams h on h.id = m.hteam_id
join teams a on a.id = m.ateam_id
join categories c on c.id = m.category_id
join groups g on g.id = m.group_id
where d.match_set_id in (36373,36374,36408)
and d.book_maker_id = 42
and d.name = 'point'
and d.is_running = 1 and is_parlay = 0
order by ts