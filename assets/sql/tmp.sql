set @todaycnt = (select count(*) from matches where DATE_FORMAT(matches.start_time, '%Y-%m-%d') = CURDATE());

set @oddscnt = (select count(distinct match_sets.match_id) from match_sets 
join match_set_offers on match_sets.id = match_set_offers.match_set_id 
and match_set_offers.available = 1 and match_sets.available = 1 
and match_set_offers.selected_odds_id is not null 
and match_set_offers.play_type_id in (select id from play_types where is_parlay = 0));

set @offercnt = (select count(distinct match_sets.match_id) from match_sets 
join match_set_offers on match_sets.id = match_set_offers.match_set_id 
and match_set_offers.available = 1 and match_sets.available = 1 );

set @runningcnt = (select count(distinct match_sets.match_id) from match_sets 
join match_set_offers on match_sets.id = match_set_offers.match_set_id 
and match_set_offers.available = 1 and match_sets.available = 1 and match_sets.is_running=1);

select @offercnt - @oddscnt as offer_not_sell,@oddscnt as offer_sell,@todaycnt as today,@runningcnt as ruunning;