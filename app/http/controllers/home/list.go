package home

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func List(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]any{
		"code": 0,
		"msg":  "success",
		"data": "hello,list",
	})
}
