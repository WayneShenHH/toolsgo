package usersvc

import (
	"github.com/WayneShenHH/toolsgo/repository"
	"github.com/WayneShenHH/toolsgo/tools/auth"
)

// UserService for user data
type UserService struct {
	repository.Repository
	AuthUser auth.User
}

// New init usersvc
func New(base repository.Repository, user auth.User) *UserService {
	return &UserService{
		Repository: base,
		AuthUser:   user,
	}
}
