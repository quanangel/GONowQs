package controller

import (
	"nowqs/frame/http/admin/models"

	"github.com/gin-gonic/gin"
)

// Login struct
type Login struct{}

// NewLogin is login exmaple function
func NewLogin() Login {
	return Login{}
}

// TODO: mebe need cpatcha code
type loginValidate struct {
	UserName   string `form:"username" json:"username" xml:"username" binding:"required"`
	Password   string `form:"password" json:"password" xml:"password" binding:"required"`
	LastIP     string `form:"last_ip" json:"last_ip" xml:"last_ip" binding:"-"`
	Cpatcha    string `form:"cpatcha" json:"cpatcha" xml:"cpatcha" binding:"-"`
	CpatchaMd5 string `form:"cpatcha_md5" json:"cpatcha_md5" xml:"cpatcha_md5" binding:"-"`
}

// Put is login function
func (a *Login) Put(c *gin.Context) {
	returnData := gin.H{
		"code": -1,
		"msg":  errorMsg(-1),
	}

	var validate loginValidate
	if err := c.Bind(&validate); err != nil {
		returnData["code"] = 10000
		c.JSON(jsonHandle(returnData))
		return
	}
	if validate.LastIP == "" {
		validate.LastIP = c.ClientIP()
	}
	modelMember := models.NewMember()
	userInfo := modelMember.Login(validate.UserName, validate.Password, validate.LastIP)
	if userInfo.UserID == 0 {
		returnData["code"] = 20000
		c.JSON(jsonHandle(returnData))
		return
	}

	c.JSON(jsonHandle(returnData))
	return
}
