package util

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/WayneShenHH/toolsgo/module/logger"
)

// CompareHashAndPassword 比對 hash & password 是否相等
func CompareHashAndPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		logger.Error(err) // nil means it is a match
	}
	return err == nil
}

// EncryptedPassword 加密字串
func EncryptedPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	logger.Debug(string(hashedPassword))
	return string(hashedPassword), nil
}
