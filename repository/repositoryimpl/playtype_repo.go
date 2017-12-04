package repositoryimpl

import "github.com/WayneShenHH/toolsgo/models/entities"

func (db *datastore) AddPlayTypeByStruct(playType *entities.PlayType) error {
	err := db.mysql.Create(playType).Error
	return err
}
