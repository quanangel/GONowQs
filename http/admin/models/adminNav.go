package models

import (
	"nowqs/frame/models/mysql"
)

// AdminNav is struct
type AdminNav struct {
	mysql.AdminNav
}

// NewAdminNav is return mysql adminNav struct
func NewAdminNav() AdminNav {
	return AdminNav{}
}

func (m *AdminNav) GetOne(userID int64, search map[string]interface{}) *AdminNav {

	return m
}
