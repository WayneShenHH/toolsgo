package jose

import (
	"encoding/json"
	"gitlab.cow.bet/bkd_tool/libgo/errors"
	"reflect"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// TAG Jose 的標籤名稱
const TAG = "jose"

// Data Claims 的預設欄位名稱
const Data = "default_data"

// 初始化 Claims 並且產生 UUID 確保編碼會不同
func initClaims(data interface{}) (jwt.MapClaims, error) {
	if data == nil {
		return nil, errors.Errorf("jose.initClaims input data was nil")
	}

	var claims jwt.MapClaims
	var err error
	typ := deTypePtr(reflect.TypeOf(data))
	switch typ.Kind() {
	case reflect.Struct:
		claims, err = structToClaims(data)
	case reflect.Map:
		val := deValuePtr(reflect.ValueOf(data))
		claims = mapToClaims(val, typ)
	default:
		claims = jwt.MapClaims{Data: data}
	}

	if err != nil {
		return nil, err
	}

	claims["UUID"] = uuid.New().String()
	return claims, nil
}

func deTypePtr(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		return deTypePtr(t.Elem())
	}
	return t
}

func deValuePtr(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Ptr {
		return deValuePtr(v.Elem())
	}
	return v
}

func structToClaims(data interface{}) (jwt.MapClaims, error) {
	temp := make(jwt.MapClaims, 0)

	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	temp[Data] = string(bytes)

	return temp, nil
}

func mapToClaims(v reflect.Value, t reflect.Type) jwt.MapClaims {
	temp := make(jwt.MapClaims, 0)
	keys := v.MapKeys()

	for i := 0; i < len(keys); i++ {
		if key, ok := keys[i].Interface().(string); ok {
			temp[key] = v.MapIndex(keys[i]).Interface()
		}
	}

	return temp
}
