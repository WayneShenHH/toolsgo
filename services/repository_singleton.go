package services

import (
	"github.com/WayneShenHH/toolsgo/repository"
	"github.com/WayneShenHH/toolsgo/repository/repositoryimpl"
)

var repo repository.Repository

func Repository() repository.Repository {
	if repo == nil {
		repo = repositoryimpl.New()
	}
	return repo
}
