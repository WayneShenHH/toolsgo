DELIMITER ||
DROP TRIGGER IF EXISTS update_match_set_offers;
||
DELIMITER @@
CREATE TRIGGER update_match_set_offers AFTER UPDATE ON match_set_offers
FOR EACH ROW
BEGIN	
	DECLARE act varchar(255);
	DECLARE modi varchar(255);
	SET @act = '';
	IF (OLD.home_threshold <> NEW.home_threshold) THEN
		SET @act = 'home_threshold_manual_change';
		SET @modi = JSON_OBJECT(
			'New',JSON_OBJECT('HomeThreshold',NEW.home_threshold),
			'Old',JSON_OBJECT('HomeThreshold',OLD.home_threshold)
		);
	END IF;
	IF (OLD.away_threshold <> NEW.away_threshold) THEN
		SET @act = 'away_threshold_manual_change';
		SET @modi = JSON_OBJECT(
			'New',JSON_OBJECT('AwayThreshold',NEW.away_threshold),
			'Old',JSON_OBJECT('AwayThreshold',OLD.away_threshold)
		);
	END IF;
	IF (OLD.draw_threshold <> NEW.draw_threshold) THEN
		SET @act = 'draw_threshold_manual_change';
		SET @modi = JSON_OBJECT(
			'New',JSON_OBJECT('DrawThreshold',NEW.draw_threshold),
			'Old',JSON_OBJECT('DrawThreshold',OLD.draw_threshold)
		);
	END IF;
	IF (OLD.offer_amount_limit <> NEW.offer_amount_limit) THEN
		SET @act = 'limit_manual_change';
		SET @modi = JSON_OBJECT(
			'New',JSON_OBJECT('OfferAmountLimit',NEW.offer_amount_limit),
			'Old',JSON_OBJECT('OfferAmountLimit',OLD.offer_amount_limit)
		);
	END IF;
	IF (OLD.juice <> NEW.juice) THEN
		SET @act = 'juice_manual_change';
		SET @modi = JSON_OBJECT(
			'New',JSON_OBJECT('Juice',NEW.juice),
			'Old',JSON_OBJECT('Juice',OLD.juice)
		);
	END IF;
	IF (OLD.home_odds_modifier <> NEW.home_odds_modifier) THEN
		SET @act = 'home_odds_manual_change';
		SET @modi = JSON_OBJECT(
			'New',JSON_OBJECT('HomeOddsModifier',NEW.home_odds_modifier),
			'Old',JSON_OBJECT('HomeOddsModifier',OLD.home_odds_modifier)
		);
	END IF;
	IF (OLD.away_odds_modifier <> NEW.away_odds_modifier) THEN
		SET @act = 'away_odds_manual_change';
		SET @modi = JSON_OBJECT(
			'New',JSON_OBJECT('AwayOddsModifier',NEW.away_odds_modifier),
			'Old',JSON_OBJECT('AwayOddsModifier',OLD.away_odds_modifier)
		);
	END IF;
	IF (OLD.home_odds_modifier <> NEW.home_odds_modifier AND OLD.away_odds_modifier <> NEW.away_odds_modifier) THEN
		SET @act = 'pair_odds_manual_change';
		SET @modi = JSON_OBJECT(
			'New',JSON_OBJECT('HomeOddsModifier',NEW.home_odds_modifier,'AwayOddsModifier',NEW.away_odds_modifier),
			'Old',JSON_OBJECT('HomeOddsModifier',OLD.home_odds_modifier,'AwayOddsModifier',OLD.away_odds_modifier)
		);
	END IF;
	IF (OLD.draw_odds_modifier <> NEW.draw_odds_modifier) THEN
		SET @act = 'draw_odds_manual_change';
		SET @modi = JSON_OBJECT(
			'New',JSON_OBJECT('DrawOddsModifier',NEW.draw_odds_modifier),
			'Old',JSON_OBJECT('DrawOddsModifier',OLD.draw_odds_modifier)
		);
	END IF;
	IF (OLD.asian_proportion_modifier <> NEW.asian_proportion_modifier) THEN
		SET @act = 'asian_manual_change';
		SET @modi = JSON_OBJECT(
			'New',JSON_OBJECT('AsianProportionModifier',NEW.asian_proportion_modifier),
			'Old',JSON_OBJECT('AsianProportionModifier',OLD.asian_proportion_modifier)
		);
	END IF;
	IF (OLD.line_modifier <> NEW.line_modifier) THEN
		SET @act = 'line_manual_change';
		SET @modi = JSON_OBJECT(
			'New',JSON_OBJECT('LineModifier',NEW.line_modifier),
			'Old',JSON_OBJECT('LineModifier',OLD.line_modifier)
		);
	END IF;
	IF (OLD.stat <> NEW.stat) THEN
		SET @act = 'stat_manual_change';
		SET @modi = JSON_OBJECT(
			'New',JSON_OBJECT('Stat',NEW.stat),
			'Old',JSON_OBJECT('Stat',OLD.stat)
		);
	END IF;
	IF (@act <> '') THEN
		SET @data = JSON_OBJECT(
			'ID',NEW.ID,
			'MatchSetID',NEW.match_set_id,
			'PlayTypeID',NEW.play_type_id,
			'HomeStake',NEW.home_stake,
			'AwayStake',NEW.away_stake,
			'DrawStake', NEW.draw_stake,
			'HomeThreshold',NEW.home_threshold,
			'AwayThreshold',NEW.away_threshold,
			'DrawThreshold',NEW.draw_threshold,
			'UserID',NEW.user_id,
			'IP',NEW.ip,
			'AsianProportionModifier',NEW.asian_proportion_modifier,
			'LineModifier',NEW.line_modifier,
			'Stat',NEW.stat,
			'Available',NEW.available,
			'HomeOddsModifier',NEW.home_odds_modifier,
			'AwayOddsModifier',NEW.away_odds_modifier,
			'DrawOddsModifier',NEW.draw_odds_modifier,
			'Juice',NEW.juice,
			'OfferAmountLimit',NEW.offer_amount_limit,
			'OnceAmountLimit',NEW.once_amount_limit,
			'OrderDelay',NEW.order_delay,
			'BookMakerID',NEW.book_maker_id,
			'SelectedOddsID',NEW.selected_odds_id
		);
		INSERT INTO log_operators 
		( 	created_at,
			updated_at,
			user_id,
			ip,
			entity,
			action,
			stat,
			changes,
			data ) 
		VALUES
		( 	NOW(),
			NOW(),
			NEW.user_id,
			NEW.ip,
			'match_set_offer',
			@act,
			NEW.stat,
			@modi,
			@data );
	END IF;
END;
@@