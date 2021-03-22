package controller

import "github.com/gin-gonic/gin"

// Nav is struct
type Nav struct{}

// NewNav is nav exmaple function
func NewNav() Nav {
	return Nav{}
}

// navGetValidate is get validate struct
type navGetValidate struct {
	Type   string `form:"type" json:"type" xml:"type" binding:"required"`
	Search string `form:"search" json:"search" xml:"search" binding:"-"`
	Limit  int    `form:"limit" json:"limit" xml:"limit" binding:"-"`
	Page   int    `form:"page" json:"page" xml:"page" binding:"-"`
}

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
	if err := c.Bind(validate); err != nil {
		returnData["code"] = 10000
		returnData["mes"] = err.Error()
		c.JSON(jsonHandle(returnData))
		return
	}

	// TODO: not finish

	c.JSON(jsonHandle(returnData))
	return
}
