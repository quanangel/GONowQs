package controller

import (
	"nowqs/frame/config"
	"nowqs/frame/models/mysql"
	"nowqs/frame/models/redis"

	"github.com/gin-gonic/gin"
)

type Captcha struct{}

func NewCaptcha() Captcha {
	return Captcha{}
}

// @Summary Catpcha
// @Tags Catpcha
// @Description Catpcha
// @Produce json
// @Success 200 {object} _returnCaptcha
// @Failure 400 {object} _returnError
// @Router /blog/v1/captcha [get]
// Get is get captcha message
func (a *Captcha) Get(c *gin.Context) {
	returnData := gin.H{
		"code": -1,
	}
	if config.AppConfig.Redis.Status {
		captcha, err := redis.GetCaptcha()
		if err != nil {
			returnData["code"] = 1
			returnData["msg"] = err.Error()
		} else {
			returnData["code"] = 0
			returnData["data"] = gin.H{
				"key":         captcha.CodeKey,
				"image":       captcha.Image,
				"expire_time": captcha.ExpireTime,
			}
		}
	} else {
		captcha := mysql.NewCaptcha()
		result, err := captcha.Add()
		if err != nil {
			returnData["code"] = 1
			returnData["msg"] = err.Error()
		} else {
			returnData["code"] = 0
			returnData["data"] = result
		}
	}

	jsonHandle(c, returnData)
}
