-- game_index
select m.sport_id,o.match_id,m.start_time,m.group_id,o.match_set_id as set_id,o.match_set_id as offer_id,o.is_running,o.is_parlay,g.is_special 
from offer_stats o join matches m on m.id=o.match_id join groups g on g.id=m.group_id 
where o.available = 1
-- total
select count(distinct m.id) as total_match_amount 
from matches m left join offer_stats o on o.match_id = m.id
where m.start_time > '%v' and (o.available = 1 or o.id is null)
-- match_with_offer_amount
select count(distinct o.match_id) as match_with_offer_amount from offer_stats o 
join odds d on d.match_set_offer_id = o.match_set_offer_id 
where d.is_book_maker_flat = 1 and o.start_time > '%v' 
-- on_sell_match_amount
select count(distinct o.match_id) from offer_stats o 
where o.available = 1
-- count running match 
select count(distinct o.match_id) from offer_stats o 
where o.is_running = 1 and start_time > '%v'
-- running_on_sell_match_amount
select count(distinct o.match_id) from offer_stats o 
where o.available = 1 and o.is_running = 1
-- today_on_sell_match_amount
select count(distinct o.match_id) from offer_stats o 
where o.available = 1 and o.start_time between '%v' and '%v'
-- GetCompetitions
select m.* from offer_stats o 
left join matches m on o.match_id=m.id 
left join match_sets s on s.id = o.match_set_id
left join match_set_offers mso on mso.id = o.match_set_offer_id
left join odds d on d.match_set_offer_id = mso.id
where 1
and o.start_time between '%v' and '%v'
and d.is_book_maker_flat = 1 
and mso.selected_odds_id is not null
and o.available = 1
and o.is_running = 1
and ((o.available = 0 and o.available_time is null) or (o.available = 0 and o.is_running = 1) or (IF(o.is_running, (s.is_payout), (o.start_time < NOW()))))
-- GetMatchSetForClose
select o.id from offer_stats o
left join match_results r on r.match_id=o.match_id and r.set_type_id=o.set_type_id
left join match_sets s on s.id = o.match_set_id and s.is_running = o.is_running
where o.available = 1 and (r.id is not null or o.start_time < '%v' or s.id is null or (o.updated_at < '%v' and o.is_running = 1))
-- GetSportLeagues
select m.category_id as id,count(distinct m.id) as match_count,o.is_running from offer_stats o 
join matches m on m.id = o.match_id and o.available = 1 
group by m.category_id,o.is_running 