// Package topic nsq topic
package topic

// SourceOdds 訂閱賠率 topic 的 channel 名稱
const SourceOdds = "source_odds"

// SourceMatch 訂閱賽事 topic 的 channel 名稱
const SourceMatch = "source_match"

// SourceEvent 訂閱開賽事件 topic 的名稱
const SourceEvent = "source_event"

// BetradarEvent Betradar LD.bookmaker events
const BetradarEvent = "br_scout"

// BetradarOddsChangeEvent Betradar odds_change events
const BetradarOddsChangeEvent = "br_uof_feed_prematch_OddsChange"

// BetradarMatchList Betradar UOF match_list
const BetradarMatchList = "br_uof_feed_match_list"
