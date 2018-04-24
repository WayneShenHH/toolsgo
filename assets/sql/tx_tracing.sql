-- find today's running match
select m.id,g.name,m.start_time,h.name as home,a.name as away from matches m
join teams h on m.hteam_id=h.id
join teams a on m.ateam_id=a.id 
join groups g on g.id=m.group_id
where 1
and m.start_time >= DATE_FORMAT(now(),'%Y-%m-%d') 
and m.start_time <= now()
and m.sport_id = 1
and m.category_id = 797
#and m.group_id = 3318 
order by start_time
-- find today's running match_set
select id,match_id,start_time,set_type_id from match_sets where start_time > DATE_FORMAT(now(),'%Y-%m-%d') and is_running = 1 order by start_time
-- find tx relation data
select s.leader_id,s.match_id from match_sources s
where s.match_id = 42831 and s.source_id=1
-- query tx_adapter and checking origin data normal
select m.home_score,m.hteam_name,m.ateam_name,m.offer_otid,m.offer_ot,
m.bookmaker_name,m.cls,m.price_oh,m.price_oa,m.price_od,m.offer_inrunning,FROM_UNIXTIME(m.offer_ts/1000) as offer_ts 
from price_updates m where 1
and match_id = 4755146
and bookmaker_id = 282
and offer_inrunning = 1
#and offer_otid in (245,6,61,59)
#and (cls = -8.5 or cls = 220.5)
order by 
m.offer_ot,m.cls,
m.bookmaker_name,m.offer_ts
-- ft-odds should recieve zero odds at the end
select d.leader_id,d.id,d.offer_line_id,d.is_running,d.name,d.origin_line,d.origin_home_odds,d.origin_away_odds,b.name,FROM_UNIXTIME(d.offer_ts/1000) as offer_ts
from odds d
join book_makers b on b.id = d.book_maker_id
where 1
and d.id in (
select max(id) as id 
from odds d
where d.match_id = 433147 
and d.is_parlay = 0
and d.is_asians = 0
and d.is_running = 1
and d.offer_line_id = 1
#and d.book_maker_id = 34
#and d.name = 'point'
#and origin_line = 1.7500
group by d.name,d.origin_line,d.book_maker_id,d.offer_line_id
)
order by d.name,d.origin_line,d.offer_ts
-- check each odds on a specific match
select d.leader_id,d.id,d.offer_line_id,d.is_running,d.name,d.origin_line,d.origin_home_odds,d.origin_away_odds,b.name,FROM_UNIXTIME(d.offer_ts/1000) as offer_ts
from odds d
join book_makers b on b.id = d.book_maker_id
where 1
and d.match_id = 232272 
and d.is_parlay = 0
and d.is_asians = 0
and d.is_running = 1
and d.offer_line_id = 1
#and d.book_maker_id = 42
#and d.origin_line = 218.5
order by 
#d.name,d.origin_line,
d.book_maker_id,d.offer_ts
-- check offer
select o.id,o.updated_at,o.match_set_id,p.code,
	p.is_running,p.is_parlay,b.name,o.selected_odds_id,o.available,
	d.origin_line,origin_home_odds,d.origin_away_odds 
from match_set_offers o 
join play_types p on p.id=o.play_type_id
join odds d on d.id=o.selected_odds_id
join book_makers b on b.id=o.book_maker_id
where o.match_set_id = 50445 and p.is_running=1;
-- full running point
select id from match_set_offers where match_set_id= (select id from match_sets where match_id = 410498 and set_type_id=1) and play_type_id = 3
-- check match count 
select hteam_name,ateam_name from price_updates where bookmaker_id = 282 
and match_time >= UNIX_TIMESTAMP('2018-04-17 04:00:00') * 1000 
and match_time <= UNIX_TIMESTAMP('2018-04-18 04:00:00') * 1000 
group by hteam_name,ateam_name