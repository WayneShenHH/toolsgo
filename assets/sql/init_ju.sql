
select name from teams where category_id=(select id from categories where name='BANBA')

truncate default_team_maps;
truncate default_group_maps;
truncate default_category_maps;

truncate matches;
truncate match_sources;

truncate match_sets;
truncate match_set_offers;
truncate odds;

truncate team_sources;
truncate category_sources;
truncate group_sources;
truncate categories;
truncate groups;
truncate teams;

delete from category_sources where source_id <> 1;
delete from group_sources where source_id <> 1;
delete from team_sources where source_id <> 1;
select * from category_sources where source_id <> 1;
select * from group_sources where source_id <> 1;
select * from team_sources where source_id <> 1;

SELECT @@global.time_zone, @@session.time_zone;

select * from matches where category_id=(select id from categories where name='BANBA') 

update users set access_token='6s6zXKlB7IGaqt5MLJzGs7xss81FjeYK45jUynRWnVk=' where username='admin';

show engine innodb status;
show status where `variable_name` = 'Threads_connected';
show processlist;

select * from match_set_offers where created_at > DATE_ADD(now(), INTERVAL -1 DAY);
select * from match_sets where created_at > DATE_ADD(now(), INTERVAL -1 DAY);
select * from matches where created_at > DATE_ADD(now(), INTERVAL -1 DAY);
select * from odds where created_at > DATE_ADD(now(), INTERVAL -1 DAY);

-- open bookmaker for tx
select * from book_makers where ref_id
in (83,126,282,285,327,365,539);
update book_makers set available = 1 where name
in ('PinnacleSports','Bet 365','Singbet','IBCBET','sbobet.com','188bet','Marathonbet');

-- set auto_increment
alter table ft_group_sources auto_increment = 57399