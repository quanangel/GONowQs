package models

import (
	"nowqs/frame/models/mysql"
)

// MemberToken is struct
type MemberToken struct {
	mysql.MemberToken
}

// NewMemberToken is return mysql member_token struct
func NewMemberToken() MemberToken {
	return MemberToken{}
}
