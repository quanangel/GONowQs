package controller

import "nowqs/frame/http/admin/models"

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

type _returnLoginPut struct {
	_returnSuccess
	Data string `json:"data" example:"token"` // token
}

type _returnLoginGet struct {
	_returnSuccess
	Data struct {
		UserName     string `json:"username" example:"username"`        // user name
		NickName     string `json:"nickname" example:"nickname"`        // nick name
		LastIP       string `json:"last_ip" example:"127.0.0.1"`        // last ip
		LastTime     string `json:"last_time" example:"1234567890"`     // last time
		RegisterTime string `json:"register_time" example:"1234567890"` // register time
	}
}

type _returnNavGetList struct {
	_returnSuccess
	// data
	Data struct {
		_returnPageLimit
		Data *[]models.AdminNav
	}
}

type _returnNavGetOnly struct {
	_returnSuccess
	// data
	Data *models.AdminNav
}
