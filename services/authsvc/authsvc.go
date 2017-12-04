package authsvc

import (
	"errors"
	"fmt"
	"time"

	"github.com/WayneShenHH/toolsgo/models"
	"github.com/WayneShenHH/toolsgo/models/entities"
	"github.com/WayneShenHH/toolsgo/repository"
	"github.com/WayneShenHH/toolsgo/tools/auth"
)

// AuthService 使用者權限相關操作
type AuthService struct {
	repository.Repository
}

// New instence AuthService
func New(ctx repository.Repository) *AuthService {
	return &AuthService{
		Repository: ctx,
	}
}
func (service *AuthService) Login(model models.UserLogin) (models.UserLoginResult, error) {
	var err error
	var result models.UserLoginResult
	user, err := service.Repository.GetUser(model.Username)
	if err != nil {
		// 帳號不正確
		err = errors.New("auth failed")
	}
	pwdEqual := auth.CompareHashAndPassword(user.EncryptedPassword, model.Password)
	fmt.Printf(`User Login Username:%v, EncryptedPassword:%v `, model.Username, auth.EncryptedPassword(model.Password))

	if !pwdEqual {
		err = errors.New("auth failed")
	} else {
		token := auth.RandomToken()
		// 驗證完成，更新 token 以及相關登入資訊
		service.Repository.UpdatesUser(user.ID, &entities.User{
			AccessToken:  token,
			LastSignInAt: time.Now(),
		})
		user.AccessToken = token
		result = models.UserLoginResult{
			User: user,
			// UserID: user.ID,
			// Token:  user.AccessToken,
			// Tier:   user.Tier,
		}
	}
	return result, err
}
