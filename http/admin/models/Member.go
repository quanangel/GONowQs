package models

import (
	"nowqs/frame/models/mysql"
	"time"
)

// Member is struct
type Member struct {
	mysql.Member
}

// NewMember is return mysql member struct
func NewMember() Member {
	return Member{}
}

// Login is member login function
func (m *Member) Login(username string, password string, lastIP string) *Member {
	db := mysql.GetDb()
	result := db.Where("user_name = ? AND password = ?", username, m.Sha512(password)).First(m)
	if result.RowsAffected == 0 {
		return m
	}
	db.Model(m).Where("user_id = ?", m.UserID).Updates(map[string]interface{}{"last_ip": lastIP, "last_time": int(time.Now().Unix())})

	return m
}
