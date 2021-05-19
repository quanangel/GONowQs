package controller

import (
	"nowqs/frame/http/blog/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BlogClassify struct{}

func NewBlogClassify() BlogClassify {
	return BlogClassify{}
}

// blogClassifyGetValidate is get validate struct
type blogClassifyGetValidate struct {
	// Type my/list/only
	Type string `form:"type" json:"type" xml:"type" binding:"required,oneof=my list only" default:"list"`
	// Search type is only the search is id, type is list the search is id/name
	Search string `form:"search" json:"search" xml:"search" binding:"required_if=Type only"`
	// Classify
	Classify int64 `form:"classify" json:"classify" xml:"classify" binding:"-"`
	// Page
	Page int `form:"page" json:"page" xml:"page" binding:"-"`
	// Limit
	Limit int `form:"limit" json:"limit" xml:"limit" binding:"-"`
	// Order
	Order string `form:"order" json:"order" xml:"order" binding:"-" default:"order_id asc, id desc"`
}

// blogClassifyPostValidate is post validate struct
type blogClassifyPostValidate struct {
	// Name
	Name string `form:"name" json:"name" xml:"name" binding:"required"`
	// PID
	PID int64 `form:"pid" json:"pid" xml:"pid" binding:"required"`
	// Type 1markdown/2quill
	Type int8 `form:"type" json:"type" xml:"type" binding:"required" default:"1"`
	// Status 1public/2privarte/3draft
	Status int8 `form:"status" json:"status" xml:"status" binding:"required" default:"1"`
	// OrderID
	OrderID int64 `form:"order_id" json:"order_id" xml:"order_id" binding:"required" default:"0"`
}

// blogClassifyPutValidate is put validate struct
type blogClassifyPutValidate struct {
	// ID
	ID int64 `form:"id" json:"id" xml:"id" binding:"required"`
	blogClassifyPostValidate
}

// blogClassifyDeleteValidate is delete validate struct
type blogClassifyDeleteValidate struct {
	// ID
	ID int64 `form:"id" json:"id" xml:"id" binding:"required"`
}

// @Summary BlogClassify
// @Tags BlogClassify
// @Description BlogClassify
// @Produce json
// @Param Auth-Token header string false "Auth-Token"
// @Param object query blogClassifyGetValidate false "get message"
// @Success 200-1 {object} _returnBlogClassifyGetList
// @Success 200-2 {object} _returnBlogClassifyGetOnly
// @Failure 400 {object} _returnError
// @Router /blog/v1/blog_classify [get]
// Get is get blog classify message
func (a *BlogClassify) Get(c *gin.Context) {
	returnData := gin.H{
		"code": -1,
	}
	authToken := c.GetHeader("Auth-Token")
	userID := userAuth(authToken)

	var validate blogClassifyGetValidate
	if err := c.Bind(&validate); err != nil {
		returnData["code"] = 10000
		returnData["msg"] = err.Error()
		jsonHandle(c, returnData)
		return
	}
	modelBlogClassify := models.NewBlogClassify()
	search := make(map[string]interface{})
	if validate.Page == 0 {
		validate.Page = 1
	}
	if validate.Limit == 0 {
		validate.Limit = 20
	}
	if validate.Classify != 0 {
		search["classify_id"] = validate.Classify
	}
	switch validate.Type {
	case "my":
		if userID == 0 {
			returnData["code"] = 2
			jsonHandle(c, returnData)
			return
		}
		if validate.Search != "" {
			search["id"] = validate.Search
			search["name"] = validate.Search
		}
		search["user_id"] = userID
		total, result := modelBlogClassify.GetList(search, validate.Page, validate.Limit, validate.Order)
		returnData["code"] = 0
		returnData["data"] = gin.H{
			"total": total,
			"page":  validate.Page,
			"limit": validate.Limit,
			"data":  result,
		}

	case "list":
		if validate.Search != "" {
			search["id"] = validate.Search
			search["name"] = validate.Search
		}
		search["status"] = 1
		total, result := modelBlogClassify.GetList(search, validate.Page, validate.Limit, validate.Order)
		returnData["code"] = 0
		returnData["data"] = gin.H{
			"total": total,
			"page":  validate.Page,
			"limit": validate.Limit,
			"data":  result,
		}
	case "only":
		id, _ := strconv.ParseInt(validate.Search, 10, 64)
		result := modelBlogClassify.GetByID(id, userID)
		if result == nil {
			returnData["code"] = 6
		} else {
			returnData["code"] = 0
			returnData["data"] = result
		}
	default:
		jsonHandle(c, returnData)
		return
	}
	jsonHandle(c, returnData)
}

// @Summary BlogClassify
// @Tags BlogClassify
// @Description BlogClassify
// @Produce json
// @Param Auth-Token header string true "Auth-Token"
// @Param object query blogClassifyPostValidate false "post message"
// @Success 200 {object} _returnSuccess
// @Failure 400 {object} _returnError
// @Router /blog/v1/blog_classify [post]
// Post is add blog classify message
func (a *BlogClassify) Post(c *gin.Context) {
	returnData := gin.H{
		"code": -1,
	}
	userID := userAuth(c.GetHeader("Auth-Token"))
	if userID == 0 {
		returnData["code"] = 2
		jsonHandle(c, returnData)
		return
	}
	var validate blogClassifyPostValidate
	if err := c.Bind(&validate); err != nil {
		returnData["code"] = 10000
		returnData["msg"] = err.Error()
		jsonHandle(c, returnData)
		return
	}
	modelBlogClassify := models.NewBlogClassify()
	result := modelBlogClassify.Add(validate.PID, userID, validate.Name, validate.Type, validate.Status, validate.OrderID)
	if result == 0 {
		returnData["code"] = 1
	} else {
		returnData["code"] = 0
	}

	jsonHandle(c, returnData)
}

// @Summary BlogClassify
// @Tags BlogClassify
// @Description BlogClassify
// @Produce json
// @Param Auth-Token header string true "Auth-Token"
// @Param object query blogClassifyPutValidate false "put message"
// @Success 200 {object} _returnSuccess
// @Failure 400 {object} _returnError
// @Router /blog/v1/blog_classify [put]
// Put is edit blog classify message
func (a *BlogClassify) Put(c *gin.Context) {
	returnData := gin.H{
		"code": -1,
	}
	userID := userAuth(c.GetHeader("Auth-Token"))
	if userID == 0 {
		returnData["code"] = 2
		jsonHandle(c, returnData)
		return
	}
	var validate blogClassifyPutValidate
	if err := c.Bind(&validate); err != nil {
		returnData["code"] = 10000
		returnData["msg"] = err.Error()
		jsonHandle(c, returnData)
		return
	}
	modelBlogClassify := models.NewBlogClassify()
	search := make(map[string]interface{})
	search["id"] = validate.ID
	search["user_id"] = userID

	editData := make(map[string]interface{})
	editData["name"] = validate.Name
	editData["pid"] = validate.PID

	result := modelBlogClassify.Edit(search, editData)
	if result {
		returnData["code"] = 0
	} else {
		returnData["code"] = 1
	}
	jsonHandle(c, returnData)
}

// @Summary BlogClassify
// @Tags BlogClassify
// @Description BlogClassify
// @Produce json
// @Param Auth-Token header string true "Auth-Token"
// @Param object query blogClassifyDeleteValidate false "delete message"
// @Success 200 {object} _returnSuccess
// @Failure 400 {object} _returnError
// @Router /blog/v1/blog_classify [delete]
// Delete is delete blog classify message
func (a *BlogClassify) Delete(c *gin.Context) {
	returnData := gin.H{
		"code": -1,
	}
	userID := userAuth(c.GetHeader("Auth-Token"))
	if userID == 0 {
		returnData["code"] = 2
		jsonHandle(c, returnData)
		return
	}
	var validate blogClassifyDeleteValidate
	if err := c.Bind(&validate); err != nil {
		returnData["code"] = 10000
		returnData["msg"] = err.Error()
		jsonHandle(c, returnData)
		return
	}

	modelBlogClassify := models.NewBlogClassify()
	search := make(map[string]interface{})
	search["id"] = validate.ID
	search["user_id"] = userID
	result := modelBlogClassify.SoftDelete(search)
	if result != nil {
		returnData["code"] = 1
		returnData["msg"] = result.Error()
	} else {
		returnData["code"] = 0
	}
	jsonHandle(c, returnData)
}
