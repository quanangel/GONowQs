package mysql

import "time"

// AuthGroupAccess is auth_group_access table struct
type AuthGroupAccess struct {
	ID int64 `gorm:"column:id;type:bigint(20) auto_increment;primaryKey;comment:id"`
	// user id
	UserID int64 `gorm:"column:user_id;type:bigint(20);index:user_id;not null;comment:user id"`
	// group id
	GroupID int `gorm:"column:group_id;type:int(10);index:group_id;not null;comment:group id"`
	// add time
	AddTime int `gorm:"column:add_time;type:int(10);not null;comment:add time"`
}

// NewAuthGroupAccess is return AuthGroupAccess struct function
func NewAuthGroupAccess() AuthGroupAccess {
	return AuthGroupAccess{}
}

// Add is add message function
func (m *AuthGroupAccess) Add(userID int64, groupID int) bool {
	m.UserID = userID
	m.GroupID = groupID
	m.AddTime = int(time.Now().Unix())
	db := GetDb()
	result := db.Create(m)
	if result.RowsAffected > 0 {
		return true
	}
	return false
}

// Del is delete message function
func (m *AuthGroupAccess) Del(search map[string]interface{}) int64 {
	db := GetDb()
	result := db.Where(search).Delete(m)
	if result.RowsAffected > 0 {
		return result.RowsAffected
	}
	return 0
}
