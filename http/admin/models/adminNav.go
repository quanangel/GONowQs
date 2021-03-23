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
