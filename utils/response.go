package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(ctx *gin.Context, code int, message string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  message,
		"data": data,
	})
}
func Response1(ctx *gin.Context, code int, message string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  message,
	})
}

// Success 成功的请求
func Success(ctx *gin.Context, data interface{}) {
	Response(ctx, 200, "ok", data)
}

func Success1(ctx *gin.Context) {
	Response1(ctx, 200, "ok")
}

// Fail 失败的请求
func Fail(ctx *gin.Context, message string) {
	Response(ctx, 400, message, nil)
}
