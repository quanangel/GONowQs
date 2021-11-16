package controller

import "nowqs/frame/http/blog/models"

type _returnSuccess struct {
	Code int    `json:"code" example:"0"`      // error code
	Msg  string `json:"msg" example:"success"` // message
}

type _returnError struct {
	Code int    `json:"code" example:"1"`    // error code
	Msg  string `json:"msg" example:"error"` // message
}

type _returnPageLimit struct {
	Total int `json:"total" example:"100"` // total
	Page  int `json:"page" example:"1"`    // page
	Limit int `json:"limit" example:"20"`  // limit
}

type _returnCaptcha struct {
	Key        string `json:"key" example:"k1111"`
	Image      string `json:"image" example:"base64:****"`
	ExpireTime int    `json:"expire_time" example:"1234567890"`
}

type _returnLoginPut struct {
	_returnSuccess
	Data string `json:"data" example:"token"` // token
}

type _returnLoginGet struct {
	_returnSuccess
	Data struct {
		UserName string `json:"username" example:"username"`    // user name
		NickName string `json:"nickname" example:"nickname"`    // nick name
		LastIP   string `json:"last_ip" example:"127.0.0.1"`    // last ip
		LastTime string `json:"last_time" example:"1234567890"` // last time
		AddTime  string `json:"add_time" example:"1234567890"`  // register time
	}
}

type _returnBlogClassifyGetList struct {
	_returnSuccess
	Data struct {
		Total int64 `json:"total" example:"100"` // total
		Page  int   `json:"page" example:"1"`    // page
		Limit int   `json:"limit" example:"20"`  // limit
		Data  *[]models.BlogClassify
	}
}

type _returnBlogClassifyGetOnly struct {
	_returnSuccess
	Data *models.BlogClassify
}

type _returnBlogGetList struct {
	_returnSuccess
	Data struct {
		Total int64 `json:"total" example:"100"` // total
		Page  int   `json:"page" example:"1"`    // page
		Limit int   `json:"limit" example:"20"`  // limit
		Data  *[]models.Blog
	}
}

type _returnBlogGetOnly struct {
	_returnSuccess
	Data *models.Blog
}

type _returnBlogPost struct {
	_returnSuccess
	Data int `json:"data" example:"1"`
}
