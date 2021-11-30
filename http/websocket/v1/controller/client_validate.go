package controller

import (
	"errors"
	"nowqs/frame/language"
	"nowqs/frame/utils"
)

func (v *_Login) validate(param *_Login) (err error) {
	if param.Action != "login" {
		err = errors.New(language.GetMsg("illegal action"))
	}
	if param.UserType == "" {
		err = errors.New(language.GetMsg("lack user type"))
	}
	if param.Token == "" {
		err = errors.New(language.GetMsg("lack token"))
	}

	return
}

func (v *_SendMessage) validate(param *_SendMessage) (err error) {
	if param.Action != "send_message" {
		err = errors.New(language.GetMsg("illegal action"))
	}
	if param.UserType == "" {
		err = errors.New(language.GetMsg("lack user type"))
	}
	if param.UserID == "" {
		err = errors.New(language.GetMsg("lack user id"))
	}
	if _, ok := utils.InSlice([]string{"image", "text"}, param.SendType); !ok {
		err = errors.New(language.GetMsg("illegal type"))
	}
	if param.Content == "" {
		err = errors.New(language.GetMsg("content cannot be empty"))
	}
	return
}
