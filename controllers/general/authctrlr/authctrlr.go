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

/*Login swagger:route POST /general/login  general UserLogin
系統登入包含不同身份別都使用同一組登入方法
*/
func Login(ctx *gin.Context) {
	// user := auth.FromContext(ctx)
	// fmt.Println(user.Username)
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

	// usersvc := usersvc.New(repo, auth.User{ID: result.User.ID})
	//usersvc.CalculateQuota()

	ctx.JSON(http.StatusOK, viewmodels.APIResult{
		Success: true,
		Data:    result,
	})
}
