package repositoryimpl

import "gitlab.cow.bet/bkd_tool/libgo/models/entities"

func (db *datastore) GetSourceCategoryByStruct(filter entities.CategorySource) *entities.CategorySource {
	c := &entities.CategorySource{}
	db.mysql.Where(filter).Find(c)
	return c
}
