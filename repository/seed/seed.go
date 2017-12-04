package seed

import "github.com/WayneShenHH/toolsgo/repository/repositoryimpl"

func Seed() {
	repo := repositoryimpl.New()
	AddPlayTypes(repo)
}
