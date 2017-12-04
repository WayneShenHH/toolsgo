package middleware

import (
	"github.com/WayneShenHH/toolsgo/app"
	"github.com/WayneShenHH/toolsgo/models/entities"
	"github.com/WayneShenHH/toolsgo/repository"

	"github.com/WayneShenHH/toolsgo/tools/auth"

	"github.com/WayneShenHH/toolsgo/errno"
	"github.com/gin-gonic/gin"
)

type authRole int

const (
	// Operator 控盤人員權限
	Operator authRole = iota
	// Agent 代理權限
	Agent authRole = iota
	// Player 玩家權限
	Player authRole = iota
)

//AuthMiddleware middleware
func AuthMiddleware(role authRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		validateToken(c, role)
		c.Next()
	}
}

// validateToken 判斷驗證狀況回傳 http header
// TODO 修改為 JWT
func validateToken(c *gin.Context, role authRole) {

	token := getToken(c)
	// fmt.Println("X-Auth-Token:" + token)
	repo := repository.FromContext(c)
	if token == "" {
		// 執行測試過程中，跳過登入的檢查
		if app.Env() == "test" || app.Env() == "ci" {
			return
		}
		// c.AbortWithStatus(401)
		errno.Abort(errno.ErrToken, nil, c)
	}
	user, err := checkToken(repo, token, role)
	if err == nil {
		auth.ToContext(c, auth.User{
			ID:       user.ID,
			Username: user.Username,
			Tier:     user.Tier,
		})
		c.Next()
	} else {
		// c.AbortWithStatus(401)
		errno.Abort(errno.ErrToken, err, c)
	}
	return
}

// getToken 從封包表頭(http)或是url(websocket)中取得token
func getToken(c *gin.Context) string {
	if token := c.Request.Header.Get("X-Auth-Token"); token != "" {
		return token
	}
	return c.Query("X-Auth-Token")
}

// checkToken 根據傳入的身份別判斷 token 正確性
/**
operator role => user table admin:true
agent role => user table tier 1 to 7
player role => user table tier 8
**/
func checkToken(repo repository.Repository, token string, role authRole) (*entities.User, error) {
	return repo.GetUserByToken(token)
}
