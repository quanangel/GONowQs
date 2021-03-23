package controller

import (
	"nowqs/frame/http/admin/models"

	"github.com/gin-gonic/gin"
)

// Nav is struct
type Nav struct{}

// NewNav is nav exmaple function
func NewNav() Nav {
	return Nav{}
}

// navGetValidate is get validate struct
type navGetValidate struct {
	// type: list、only
	Type string `form:"type" json:"type" xml:"type" binding:"required,oneof=list only"`
	// search: type is only the search is id, type is list the search is id/name/url
	Search string `form:"search" json:"search" xml:"search" binding:"required_if=Type only"`
	// page
	Page int `form:"page" json:"page" xml:"page" binding:"-"`
	// limit
	Limit int `form:"limit" json:"limit" xml:"limit" binding:"-"`
}

// navPostValidate is post validate struct
type navPostValidate struct {
	// name
	Name string `form:"name" json:"name" xml:"name" binding:"required"`
	// pid
	PID int `form:"pid" json:"pid" xml:"pid" binding:"omitempty"`
	// url
	Url string `form:"url" json:"url" xml:"url" binding:"required"`
	// status: 1normal、2disable
	Status int8 `form:"status" json:"status" xml:"status" binding:"oneof=1 2"`
}

type navPutValidate struct {
	// id
	ID int `form:"id" json:"id" xml:"id" binding:"required"`
	// pid
	PID string `form:"pid" json:"pid" xml:"pid" binding:"omitempty"`
	// name
	Name string `form:"name" json:"name" xml:"name" binding:"omitempty"`
	// url
	Url string `form:"url" json:"url" xml:"url" binding:"omitempty"`
	// status: 1normal、2disable
	Status string `form:"status" json:"status" xml:"status" binding:"omitempty"`
}

// @Summary Nav
// @Tags Nav
// @Description admin nav
// @Produce json
// @Param Auth-Token header string true "Auth-Token"
// @Param object query navGetValidate false "get message"
// @Success 200-1 {object} _returnNavGetList
// @Success 200-2 {object} _returnNavGetOnly
// @Failure 400 {object} _returnError
// @Router /admin/nav/index [get]
// Get is get nav message
func (a *Nav) Get(c *gin.Context) {
	returnData := gin.H{
		"code": -1,
	}
	isPower := checkRuleByUser(c)
	if !isPower {
		returnData["code"] = 2
		c.JSON(jsonHandle(returnData))
		return
	}

	var validate navGetValidate
	if err := c.Bind(&validate); err != nil {
		returnData["code"] = 10000
		returnData["mes"] = err.Error()
		c.JSON(jsonHandle(returnData))
		return
	}

	model := models.NewAdminNav()
	search := make(map[string]interface{})
	switch validate.Type {
	case "list":
		if "" != validate.Search {
			search["id"] = validate.Search
			search["name"] = validate.Search
			search["url"] = validate.Search
		}
		if 0 == validate.Page {
			validate.Page = 1
		}
		if 0 == validate.Limit {
			validate.Limit = 20
		}
		result := model.GetList(search, validate.Page, validate.Limit)
		if 0 == len(*result) {
			returnData["code"] = 6

		} else {
			returnData["code"] = 0
			returnData["data"] = result
		}
	case "only":
		search["id"] = validate.Search
		result := model.GetOne(search)
		if 0 == result.ID {
			returnData["code"] = 6
		} else {
			returnData["code"] = 0
			returnData["data"] = result
		}
	}

	c.JSON(jsonHandle(returnData))
	return
}

// @Summary Nav
// @Tags Nav
// @Description admin
// @Produce json
// @Param Auth-Token header string true "Auth-Token"
// @Param object query navPostValidate false "post message"
// @Success 200 {object} _returnSuccess
// @Failure 400 {object} _returnError
// @Router /admin/nav/index [post]
// Post is post nav message
func (m *Nav) Post(c *gin.Context) {
	returnData := gin.H{
		"code": -1,
	}

	// isPower := checkRuleByUser(c)
	// if !isPower {
	// 	returnData["code"] = 2
	// 	c.JSON(jsonHandle(returnData))
	// 	return
	// }

	var validate navPostValidate
	if err := c.Bind(&validate); err != nil {
		returnData["code"] = 10000
		returnData["mes"] = err.Error()
		c.JSON(jsonHandle(returnData))
		return
	}

	model := models.NewAdminNav()
	result := model.Add(validate.Name, validate.PID, validate.Url, validate.Status)
	if 0 == result {
		returnData["code"] = 1
		c.JSON(jsonHandle(returnData))
		return
	}

	returnData["code"] = 0
	c.JSON(jsonHandle(returnData))
	return
}

// @Summary Nav
// @Tags Nav
// @Description admin nav
// @Produce json
// @Param Auth-Token header string true "Auth-Token"
// @Param object query navPutValidate false "put message"
// @Success 200 {object} _returnSuccess
// @Failure 400 {string} _returnError
// @Router /admin/nav/index [put]
// Put is put nav message
func (m *Nav) Put(c *gin.Context) {
	returnData := gin.H{
		"code": -1,
	}

	// isPower := checkRuleByUser(c)
	// if !isPower {
	// 	returnData["code"] = 2
	// 	c.JSON(jsonHandle(returnData))
	// 	return
	// }

	var validate navPutValidate

	if err := c.Bind(&validate); err != nil {
		returnData["code"] = 1
		c.JSON(jsonHandle(returnData))
		return
	}
	model := models.NewAdminNav()
	search := make(map[string]interface{})
	search["id"] = validate.ID
	update := make(map[string]interface{})
	if "" != validate.Name {
		update["name"] = validate.Name
	}
	if "" != validate.PID {
		update["pid"] = validate.PID
	}
	result := model.Edit(search, update)
	if result {
		returnData["code"] = 0
	} else {
		returnData["code"] = 1
	}
	c.JSON(jsonHandle(returnData))
	return
}
