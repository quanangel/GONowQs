package controller

import (
	"nowqs/frame/http/admin/models"

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

// type authGroupPostValidate struct {
// }

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
	// TODO:
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
	}

	jsonHandle(c, returnData)
}

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
	// TODO:

	jsonHandle(c, returnData)
}

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
	// TODO:

	jsonHandle(c, returnData)
}

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
	// TODO:

	jsonHandle(c, returnData)
}
