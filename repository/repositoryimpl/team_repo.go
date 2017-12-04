package repositoryimpl

import "gitlab.cow.bet/bkd_tool/libgo/models/entities"

func (db *datastore) GetSourceTeamByStruct(filter entities.TeamSource) *entities.TeamSource {
	t := &entities.TeamSource{}
	db.mysql.Where(filter).Find(t)
	return t
}
