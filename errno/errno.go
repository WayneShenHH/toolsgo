package errno

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/WayneShenHH/toolsgo/models/viewmodels"
	"github.com/gin-gonic/gin"
)

const (
	// CodeSuccess success
	CodeSuccess = 0
	// CodeFailed failed
	CodeFailed = 1
	// CodeAuthError authorize error
	CodeAuthError = 2
	// CodeValidError validation error
	CodeValidError = 3
	// CodeServerError server error
	CodeServerError = 9
)

// Err represents an error, `Path`, `Line`, `Code` will be automatically filled.
type Err struct {
	ResultCode uint
	Code       string
	Message    string
	StatusCode int
	Path       string
	Line       int
}

// Error returns the error message.
func (e *Err) Error() string {
	return e.Code
}

// Fill the error struct with the detail error information.
func Fill(err *Err) *Err {
	_, fn, line, _ := runtime.Caller(1)

	// Fill the error occurred path, line, code.
	err.Path = strings.Replace(fn, os.Getenv("GOPATH"), "", -1)
	err.Line = line
	return err
}

// DefaultSuccess return default success APIResult
func DefaultSuccess(data interface{}) viewmodels.APIResult {
	return viewmodels.APIResult{
		Code:    CodeSuccess,
		Success: true,
		Data:    data,
	}
}

// Success return default success result
func Success(data interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, DefaultSuccess(data))
}

// SuccessPager return default success result with total count info for paging
func SuccessPager(data interface{}, count uint, ctx *gin.Context) {
	model := DefaultSuccess(data)
	model.Count = count
	ctx.JSON(http.StatusOK, model)
}

// DefaultFailed return default failed APIResult
func DefaultFailed(err Err) viewmodels.APIResult {
	return viewmodels.APIResult{
		Code:    err.ResultCode,
		Success: false,
		Data:    nil,
		Message: err.Code,
	}
}

// Abort the current request with the specified error code.
func Abort(errStruct *Err, err error, c *gin.Context) {
	if err != nil && errStruct != nil {
		fmt.Println(`[Abort] `, errStruct.Code, err.Error())
	}
	c.Error(err)
	c.Error(Fill(errStruct))
	c.AbortWithStatus(errStruct.StatusCode)
}
