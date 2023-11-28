package qrcode

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"
	"net/http"
)

func Gen(ctx *gin.Context) {

	var png []byte
	png, err := qrcode.Encode("https://www.baidu.com", qrcode.Medium, 256)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"code": 0,
			"msg":  "failed",
			"data": "whoos",
		})
	}

	ctx.JSON(http.StatusOK, map[string]any{
		"code": 0,
		"msg":  "success",
		"data": gin.H{
			"url": "data:image/png:base64," + base64.StdEncoding.EncodeToString(png),
		},
	})

}
