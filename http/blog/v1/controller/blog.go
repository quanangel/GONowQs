package controller

import (
	"nowqs/frame/http/blog/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Blog struct{}

func NewBlog() Blog {
	return Blog{}
}

// blogGetValidate is get validate struct
type blogGetValidate struct {
	// Type: my/list/only
	Type string `form:"type" json:"type" xml:"type" binding:"required,oneof=my list only" default:"list"`
	// Search: type is only the search is id, type is list the search is id/name
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

type blogPostValidate struct {
	// ClassifyID
	ClassifyID int64 `form:"classify_id" json:"classify_id" xml:"classify_id" binding:"required"`
	// Cover
	Cover string `form:"cover" json:"cover" xml:"cover" binding:"-"`
	// Title
	Title string `form:"title" json:"title" xml:"title" binding:"required"`
	// Content
	Content string `form:"content" json:"content" xml:"content" binding:"-"`
	// Status 1public/2private/3draft
	Status int8 `form:"status" json:"status" xml:"status" binding:"required" default:"1"`
	// Type 1markdown/2quill
	Type int8 `form:"type" json:"type" xml:"type" binding:"required" default:"1"`
}

type blogPutValidate struct {
	// ID
	ID int64 `form:"id" json:"id" xml:"id" binding:"required"`
	blogPostValidate
}

type blogDeleteValidate struct {
	// ID
	ID int64 `form:"id" json:"id" xml:"id" binding:"required"`
}

// @Summary Blog
// @Tags Blog
// @Description Blog
// @Produce json
// @Param Auth-Token header string false "Auth-Token"
// @Param object query blogGetValidate false "get message"
// @Success 2001 {object} _returnBlogGetList
// @Success 2002 {object} _returnBlogGetOnly
// @Failure 400 {object} _returnError
// @Router /blog/v1/blog [get]
// Get is get blog message
func (a *Blog) Get(c *gin.Context) {
	returnData := gin.H{
		"code": -1,
	}
	authToken := c.GetHeader("Auth-Token")
	userID := userAuth(authToken)

	var validate blogGetValidate
	if err := c.Bind(&validate); err != nil {
		returnData["code"] = 10000
		returnData["msg"] = err.Error()
		jsonHandle(c, returnData)
		return
	}

	modelBlog := models.NewBlog()
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
		total, result := modelBlog.GetList(search, validate.Page, validate.Limit, validate.Order)
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
		total, result := modelBlog.GetList(search, validate.Page, validate.Limit, validate.Order)
		returnData["code"] = 0
		returnData["data"] = gin.H{
			"total": total,
			"page":  validate.Page,
			"limit": validate.Limit,
			"data":  result,
		}
	case "only":
		id, _ := strconv.ParseInt(validate.Search, 10, 64)
		result := modelBlog.GetByID(id, userID)
		if result == nil {
			returnData["code"] = 6
		} else {
			returnData["code"] = 0
			returnData["data"] = result
		}
	}

	jsonHandle(c, returnData)
}

// @Summary Blog
// @Tags Blog
// @Description Blog
// @Produce json
// @Param Auth-Token header string true "Auth-Token"
// @Param object query blogPostValidate false "add message"
// @Success 200 {object} _returnBlogPost
// @Failure 400 {object} _returnError
// @Router /blog/v1/blog [post]
// Post is add blog message
func (a *Blog) Post(c *gin.Context) {
	returnData := gin.H{
		"code": -1,
	}
	userID := userAuth(c.GetHeader("Auth-Token"))
	if userID == 0 {
		returnData["code"] = 2
		jsonHandle(c, returnData)
		return
	}
	var validate blogPostValidate
	if err := c.Bind(&validate); err != nil {
		returnData["code"] = 10000
		returnData["msg"] = err.Error()
		jsonHandle(c, returnData)
		return
	}

	modelBlog := models.NewBlog()
	result := modelBlog.Add(validate.ClassifyID, userID, validate.Cover, validate.Title, validate.Content, validate.Status, validate.Type, 2)
	if result == 0 {
		returnData["code"] = 1
		jsonHandle(c, returnData)
		return
	}
	returnData["code"] = 0
	returnData["data"] = result

	jsonHandle(c, returnData)
}

// @Summary Blog
// @Tags Blog
// @Description Blog
// @Produce json
// @Param Auth-Token header string true "Auth-Token"
// @Param object query blogPutValidate false "put message"
// @Success 200 {object} _returnSuccess
// @Failure 400 {object} _returnError
// @Router /blog/v1/blog [put]
// Put is edit blog message
func (a *Blog) Put(c *gin.Context) {
	returnData := gin.H{
		"code": -1,
	}
	userID := userAuth(c.GetHeader("Auth-Token"))
	if userID == 0 {
		returnData["code"] = 2
		jsonHandle(c, returnData)
		return
	}
	var validate blogPutValidate
	if err := c.Bind(&validate); err != nil {
		returnData["code"] = 10000
		returnData["msg"] = err.Error()
		jsonHandle(c, returnData)
		return
	}

	modelBlog := models.NewBlog()
	search := make(map[string]interface{})
	search["id"] = validate.ID
	search["user_id"] = userID
	editData := make(map[string]interface{})
	editData["classify_id"] = validate.ClassifyID
	editData["title"] = validate.Title
	editData["cover"] = validate.Cover
	editData["content"] = validate.Content
	editData["status"] = validate.Status
	editData["type"] = validate.Type
	result := modelBlog.Edit(search, editData)
	if result {
		returnData["code"] = 0
	} else {
		returnData["code"] = 1
	}

	jsonHandle(c, returnData)
}

// @Summary Blog
// @Tags Blog
// @Description Blog
// @Produce json
// @Param Auth-Token header string true "Auth-Token"
// @Param object query blogDeleteValidate false "delete message"
// @Success 200 {object} _returnSuccess
// @Failure 400 {object} _returnError
// @Router /blog/v1/blog [delete]
// Delete is delete blog message
func (a *Blog) Delete(c *gin.Context) {
	returnData := gin.H{
		"code": -1,
	}
	userID := userAuth(c.GetHeader("Auth-Token"))
	if userID == 0 {
		returnData["code"] = 2
		jsonHandle(c, returnData)
		return
	}
	var validate blogDeleteValidate
	if err := c.Bind(&validate); err != nil {
		returnData["code"] = 10000
		returnData["msg"] = err.Error()
		jsonHandle(c, returnData)
		return
	}
	search := make(map[string]interface{})
	search["id"] = validate.ID
	search["user_id"] = userID
	modelBlog := models.NewBlog()
	result := modelBlog.SoftDelete(search)
	if result {
		returnData["code"] = 0
	} else {
		returnData["code"] = 1
	}
	jsonHandle(c, returnData)
}
