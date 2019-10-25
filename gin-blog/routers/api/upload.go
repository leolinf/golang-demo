package api

import (
	"gin-blog/pkg/e"
	"gin-blog/pkg/logging"
	"gin-blog/pkg/setting"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	code := e.SUCCESS
	data := make(map[string]interface{})
	file, image, err := c.Request.FormFile("image")
	logging.Info(file, image, err)
	if err != nil {
		logging.Warn(err)
		code = e.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
		return
	}
	if image == nil {
		code = e.INVALID_PARAMS
	} else {
		imageName := image.Filename
		if err := c.SaveUploadedFile(image, setting.AppSetting.RuntimeRootPath+"test.py"); err != nil {
			code = e.INVALID_PARAMS
			logging.Warn("Save images err %v", err)
		}
		logging.Info(file)
		logging.Info(imageName)
		logging.Info(image)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
