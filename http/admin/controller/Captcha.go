package controller

import "github.com/gin-gonic/gin"

type Captcha struct{}

func NewCaptcha() Captcha {
	return Captcha{}
}

// TODO:
func (a *Captcha) Get(c *gin.Context) {

}
