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

// blogClassifyValidate is get validate struct
type blogClassifyValidate struct {
	// Type: my/list/only
	Type string `form:"type" json:"type" xml:"type" binding:"required,oneof=my list only" default:"list"`
	// Search: type is only the search is id, type is list the search is id/name/url
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

// @Summary BlogClassify
// @Tags BlogClassify
// @Description BlogClassify
// @Produce json
// @Param Auth-Token header string false "Auth-Token"
// @Param object query blogClassifyValidate false "get message"
// @Success 200-1 {object} _returnBlogClassifyGetList
// @Success 200-2 {object} _returnBlogClassifyGetOnly
// @Failure 400 {object} _returnError
// @Router /blog/v1/blog_classify [get]
// Get is get user message
func (a *BlogClassify) Get(c *gin.Context) {
	returnData := gin.H{
		"code": -1,
	}
	authToken := c.GetHeader("Auth-Token")
	userID := userAuth(authToken)

	var validate blogClassifyValidate
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
		result := modelBlogClassify.GetOne(id)
		returnData["code"] = 0
		returnData["data"] = result
	default:
		jsonHandle(c, returnData)
		return
	}
	jsonHandle(c, returnData)
}
