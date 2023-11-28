package http

import (
	"github.com/gin-gonic/gin"
	"web_tpl1/app"
	"web_tpl1/app/http/routes"
)

func NewServer() error {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routes.Reg(r)
	return r.Run(app.Config.HTTP.Listen) // listen and serve on 0.0.0.0:8080

}
