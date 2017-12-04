package repositoryimpl

import "github.com/WayneShenHH/toolsgo/models/entities"

func (db *datastore) GetSourceTeamByStruct(filter entities.TeamSource) *entities.TeamSource {
	t := &entities.TeamSource{}
	db.mysql.Where(filter).Find(t)
	return t
}
