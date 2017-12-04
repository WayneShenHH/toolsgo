package repositoryimpl

import "gitlab.cow.bet/bkd_tool/libgo/models/entities"

func (db *datastore) GetSourceGroupByStruct(filter entities.GroupSource) *entities.GroupSource {
	g := &entities.GroupSource{}
	db.mysql.Where(filter).Find(g)
	return g
}
