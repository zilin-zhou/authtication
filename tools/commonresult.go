package tools

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//操作成功和失败的封装
const (
	SUCCESS int = 0 //操作成功
	FAILED  int = 1 //操作失败
)

//操作成功函数
func Success(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": SUCCESS,
		"msg":  "操作成功",
		"data": v,
	})
}

//操作失败函数
func Failed(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": FAILED,
		"msg":  v,
	})
}
