package repositoryimpl

import "github.com/WayneShenHH/toolsgo/models/entities"

func (db *datastore) GetSourceCategoryByStruct(filter entities.CategorySource) *entities.CategorySource {
	c := &entities.CategorySource{}
	db.mysql.Where(filter).Find(c)
	return c
}
