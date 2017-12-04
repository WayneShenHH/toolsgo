package models

import "github.com/WayneShenHH/toolsgo/models/entities"

// UserLogin 使用者登入參數
// swagger:parameters UserLogin
type UserLogin struct {
	Username string
	Password string
}

// UserLoginResult 使用者登入回傳
// swagger:model UserLoginResult
type UserLoginResult struct {
	// Token  string
	// UserID uint
	// Tier   uint
	EncryptedPassword  string `json:"-"`
	ResetPasswordToken string `json:"-"`
	*entities.User
}
