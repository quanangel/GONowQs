package models

import (
	"nowqs/frame/models/mysql"
)

// UsersToken is struct
type UsersToken struct {
	mysql.UsersToken
}

// NewUsersToken is return mysql users_token struct
func NewUsersToken() UsersToken {
	return UsersToken{}
}
