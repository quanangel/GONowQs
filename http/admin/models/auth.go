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
func (m *Auth) CheckUser(userID int64, rule string, condition string) bool {
	authGroup := mysql.AuthGroup{}
	return authGroup.CheckUser(userID, rule, condition)
}

// GetRule is get user rule map list
func (m *Auth) GetRule(userID int64) map[int]int {
	authGroup := mysql.AuthGroup{}
	return authGroup.GetRules(userID)
}

// GetGroupList is get auth group list
func (m *Auth) GetGroupList(search map[string]interface{}, page int, limit int) []mysql.AuthGroup {
	authGroup := mysql.AuthGroup{}
	return authGroup.GetList(search, page, limit)
}

// GetGroupByID is get group message by id
func (m *Auth) GetGroupByID(id int) *mysql.AuthGroup {
	authGroup := mysql.AuthGroup{}
	return authGroup.GetByID(id)
}

// AddGroup is add group message function
func (m *Auth) AddGroup(name string, status int8, rules string) int {
	authGroup := mysql.AuthGroup{}
	return authGroup.Add(name, status, rules)
}

// EditGroup is edit group message
func (m *Auth) EditGroup(search map[string]interface{}, data map[string]interface{}) bool {
	authGroup := mysql.AuthGroup{}
	return authGroup.Edit(search, data)
}

// DelGroup is delete group mesage
func (m *Auth) DelGroup(search map[string]interface{}) bool {
	authGroup := mysql.AuthGroup{}
	return authGroup.Del(search)
}

// GetRuleList is get rule list function
func (m *Auth) GetRuleList(search map[string]interface{}, page int, limit int) *[]mysql.AuthRule {
	authRule := mysql.AuthRule{}
	return authRule.GetList(search, page, limit)
}

// AddRule is add rule function
func (m *Auth) AddRule(pid int, name string, url string, condition string, status int8) int {
	authRule := mysql.AuthRule{}
	return authRule.Add(pid, name, url, condition, status)
}

// EditRule is edit rule function
func (m *Auth) EditRule(search map[string]interface{}, data map[string]interface{}) bool {
	authRule := mysql.AuthRule{}
	return authRule.Edit(search, data)
}

// DelRule is del rule function
func (m *Auth) DelRule(search map[string]interface{}) bool {
	authRule := mysql.AuthRule{}
	return authRule.Del(search)
}

// AddGroupAccess is add group member
func (m *Auth) AddGroupAccess(userID int64, groupID int) bool {
	groupAccess := mysql.AuthGroupAccess{}
	return groupAccess.Add(userID, groupID)
}

// DelGroupAccess is delete group access
func (m *Auth) DelGroupAccess(search map[string]interface{}) int64 {
	groupAccess := mysql.AuthGroupAccess{}
	return groupAccess.Del(search)
}
