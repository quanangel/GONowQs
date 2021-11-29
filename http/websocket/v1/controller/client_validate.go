package controller

import (
	"errors"
	"nowqs/frame/language"
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
