UPDATE sport_offer_types SET leader_offer_type_id='2' WHERE id='7';
UPDATE sport_offer_types SET leader_offer_type_id='32' WHERE id='21';

UPDATE sport_offer_types SET leader_offer_type_id='648' WHERE id='49';
UPDATE sport_offer_types SET leader_offer_type_id='64' WHERE id='55';
UPDATE sport_offer_types SET leader_offer_type_id='63' WHERE id='43';
--
ALTER TABLE match_set_offers 
DROP INDEX idx_offer ,
ADD UNIQUE INDEX idx_offer (match_id ASC, offer_type_id ASC, line_id ASC, is_running ASC, is_parlay ASC, is_asians ASC);

