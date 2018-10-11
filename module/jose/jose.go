package jose

import (
	"encoding/json"
	"gitlab.cow.bet/bkd_tool/libgo/errors"

	"github.com/dgrijalva/jwt-go"
)

var secret = []byte("MAKE_FTIT_GREAT_AGIN")

// GenerateJWT 產生 JWT
func GenerateJWT(data interface{}) (string, error) {
	if data == nil {
		return "", errors.Errorf("Input data was nil")
	}

	claims, err := initClaims(data)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret)
}

// ParseJWT 解析 JWT
func ParseJWT(jwtString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(jwtString, CheckToken)
	if err != nil {
		return nil, errors.Errorf("Parse JWT failed: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.Errorf("Cast to jwt.MapClaims failed: orig: %v", token.Claims)
	}

	return claims, nil
}

// Bind 將 Claims 的 JSON Unmarshal 到 Target
func Bind(jwtString string, target interface{}) error {
	claims, err := ParseJWT(jwtString)
	if err != nil {
		return err
	}

	data, exist := claims[Data]
	if !exist {
		return errors.Errorf("Can't find data value: claims: %+v", claims)
	}

	return json.Unmarshal([]byte(data.(string)), target)
}

// CheckToken 檢查 Token 正確性
func CheckToken(token *jwt.Token) (interface{}, error) {
	// 只接受 HS256, HS384, HS512
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.Errorf("Not allow this algorithm: %v", token.Header["alg"])
	}
	return []byte(secret), nil
}
