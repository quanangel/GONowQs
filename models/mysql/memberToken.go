package mysql

import "time"

// MemberToken is member_token table struct
type MemberToken struct {
	// id
	ID int64 `gorm:"column:id;type:bigint(20) auto_increment;primaryKey;comment:id"`
	// user id
	UserID int64 `gorm:"column:user_id;type:bigint(20);index;not null;comment:user id"`
	// token
	Token string `gorm:"column:token;type:varchar(100);index;not null;comment:token"`
	// add time
	AddTime int `gorm:"column:add_time;type:int(10);not null;comment:add time"`
}

// NewMemberToken is return MemberToken struct
func NewMemberToken() MemberToken {
	return MemberToken{}
}

// Add is add token function
func (m *MemberToken) Add(userID int64, token string) bool {
	m.UserID = userID
	m.Token = token
	m.AddTime = int(time.Now().Unix())
	db := GetDb()
	result := db.Create(m)
	if result.RowsAffected > 0 {
		return true
	}
	return false
}

// GetTokenByID is get token by userID function
func (m *MemberToken) GetTokenByID(userID int64) string {
	db := GetDb()
	db.Where("user_id=?", userID).First(m)
	return m.Token
}

// Del is del token function
func (m *MemberToken) Del(userID int64, token string) bool {
	db := GetDb()
	result := db.Where("user_id=? and token=?", userID, token).Delete(m)
	if result.RowsAffected > 0 {
		return true
	}
	return false
}
