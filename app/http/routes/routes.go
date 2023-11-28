package routes

import (
	"github.com/gin-gonic/gin"
	"web_tpl1/app/http/controllers/home"
	"web_tpl1/app/http/controllers/qrcode"
)

func Reg(r *gin.Engine) {

	r.GET("/", home.Index)
	r.GET("/redis", home.Redis)
	r.GET("/list", home.List)
	r.GET("/v1/qrcode", qrcode.Gen)
}
