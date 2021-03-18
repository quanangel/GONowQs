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

// @Summary Login
// @Tags Login
// @Description Login
// @Produce json
// @Param username query string true "username"
// @Param password query string true "password"
// @Param last_ip query string false "last_ip"
// @Param cpatcha query string false "cpatcha"
// @Param cpatcha_md5 query string false "cpatcha_md5"
// @Success 200 {string} json "{"code": 0,"msg": "success","data": ""}"
// @Failure 400 {string} json "{"code": 1, "msg": "error"}"
// @Router /admin/login/index [put]
// Put is login function
func (a *Login) Put(c *gin.Context) {
	returnData := gin.H{
		"code": -1,
	}

	var validate loginValidate
	if err := c.Bind(&validate); err != nil {
		returnData["code"] = 10000
		returnData["msg"] = err.Error()
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

	userToken := buildToken(userInfo.UserID)
	err := tokenSave(userInfo.UserID, userToken)
	if nil != err {
		returnData["code"] = 30000
		returnData["msg"] = err.Error()
		c.JSON(jsonHandle(returnData))
		return
	}

	returnData["code"] = 0
	returnData["data"] = userToken
	c.JSON(jsonHandle(returnData))
	return
}

// @Summary Login
// @Tags Login
// @Description Login
// @Produce json
// @Param Auth-Token header string true "Auth-Token"
// @Success 200 {string} json "{"code": 0,"msg": "success","data": ""}"
// @Failure 400 {string} json "{"code": 1, "msg": "error"}"
// @Router /admin/login/index [get]
func (a *Login) Get(c *gin.Context) {
	checkRuleByUser(c)

	returnData := gin.H{
		"code": -1,
	}
	authToken := c.GetHeader("Auth-Token")
	if "" == authToken {
		returnData["code"] = 3
		c.JSON(jsonHandle(returnData))
		return
	}

	userID := userAuth(authToken)
	if 0 == userID {
		returnData["code"] = 2
		c.JSON(jsonHandle(returnData))
		return
	}

	modelMember := models.NewMember()
	userInfo := modelMember.GetByID(userID)
	if 1 != userInfo.Status {
		returnData["code"] = 20000
		c.JSON(jsonHandle(returnData))
		return
	}

	returnData["code"] = 0
	returnData["data"] = gin.H{
		"username":      userInfo.UserName,
		"nickname":      userInfo.NickName,
		"last_ip":       userInfo.LastIP,
		"last_time":     userInfo.LastTime,
		"register_time": userInfo.RegisterTime,
	}

	c.JSON(jsonHandle(returnData))
	return
}
