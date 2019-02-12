package repository

import (
	"time"

	"github.com/WayneShenHH/toolsgo/models"
	"github.com/WayneShenHH/toolsgo/models/entities"
)

// Repository interface for dao
type Repository interface {
	GetMatchesByTime(start, end time.Time) []entities.Match
	GetMatchByID(id uint) *entities.Match
	GetSourceCategoryByStruct(filter entities.CategorySource) *entities.CategorySource
	GetSourceGroupByStruct(filter entities.GroupSource) *entities.GroupSource
	GetSourceTeamByStruct(filter entities.TeamSource) *entities.TeamSource
	GetSourceMatchByStruct(filter entities.MatchSource) *entities.MatchSource
	GetUser(username string) (*entities.User, error)
	UpdatesUser(id uint, fields *entities.User) error
	GetUserByToken(token string) (*entities.User, error)
	AddPlayTypeByStruct(playType *entities.PlayType) error
	Hset(key string, field string, value []byte)
	Rpush(key string, value []byte)
	Blpop(key string) []byte
	LRange(key string, start int, end int) []interface{}
	FlushDB()
	ClearWorkerData()
	TxMessage(mid uint) []models.TxMessage
	GetOldData() *[]models.OfferHierarchy
	ClearOldData()
	ClockIn()
}
