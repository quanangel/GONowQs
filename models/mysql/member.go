package mysql

import (
	"crypto/sha512"
	"encoding/hex"
	"time"
)

// Member is member table struct
type Member struct {
	// user id
	UserID int64 `gorm:"column:user_id;type:bigint(20);primaryKey;autoIncrement;comment:user id"`
	// user name
	UserName string `gorm:"column:user_name;type:varchar(50);index;not null;comment:user name"`
	// nick name
	NickName string `gorm:"column:nick_name;type:varchar(50);not null;comment:nick name"`
	// password
	Password string `gorm:"column:password:type:varchar(1024);not null;comment:password"`
	// status: 1normal、2disable
	Status int8 `gorm:"column:status;type:tinyint(1);default:1;comment:status:1normal、2disable"`
	// last login ip
	LastIP string `gorm:"column:last_ip;type:varchar(20);null;comment:last login ip"`
	// last login time
	LastTime int `gorm:"column:last_time;type:int(10);null;comment:last login time"`
	// register time
	RegisterTime int `gorm:"column:register_time;type:int(10);not null;comment:register time"`
}

// NewMember is return Member struct
func NewMember() Member {
	return Member{}
}

// Add is add member function
func (m *Member) Add(username string, nickname string, password string) int64 {
	m.UserName = username
	m.NickName = nickname
	m.Password = m.Sha512(password)
	m.RegisterTime = int(time.Now().Unix())
	db := GetDb()
	result := db.Create(m)
	if result.RowsAffected > 0 {
		return m.UserID
	}
	return 0
}

// GetAll is get all membmer message
func (m *Member) GetAll() (users *[]Member) {
	db := GetDb()
	db.Find(users)
	return users
}

// GetList is get all membmer message
func (m *Member) GetList(search map[string]string, page int, limit int) (users *[]Member) {
	db := GetDb()
	for key := range search {
		db.Or(key+" LIKE ?", "%"+search[key]+"%")
	}
	db.Offset((page - 1) * limit).Limit(page * limit).Find(users)
	return users
}

// GetByID is get member message by id
func (m *Member) GetByID(userID int64) *Member {
	db := GetDb()
	db.Where("user_id=?", userID).First(m)
	return m
}

// Sha512 is sha-512 encode
func (m *Member) Sha512(content string) string {
	newSha := sha512.New()
	newSha.Write([]byte(content))
	return hex.EncodeToString(newSha.Sum(nil))
}
