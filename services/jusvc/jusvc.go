package jusvc

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/WayneShenHH/toolsgo/models"
	"github.com/WayneShenHH/toolsgo/models/entities"
	"github.com/WayneShenHH/toolsgo/repository"
	"github.com/WayneShenHH/toolsgo/tools"
	"github.com/WayneShenHH/toolsgo/tools/timeutil"
)

type JuService struct {
	repository.Repository
}

// New instence JuService
func New(ctx repository.Repository) *JuService {
	return &JuService{
		Repository: ctx,
	}
}

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
	bytes := tools.LoadJson("msg_setting")
	msgSetting := &models.Message{}
	json.Unmarshal(bytes, msgSetting)
	if msgSetting.Offer.Bid == 0 {
		msgSetting.Offer = models.SourceOffer{
			Bid:        999,
			OtName:     "point",
			Head:       -1.0,
			Proportion: 50,
			Halves:     "full",
			Hodd:       0.95,
			Aodd:       0.95,
			IsAsians:   true,
		}
	}
	o := &msgSetting.Offer
	msgSetting.Offer.PushID = fmt.Sprintf("%v_%v_%v_%v_%v", mid, o.Halves, o.OtName, o.Head, o.Bid)
	return msgSetting
}
func (service *JuService) CreateTxMatch(mid uint) {
	var sid uint = 1
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
		fmt.Println("data may not complete")
	}
	config := getMsgConfig(mid)
	message := models.Message{
		Match: models.SourceMatch{
			ID:         0,
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
			Ts:       timeutil.TimeToStamp(time.Now()),
			AdpterTs: timeutil.TimeToStamp(time.Now()),
			OfferTs:  timeutil.TimeToStamp(time.Now()),
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
		message.Match.Minute = config.Match.Minute
	}
	tools.Log(message)
	bytes, _ := json.Marshal(message)
	service.Repository.Rpush("worker:match:message", bytes)
}
func (service *JuService) Clear() {
	service.Repository.FlushDB()
	service.Repository.ClearWorkerData()
}
