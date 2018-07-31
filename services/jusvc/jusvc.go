package jusvc

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/WayneShenHH/toolsgo/models"
	"github.com/WayneShenHH/toolsgo/models/entities"
	"github.com/WayneShenHH/toolsgo/repository"
	"github.com/WayneShenHH/toolsgo/tools"
	"github.com/WayneShenHH/toolsgo/tools/timeutil"
)

// JuService deal with ju message
type JuService struct {
	repository.Repository
}

// New instence JuService
func New(ctx repository.Repository) *JuService {
	return &JuService{
		Repository: ctx,
	}
}

// CreateJuMatch create a ju message for testing comparer
func (service *JuService) CreateJuMatch(mid uint) {
	var sid uint = 2
	m := service.Repository.GetMatchByID(mid)
	spid := m.SportID
	h := entities.TeamSource{
		SourceID: sid,
		TeamID:   m.HteamID,
	}
	a := entities.TeamSource{
		SourceID: sid,
		TeamID:   m.AteamID,
	}
	g := entities.GroupSource{
		SourceID: sid,
		GroupID:  m.GroupID,
	}
	c := entities.CategorySource{
		SourceID:   sid,
		CategoryID: m.CategoryID,
	}
	c2 := service.Repository.GetSourceCategoryByStruct(c)
	g2 := service.Repository.GetSourceGroupByStruct(g)
	h2 := service.Repository.GetSourceTeamByStruct(h)
	a2 := service.Repository.GetSourceTeamByStruct(a)
	if c2.ID == 0 || g2.ID == 0 || h2.ID == 0 || a2.ID == 0 {
		panic("data may not complete")
	}
	config := getMsgConfig(mid)
	message := models.Message{
		Match: models.SourceMatch{
			ID:           0,
			SportID:      spid,
			StartTime:    timeutil.TimeToString(m.StartTime),
			StartDate:    timeutil.TimeToYMD(m.StartTime),
			StartTS:      timeutil.TimeToStamp(m.StartTime),
			HteamCH:      h2.Name,
			AteamCH:      a2.Name,
			GroupNameCh:  g2.Name,
			CategoryName: c2.Name,
		},
		Offer: config.Offer,
		MessageTime: models.MessageTime{
			Ts:       timeutil.TimeToStamp(time.Now()),
			AdpterTs: timeutil.TimeToStamp(time.Now()),
			OfferTs:  timeutil.TimeToStamp(time.Now()),
		},
		SourceType: "ju",
	}
	message.Offer.Bid = 999
	tools.Log(message)
	bytes, _ := json.Marshal(message)
	service.Repository.Rpush("worker:match:message", bytes)
}
func getMsgConfig(mid uint) *models.Message {
	bytes := tools.LoadJSON("msg_setting")
	msgSetting := &models.Message{}
	json.Unmarshal(bytes, msgSetting)
	if msgSetting.Offer.Bid == 0 {
		msgSetting.Offer = models.SourceOffer{
			Bid:        999,
			OtName:     "point",
			Head:       "-1.0",
			Proportion: 50,
			Halves:     "full",
			Hodd:       0.95,
			Aodd:       0.95,
			IsAsians:   true,
		}
	}
	o := &msgSetting.Offer
	msgSetting.Offer.PushID = fmt.Sprintf("%v_%v_%v_%v_%v", mid, o.HalvesType, o.PlayType, o.Head, o.Bid)
	return msgSetting
}

// CreateTxMatch create a tx message for testing comparer
func (service *JuService) CreateTxMatch(mid uint) {
	var sid uint = 1
	m := service.Repository.GetMatchByID(mid)
	spid := m.SportID
	m2 := service.Repository.GetSourceMatchByStruct(entities.MatchSource{MatchID: m.ID, SourceID: sid})
	c2 := service.Repository.GetSourceCategoryByStruct(entities.CategorySource{SourceID: sid, CategoryID: m.CategoryID})
	g2 := service.Repository.GetSourceGroupByStruct(entities.GroupSource{SourceID: sid, GroupID: m.GroupID})
	h2 := service.Repository.GetSourceTeamByStruct(entities.TeamSource{SourceID: sid, TeamID: m.HteamID})
	a2 := service.Repository.GetSourceTeamByStruct(entities.TeamSource{SourceID: sid, TeamID: m.AteamID})
	if c2.ID == 0 || g2.ID == 0 || h2.ID == 0 || a2.ID == 0 {
		fmt.Println("data may not complete")
	}
	config := getMsgConfig(mid)
	if m2.LeaderID == 0 {
		m2.LeaderID = config.Match.ID
	}
	message := models.Message{
		Match: models.SourceMatch{
			ID:         m2.LeaderID,
			SportID:    spid,
			StartTime:  timeutil.TimeToString(m.StartTime),
			StartDate:  timeutil.TimeToYMD(m.StartTime),
			StartTS:    timeutil.TimeToStamp(m.StartTime),
			HteamID:    h2.LeaderID,
			AteamID:    a2.LeaderID,
			GroupID:    g2.LeaderID,
			CategoryID: c2.LeaderID,
		},
		Offer: config.Offer,
		MessageTime: models.MessageTime{
			OfferTs:  config.OfferTs,
			AdpterTs: timeutil.TimeToStamp(time.Now()),
			Ts:       timeutil.TimeToStamp(time.Now()),
		},
		SourceType: "tx",
	}
	if message.Offer.Hodd != 0 {
		message.Offer.Hodd++
	}
	if message.Offer.Aodd != 0 {
		message.Offer.Aodd++
	}
	if message.Offer.Dodd != 0 {
		message.Offer.Dodd++
	}
	if message.Offer.IsRunning {
		message.Match.MatchState = config.Match.MatchState
		message.Match.StateString = config.Match.StateString
		message.Match.HomeScore = config.Match.HomeScore
		message.Match.AwayScore = config.Match.AwayScore
		message.Match.HomeRedcard = config.Match.HomeRedcard
		message.Match.AwayRedcard = config.Match.AwayRedcard
		message.Match.Gametime = config.Match.Gametime
		message.Match.GameMinute = config.Match.GameMinute
	}
	if config.OfferTs == 0 {
		message.OfferTs = timeutil.TimeToStamp(time.Now())
	}
	service.MultiInsert(config.Mul, message)
}

//MultiInsert multiple insert messages
func (service *JuService) MultiInsert(cnt int, message models.Message) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < cnt; i++ {
		tools.Log(message)
		bytes, _ := json.Marshal(message)
		service.Repository.Rpush("worker:match:message", bytes)
		message.OfferTs++
		message.Offer.Hodd = r.Float64()/10 + 1.9
		message.Offer.Aodd = r.Float64()/10 + 1.9
	}
}

// Clear init all data related offer
func (service *JuService) Clear() {
	service.Repository.FlushDB()
	service.Repository.ClearWorkerData()
}
