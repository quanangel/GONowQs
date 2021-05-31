package controller

import (
	"nowqs/frame/http/admin/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AuthGroup struct
type AuthGroup struct{}

// NewAuthGroup is AuthGroup exmaple
func NewAuthGroup() AuthGroup {
	return AuthGroup{}
}

// authGroupGetValidate is get validate struct
type authGroupGetValidate struct {
	// Type: list、only
	Type string `form:"type" json:"type" xml:"type" binding:"required,oneof=list only"`
	// Search: type is only the search is id, type is list the search is id/name/url
	Search string `form:"search" json:"search" xml:"search" binding:"required_if=Type only"`
	// Page
	Page int `form:"page" json:"page" xml:"page" binding:"-"`
	// Limit
	Limit int `form:"limit" json:"limit" xml:"limit" binding:"-"`
	// Order
	Order string `form:"order" json:"order" xml:"order" binding:"-"`
}

// authGroupPostValidate is post validate struct
type authGroupPostValidate struct {
	// Name
	Name string `form:"name" json:"name" xml:"name" binding:"required"`
	// Status: 1normal/2disable
	Status int8 `form:"status" json:"status" xml:"status" binding:"oneof=1 2"`
	// Rules
	Rules string `form:"rules" json:"rules" xml:"rules" binding:"-"`
}

// authGroupPutValidate is put validate struct
type authGroupPutValidate struct {
	authGroupDeleteValidate
	authGroupPostValidate
}

// authGroupDeleteValidate is delete validate struct
type authGroupDeleteValidate struct {
	// ID
	ID int `form:"id" json:"id" xml:"id" binding:"required"`
}

// @Summary AuthGroup
// @Tags AuthGroup
// @Description auth group
// @Produce json
// @Param Auth-Token header string true "Auth-Token"
// @Param object query authGroupGetValidate false "get message"
// @Success 200-1 {object} _returnNavGetList
// @Success 200-2 {object} _returnNavGetOnly
// @Failure 400 {object} _returnError
// @Router /admin/auth_group/index [get]
// Get is get auth group message
func (a *AuthGroup) Get(c *gin.Context) {
	returnData := gin.H{
		"code": -1,
	}
	isPower := checkRuleByUser(c)
	if !isPower {
		returnData["code"] = 2
		jsonHandle(c, returnData)
		return
	}
	var validate authGroupGetValidate
	if err := c.Bind(&validate); err != nil {
		returnData["code"] = 10000
		returnData["msg"] = err.Error()
		jsonHandle(c, returnData)
		return
	}

	model := models.NewAuth()
	search := make(map[string]interface{})
	switch validate.Type {
	case "list":
		if validate.Search != "" {
			search["id"] = validate.Search
			search["name"] = validate.Search
		}
		if validate.Page == 0 {
			validate.Page = 1
		}
		if validate.Limit == 0 {
			validate.Limit = 20
		}
		result := model.GetGroupList(search, validate.Page, validate.Limit)
		if len(result) == 0 {
			returnData["code"] = 6
		} else {
			returnData["code"] = 0
			returnData["data"] = result
		}
	case "only":
		searchId, err := strconv.Atoi(validate.Search)
		if err != nil {
			returnData["code"] = 1
			returnData["msg"] = err.Error()
			jsonHandle(c, returnData)
			return
		}
		result := model.GetGroupByID(searchId)
		if result.ID == 0 {
			returnData["code"] = 6
		} else {
			returnData["code"] = 0
			returnData["data"] = result
		}
	}

	jsonHandle(c, returnData)
}

// @Summary AuthGroup
// @Tags AuthGroup
// @Description auth group
// @Produce json
// @Param Auth-Token header string true "Auth-Token"
// @Param object query authGroupPostValidate false "post message"
// @Success 200 {object} _returnAuthGroupPost
// @Failure 400 {object} _returnError
// @Router /admin/auth_group/index [post]
// Post is add auth group message
func (a *AuthGroup) Post(c *gin.Context) {
	returnData := gin.H{
		"code": -1,
	}
	isPower := checkRuleByUser(c)
	if !isPower {
		returnData["code"] = 2
		jsonHandle(c, returnData)
		return
	}
	var validate authGroupPostValidate
	if err := c.Bind(&validate); err != nil {
		returnData["code"] = 10000
		returnData["msg"] = err.Error()
		jsonHandle(c, returnData)
		return
	}
	model := models.NewAuth()
	result := model.AddGroup(validate.Name, validate.Status, validate.Rules)
	if result == 0 {
		returnData["code"] = 1
	} else {
		returnData["code"] = 0
		returnData["data"] = result
	}

	jsonHandle(c, returnData)
}

// @Summary AuthGroup
// @Tags AuthGroup
// @Description auth group
// @Produce json
// @Param Auth-Token header string true "Auth-Token"
// @Param object query authGroupPutValidate false "put message"
// @Success 200 {object} _returnSuccess
// @Failure 400 {object} _returnError
// @Router /admin/auth_group/index [put]
// Put is edit auth group message
func (a *AuthGroup) Put(c *gin.Context) {
	returnData := gin.H{
		"code": -1,
	}
	isPower := checkRuleByUser(c)
	if !isPower {
		returnData["code"] = 2
		jsonHandle(c, returnData)
		return
	}
	var validate authGroupPutValidate
	if err := c.Bind(&validate); err != nil {
		returnData["code"] = 10000
		returnData["msg"] = err.Error()
		jsonHandle(c, returnData)
		return
	}
	model := models.NewAuth()
	search := make(map[string]interface{})
	search["id"] = validate.ID
	data := make(map[string]interface{})
	data["name"] = validate.Name
	data["status"] = validate.Status
	if validate.Rules != "" {
		data["rules"] = validate.Rules
	}
	result := model.EditGroup(search, data)
	if result {
		returnData["code"] = 0
	} else {
		returnData["code"] = 1
	}

	jsonHandle(c, returnData)
}

// @Summary AuthGroup
// @Tags AuthGroup
// @Description auth group
// @Produce json
// @Param Auth-Token header string true "Auth-Token"
// @Param object query authGroupDeleteValidate false "delete message"
// @Success 200 {object} _returnSuccess
// @Failure 400 {object} _returnError
// @Router /admin/auth_group/index [delete]
// Delete is delete auth group message
func (a *AuthGroup) Delete(c *gin.Context) {
	returnData := gin.H{
		"code": -1,
	}
	isPower := checkRuleByUser(c)
	if !isPower {
		returnData["code"] = 2
		jsonHandle(c, returnData)
		return
	}
	var validate authGroupDeleteValidate
	if err := c.Bind(&validate); err != nil {
		returnData["code"] = 10000
		returnData["msg"] = err.Error()
		jsonHandle(c, returnData)
		return
	}
	model := models.NewAuth()
	search := make(map[string]interface{})
	search["id"] = validate.ID
	result := model.DelGroup(search)
	if result {
		returnData["code"] = 0
	} else {
		returnData["code"] = 1
	}

	jsonHandle(c, returnData)
}
