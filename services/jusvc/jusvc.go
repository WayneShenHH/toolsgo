package jusvc

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/WayneShenHH/toolsgo/models"
	"github.com/WayneShenHH/toolsgo/models/entities"
	"github.com/WayneShenHH/toolsgo/repository"
	"github.com/WayneShenHH/toolsgo/tools"
	"github.com/WayneShenHH/toolsgo/tools/timezone"
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
	message := models.Message{
		Match: models.SourceMatch{
			ID:           0,
			SportID:      spid,
			StartTime:    m.StartTime.Format("2006-01-02 15:04:05 +00:00"),
			StartDate:    m.StartTime.Format("2006-01-02"),
			StartTS:      timezone.TimeToStamp(m.StartTime),
			HteamCH:      h2.Name,
			AteamCH:      a2.Name,
			GroupNameCh:  g2.Name,
			CategoryName: c2.Name,
		},
		Offer: models.SourceOffer{
			PushID: fmt.Sprint(mid, "_full_point_-1.0_999"),
			Bid:    999,
			OtName: "point",
			Head:   -1.0,
			Halves: "full",
			Hodd:   0.77,
			Aodd:   1.02,
		},
		MessageTime: models.MessageTime{
			Ts:       timezone.TimeToStamp(time.Now()),
			AdpterTs: timezone.TimeToStamp(time.Now()),
			OfferTs:  timezone.TimeToStamp(time.Now()),
		},
		SourceType: "ju",
	}
	tools.Log(message)
	bytes, _ := json.Marshal(message)
	service.Repository.Rpush("worker:match:message", bytes)
}
