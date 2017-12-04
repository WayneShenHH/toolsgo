// Package auth 取得身份驗證相關功能
package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// key is the key of auth user in context
const key = "authuser"

// User 驗證 user 資料
type User struct {
	ID       uint
	Username string
	Tier     uint
}

// Setter defines a context that enables setting values.
type Setter interface {
	Set(string, interface{})
}

// FromContext 從 context 中取得身份驗證資訊
func FromContext(c context.Context) (User, bool) {
	u, ok := c.Value(key).(User)
	return u, ok
}

// ToContext 將身份驗證資料傳入 context
func ToContext(c Setter, user User) {
	c.Set(key, user)
}

// EncryptedPassword 加密字串
func EncryptedPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(hashedPassword))
	return string(hashedPassword)
}

// CompareHashAndPassword 比對 hash & password 是否相等
func CompareHashAndPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		fmt.Println(err) // nil means it is a match
	}
	return err == nil
}

// RandomToken generate token
func RandomToken() string {
	token, err := generateRandomString(32)
	if err != nil {
		// Serve an appropriately vague error to the
		// user, but log the details internally.
	}
	return token
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func generateRandomString(s int) (string, error) {
	b, err := generateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
