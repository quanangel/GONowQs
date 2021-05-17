package models

import (
	"nowqs/frame/models/mysql"
	"time"
)

// Users is struct
type Users struct {
	mysql.Users
}

// NewUsers is return mysql users struct
func NewUsers() Users {
	return Users{}
}

// Login is user login function
func (m *Users) Login(username string, password string, lastIP string) *Users {
	db := mysql.GetDb()
	result := db.Where("user_name = ? AND password = ?", username, m.Sha512(password)).First(m)
	if result.RowsAffected == 0 {
		return m
	}
	db.Model(m).Where("user_id = ?", m.UserID).Updates(map[string]interface{}{"last_ip": lastIP, "last_time": int(time.Now().Unix())})

	return m
}
