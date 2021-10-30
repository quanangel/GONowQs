package mysql

import "time"

// UsersToken is users token table struct
type UsersToken struct {
	// id
	ID int64 `gorm:"column:id;type:bigint(20) auto_increment;primaryKey;comment:id"`
	// user id
	UserID int64 `gorm:"column:user_id;type:bigint(20);index;not null;comment:user id"`
	// token
	Token string `gorm:"column:token;type:varchar(100);index;not null;comment:token"`
	// BaseTimeModel
	BaseTimeModel
}

// NewUsersToken is return MemberToken struct
func NewUsersToken() UsersToken {
	return UsersToken{}
}

// Add is add token function
func (m *UsersToken) Add(userID int64, token string) bool {
	m.UserID = userID
	m.Token = token
	m.AddTime = int(time.Now().Unix())
	m.UpdateTime = m.AddTime
	db := GetDb()
	result := db.Create(m)
	return result.RowsAffected > 0
}

// GetTokenByID is get token by userID function
func (m *UsersToken) GetTokenByID(userID int64) string {
	db := GetDb()
	db.Where("user_id=?", userID).First(m)
	return m.Token
}

// Del is del token function
func (m *UsersToken) Del(userID int64, token string) bool {
	db := GetDb()
	result := db.Where("user_id=? and token=?", userID, token).Delete(m)
	return result.RowsAffected > 0
}
