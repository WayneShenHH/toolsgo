package repository

import "gitlab.cow.bet/bkd_tool/libgo/models/entities"

type Repository interface {
	GetMatchByID(id uint) *entities.Match
	GetSourceCategoryByStruct(filter entities.CategorySource) *entities.CategorySource
	GetSourceGroupByStruct(filter entities.GroupSource) *entities.GroupSource
	GetSourceTeamByStruct(filter entities.TeamSource) *entities.TeamSource
	Hset(key string, field string, value []byte)
	Rpush(key string, value []byte)
}
