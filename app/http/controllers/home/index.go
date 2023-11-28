package home

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_tpl1/app"
	"web_tpl1/app/http/models"
)

func Index(ctx *gin.Context) {

	var rel []models.User

	err := app.DBW().Debug().Find(&rel).Error

	if err != nil {
		ctx.JSON(400, map[string]any{
			"code": 1,
			"msg":  "error",
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]any{
		"code": 0,
		"data": rel,
	})

}

func Redis(ctx *gin.Context) {

	rdb := app.RedisW()

	val, _ := rdb.Get(ctx, "aaa").Result()

	ctx.JSON(http.StatusOK, map[string]any{
		"code": 0,
		"data": val,
	})

}
