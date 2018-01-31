SET @cid = (select id from categories where name = 'BANBA');
select m.id,m.start_time,h.name as home,a.name as away from matches m
join teams h on m.hteam_id=h.id
join teams a on m.ateam_id=a.id
where m.category_id = @cid and m.created_at > DATE_ADD(now(), INTERVAL -1 DAY)
order by id desc;
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