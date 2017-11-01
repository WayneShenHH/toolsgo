DELIMITER ||
DROP TRIGGER IF EXISTS update_match_set_offers;
||
DELIMITER @@
CREATE TRIGGER update_match_set_offers AFTER UPDATE ON match_set_offers
  FOR EACH ROW
  BEGIN	
    DECLARE act varchar(255);
	DECLARE modi decimal(7,4);
	IF (OLD.juice <> NEW.juice) THEN
		SET @act = 'juice_manual_change';
		SET @modi = (NEW.juice - OLD.juice);
	END IF;
	IF (OLD.home_odds_modifier <> NEW.home_odds_modifier) THEN
		SET @act = 'home_odds_manual_change';
		SET @modi = NEW.home_odds_modifier - OLD.home_odds_modifier;
	END IF;
	IF (OLD.away_odds_modifier <> NEW.away_odds_modifier) THEN
		SET @act = 'away_odds_manual_change';
		SET @modi = NEW.away_odds_modifier - OLD.away_odds_modifier;
	END IF;
	IF (OLD.home_odds_modifier <> NEW.home_odds_modifier AND OLD.away_odds_modifier <> NEW.away_odds_modifier) THEN
		SET @act = 'pair_odds_manual_change';
		SET @modi = NEW.home_odds_modifier - OLD.home_odds_modifier;
	END IF;
	IF (OLD.draw_odds_modifier <> NEW.draw_odds_modifier) THEN
		SET @act = 'draw_odds_manual_change';
		SET @modi = NEW.draw_odds_modifier - OLD.draw_odds_modifier;
	END IF;
	IF (OLD.asian_proportion_modifier <> NEW.asian_proportion_modifier) THEN
		SET @act = 'asian_manual_change';
		SET @modi = NEW.asian_proportion_modifier - OLD.asian_proportion_modifier;
	END IF;
	IF (OLD.line_modifier <> NEW.line_modifier) THEN
		SET @act = 'line_manual_change';
		SET @modi = NEW.line_modifier - OLD.line_modifier;
	END IF;
	IF (OLD.stat <> NEW.stat) THEN
		SET @act = 'stat_manual_change';
	END IF;
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
	( created_at,
	  updated_at,
	  user_id,
	  ip,
	  action,
	  stat,
	  modifier,
      data ) 
	VALUES
	( NOW(),
	  NOW(),
	  NEW.user_id,
	  NEW.ip,
	  @act,
	  NEW.stat,
	  @modi,
	  @data );
  END;
@@