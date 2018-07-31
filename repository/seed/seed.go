package seed

import "github.com/WayneShenHH/toolsgo/repository/repositoryimpl"

// Seed seed on db reset
func Seed() {
	repo := repositoryimpl.New()
	AddPlayTypes(repo)
}
