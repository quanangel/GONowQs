package mysql

import (
	"crypto/sha512"
	"encoding/hex"
	"time"
)

// Member is member table struct
type Member struct {
	// 用户ID
	UserID int64 `gorm:"column:user_id;type:bigint(20);primaryKey;autoIncrement"`
	// 用户名
	UserName string `gorm:"column:user_name;type:varchar(50);index;not null"`
	// 用户昵称
	NickName string `gorm:"column:nick_name;type:varchar(50);not null"`
	// 密码
	Password string `gorm:"column:password:type:varchar(1024);not null"`
	// 状态:1正常、2禁用
	Status int8 `gorm:"column:status;type:tinyint(1);default:1"`
	// 最后登录IP
	LastIP string `gorm:"column:last_ip;type:varchar(20);null"`
	// 最后登录时间
	LastTime int `gorm:"column:last_time;type:int(10);null"`
	// 注册时间
	RegisterTime int `gorm:"column:register_time;type:int(10);not null"`
}

// NewMember is return Member struct
func NewMember() Member {
	return Member{}
}

// Add is add member function
func (m *Member) Add(username string, nickname string, password string) int64 {
	m.UserName = username
	m.NickName = nickname
	m.Password = m.sha512(password)
	m.RegisterTime = int(time.Now().Unix())
	db := GetDb()
	result := db.Create(m)
	if result.RowsAffected > 0 {
		return m.UserID
	}
	return 0
}

// GetAll is get all membmer message
func (m *Member) GetAll() (users []Member) {
	db := GetDb()
	db.Find(&users)
	return users
}

// GetList is get all membmer message
func (m *Member) GetList(search map[string]string, page int, limit int) (users *Member) {
	db := GetDb()
	for key := range search {
		db.Or(key+" LIKE ?", "%"+search[key]+"%")
	}
	db.Offset((page - 1) * limit).Limit(page * limit).Find(users)
	return users
}

// Login is member login function
func (m *Member) Login(username string, password string, lastIP string) *Member {
	db := GetDb()
	result := db.Where("user_name = ? AND password = ?", username, m.sha512(password)).First(m)
	if result.RowsAffected == 0 {
		return m
	}
	db.Model(m).Where("user_id = ?", m.UserID).Updates(map[string]interface{}{"last_ip": lastIP, "last_time": int(time.Now().Unix())})

	return m
}

// GetByID is get member message by id
func (m *Member) GetByID(userID int64) *Member {
	db := GetDb()
	db.Where("user_id=?", userID).First(m)
	return m
}

// sha512 is sha-512 encode
func (m *Member) sha512(content string) string {
	newSha := sha512.New()
	newSha.Write([]byte(content))
	return hex.EncodeToString(newSha.Sum(nil))
}
