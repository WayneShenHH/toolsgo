set @ids = (select match_set_id from log_closes where id = 2);
set @q = (SELECT CONCAT('select * from match_sets where id in (' , @ids , ')'));
PREPARE stmt FROM @q;
EXECUTE stmt;