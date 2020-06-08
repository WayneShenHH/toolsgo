package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/WayneShenHH/toolsgo/pkg/environment"
	"github.com/WayneShenHH/toolsgo/pkg/errors"
	"github.com/WayneShenHH/toolsgo/pkg/module/jose"
	"github.com/WayneShenHH/toolsgo/pkg/module/logger"
)

type authRole uint8

const (
	// Operator 控盤人員權限
	Operator authRole = 0
	// Station 子站權限
	Station authRole = 1
	// Player 玩家權限
	Player authRole = 2
)

// 原本使用的名稱是 X-Auth-Token 為了標準化改為使用 Authorization 並且加上 Bearer 開頭
const (
	xAuthToken = "X-Auth-Token" //nolint
	tokenName  = "Authorization"
	prefix     = "Bearer "
)

//AuthMiddleware middleware
func AuthMiddleware(role authRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		validateToken(c, role)
		c.Next()
	}
}

// validateToken 判斷驗證狀況回傳 http header
func validateToken(c *gin.Context, role authRole) {
	token := getToken(c)
	if token == "" {
		// errno.Abort(errno.ErrToken, errors.E("token not found"), c)
		c.Error(errors.E("token not found")) //nolint
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	testConfig := environment.Setting.Test
	if testConfig.Enable && token == testConfig.Token {
		// 執行測試過程中，跳過登入的檢查
		logger.Debug(`test mode, use mock auth user`)
		//mockAuthUserForTest(c, repo, role)
		return
	}
	user, err := checkAuth(token, role)
	if err != nil {
		c.Error(err) //nolint
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if ok := (user.Status == jose.ValidStatus); !ok { // 檢查帳號狀態是否正常
		errors.E("user invalid")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	jose.ToContext(c, user)
	c.Next()
}

// getToken 從封包表頭(http)或是url(websocket)中取得token
func getToken(c *gin.Context) string {
	token := tryFindToken(c, tokenName)
	if token == "" {
		// 未來將不支援 X-Auth-Token
		token = tryFindToken(c, xAuthToken)
		logger.Debug(xAuthToken + token)
	} else {
		logger.Debug(tokenName + token)
	}
	return strings.TrimPrefix(token, prefix)
}

// 試著從 Request 、 From 、 QueryString 中找出 Token 字串，找不到時回傳空字串
func tryFindToken(ctx *gin.Context, key string) string {
	token := ctx.Request.Header.Get(key)
	if token == "" {
		token = ctx.PostForm(key)
	}
	if token == "" {
		token = ctx.Query(key)
	}
	return token
}

// 根據傳入的身份別判斷 token 正確性
func checkAuth(token string, role authRole) (*jose.UserToken, error) {
	// 判斷個別等級權限對應功能
	ut, err := authJWT(token)
	if err != nil {
		return ut, errors.E("Unauthorized")
	}
	switch role {
	case Operator:
		if ut.OwnerType == "operators" {
			return ut, nil
		}
	case Station:
		if ut.OwnerType == "stations" {
			return ut, nil
		}
	case Player:
		if ut.OwnerType == "players" {
			return ut, nil
		}
	}
	return ut, errors.E("Unauthorized")
}

/*
// 驗證 JWT，回傳使用者資訊
func authJWT(cache memcache.MemCache, jwt string) (*jose.UserToken, error) {
	const op errors.Op = "middlewares/authJWT"

	jwt = strings.Replace(jwt, "Bearer ", "", 1)
	user := new(jose.UserToken)
	if err := jose.Bind(jwt, user); err != nil {
		return nil, errors.E(op, err)
	}

	if user.Status != jose.ValidStatus {
		return nil, errors.E(op, errors.Data, errors.Token, "user is not in valid status")
	}
	cacheJwt, err := cache.GETSTR(jose.CacheKey(user.ID))
	if err != nil {
		return nil, errors.E(op, errors.Datastore, err)
	}

	if len(cacheJwt) == 0 {
		return nil, errors.E(op, errors.Data, errors.Token, fmt.Sprintf("token not exist on redis, user.ID=%d", user.ID))
	}

	if cacheJwt != jwt {
		return nil, errors.E(op, errors.Data, errors.Token, "token not valid")
	}
	return user, nil
}*/

// 驗證 JWT，回傳使用者資訊（統一層暫時不驗證 redis token）
func authJWT(jwt string) (*jose.UserToken, error) {
	const op errors.Op = "middlewares/authJWT"

	jwt = strings.Replace(jwt, "Bearer ", "", 1)
	user := new(jose.UserToken)
	if err := jose.Bind(jwt, user); err != nil {
		return nil, errors.E(op, err)
	}

	if user.Status != jose.ValidStatus {
		return nil, errors.E(op, errors.Data, errors.Token, "user is not in valid status")
	}

	return user, nil
}
