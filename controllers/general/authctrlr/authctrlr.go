// Package authctrlr 身份驗證
package authctrlr

import (
	"net/http"

	"github.com/WayneShenHH/toolsgo/models/viewmodels"

	"github.com/WayneShenHH/toolsgo/errno"
	"github.com/WayneShenHH/toolsgo/models"
	"github.com/WayneShenHH/toolsgo/repository"
	"github.com/WayneShenHH/toolsgo/services/authsvc"
	"github.com/gin-gonic/gin"
)

// Login 使用者登入
func Login(ctx *gin.Context) {
	var model models.UserLogin
	repo := repository.FromContext(ctx)
	err := ctx.Bind(&model)
	if err != nil {
		errno.Abort(errno.ErrBind, err, ctx)
		return
	}
	svc := authsvc.New(repo)
	result, err := svc.Login(model)
	if err != nil {
		errno.Abort(errno.ErrUserNotFound, err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, viewmodels.APIResult{
		Success: true,
		Data:    result,
	})
}

// swagger:operation POST /general/login general Login
//
// 使用者登入API
//
// 系統登入包含不同身份別都使用同一組登入方法
//
// ---
// parameters:
// - name: Username
//   in: query
//   description: 使用者名稱
//   type: string
//   required: true
// - name: Password
//   in: query
//   description: 密碼
//   type: string
//   required: true
// - name: Longitude
//   in: query
//   description: 經度
//   required: true
// - name: Latitude
//   in: query
//   description: 緯度
//   required: true
// - name: AuthType
//   in: query
//   description: 登入者角色
//   required: true
// responses:
//   200:
//     description: 成功訊息