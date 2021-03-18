package models

import (
	"nowqs/frame/models/mysql"
)

// Auth is auth model
type Auth struct{}

func NewAuth() Auth {
	return Auth{}
}

// CheckUser is check user rule by user id
func (m *Auth) CheckUser(userID int64, rule string) bool {
	authGroup := mysql.AuthGroup{}
	return authGroup.CheckUser(userID, rule)
}

// GetRule is get user rule map list
func (m *Auth) GetRule(userID int64) map[int]string {
	authGroup := mysql.AuthGroup{}
	return authGroup.GetRules(userID)
}
