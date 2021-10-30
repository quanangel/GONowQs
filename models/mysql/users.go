package mysql

import (
	"crypto/sha512"
	"encoding/hex"
	"time"
)

// Users is users table struct
type Users struct {
	// user id
	UserID int64 `gorm:"column:user_id;type:bigint(20) auto_increment;primaryKey;comment:user id"`
	// user name
	UserName string `gorm:"column:user_name;type:varchar(50);index:user_name;not null;comment:user name"`
	// nick name
	NickName string `gorm:"column:nick_name;type:varchar(50);not null;comment:nick name"`
	// portrait
	Portrait string `gorm:"column:portrait;type:varchar(255);comment:portrait"`
	// password
	Password string `gorm:"column:password:type:varchar(1024);not null;comment:password"`
	// status: 1normal、2disable
	Status int8 `gorm:"column:status;type:tinyint(1);default:1;comment:status:1normal、2disable"`
	// last login ip
	LastIP string `gorm:"column:last_ip;type:varchar(20);null;comment:last login ip"`
	// last login time
	LastTime int `gorm:"column:last_time;type:int(10);null;comment:last login time"`
	// BaseTimeModel
	BaseTimeModel
}

// NewUsers is return users struct
func NewUsers() Users {
	return Users{}
}

// Add is add user function
func (m *Users) Add(username string, nickname string, password string) int64 {
	m.UserName = username
	m.NickName = nickname
	m.Password = m.Sha512(password)
	m.AddTime = int(time.Now().Unix())
	m.UpdateTime = m.AddTime
	db := GetDb()
	result := db.Create(m)
	if result.RowsAffected > 0 {
		return m.UserID
	}
	return 0
}

// GetAll is get all user message
func (m *Users) GetAll() (users *[]Users) {
	db := GetDb()
	db.Find(users)
	return users
}

// GetList is get all user message
func (m *Users) GetList(search map[string]string, page int, limit int) (users *[]Users) {
	db := GetDb()
	for key := range search {
		db.Or(key+" LIKE ?", "%"+search[key]+"%")
	}
	db.Offset((page - 1) * limit).Limit(page * limit).Find(users)
	return users
}

// GetByID is get user message by id
func (m *Users) GetByID(userID int64) *Users {
	db := GetDb()
	db.Where("user_id=?", userID).First(m)
	return m
}

// Sha512 is sha-512 encode
func (m *Users) Sha512(content string) string {
	newSha := sha512.New()
	newSha.Write([]byte(content))
	return hex.EncodeToString(newSha.Sum(nil))
}
