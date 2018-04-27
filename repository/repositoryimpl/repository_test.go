package repositoryimpl

import (
	"testing"
	"time"

	"github.com/WayneShenHH/toolsgo/repository"
)

func repo() repository.Repository {
	dbInstance := repository.DBConnect()
	r := repository.RedisConnect()
	dbInstance.LogMode(true)
	return &datastore{
		mysql: dbInstance,
		cache: r,
	}
}
func Test_CheckTxSchdule(t *testing.T) {
	s := time.Now()
	e := s.Add(time.Hour * 2)
	repo().GetMatchesByTime(s, e)
}
