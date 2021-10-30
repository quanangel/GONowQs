package controller

import (
	"nowqs/frame/http/blog/models"

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
// @Param Auth-Token header string true "Auth-Token"
// @Success 200 {object} _returnLoginGet
// @Failure 400 {object} _returnError
// @Router /blog/v1/login [get]
// Get is get user message
func (a *Login) Get(c *gin.Context) {

	returnData := gin.H{
		"code": -1,
	}
	authToken := c.GetHeader("Auth-Token")
	if authToken == "" {
		returnData["code"] = 3
		jsonHandle(c, returnData)
		return
	}

	userID := userAuth(authToken)
	if userID == 0 {
		returnData["code"] = 2
		jsonHandle(c, returnData)
		return
	}

	modelMember := models.NewUsers()
	userInfo := modelMember.GetByID(userID)
	if userInfo.Status != 1 {
		returnData["code"] = 20000
		jsonHandle(c, returnData)
		return
	}

	returnData["code"] = 0
	returnData["data"] = gin.H{
		"username":  userInfo.UserName,
		"nickname":  userInfo.NickName,
		"last_ip":   userInfo.LastIP,
		"last_time": userInfo.LastTime,
		"add_time":  userInfo.AddTime,
	}

	jsonHandle(c, returnData)
}

// @Summary Login
// @Tags Login
// @Description Login
// @Produce json
// @Param object query loginValidate false "put message"
// @Success 200 {object} _returnLoginPut
// @Failure 400 {object} _returnError
// @Router /blog/v1/login [put]
// Put is login function
func (a *Login) Put(c *gin.Context) {
	returnData := gin.H{
		"code": -1,
	}

	var validate loginValidate
	if err := c.Bind(&validate); err != nil {
		returnData["code"] = 10000
		returnData["msg"] = err.Error()
		jsonHandle(c, returnData)
		return
	}
	if validate.LastIP == "" {
		validate.LastIP = c.ClientIP()
	}
	modelMember := models.NewUsers()
	userInfo := modelMember.Login(validate.UserName, validate.Password, validate.LastIP)
	if userInfo.UserID == 0 {
		returnData["code"] = 20000
		jsonHandle(c, returnData)
		return
	}

	userToken := buildToken(userInfo.UserID)
	err := tokenSave(userInfo.UserID, userToken)
	if nil != err {
		returnData["code"] = 30000
		returnData["msg"] = err.Error()
		jsonHandle(c, returnData)
		return
	}

	returnData["code"] = 0
	returnData["data"] = userToken
	jsonHandle(c, returnData)
}
