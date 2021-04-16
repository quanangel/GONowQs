package mysql

import (
	"fmt"
	"time"
)

type AuthRule struct {
	// id
	ID int `gorm:"column:id;type:int(10) auto_increment;primaryKey;comment:id"`
	// Pid
	PID int `gorm:"column:pid;type:int(10);index:pid;default:0;comment:up one level id"`
	// name
	Name string `gorm:"column:name;type:varchar(20);not null;comment:name"`
	// type
	Type int8 `gorm:"column:type;type:tinyint(1);default:2;comment:1url/2method"`
	// url
	Url string `gorm:"column:url;type:varchar(200);comment:url"`
	// Condition
	Condition string `gorm:"column:condition;type:varchar(100);comment:condition"`
	// status: 1normal、2disable
	Status int8 `gorm:"column:status;type:tinyint(1);default:1;comment:1normal、2disable"`
	// add time
	AddTime int `gorm:"column:add_time;type:int(10);not null;add time"`
}

// NewAuthRule is return AuthRule struct function
func NewAuthRule() AuthRule {
	return AuthRule{}
}

// Add is add auth rule function
func (m *AuthRule) Add(pid int, name string, url string, condition string, status int8) int {
	m.PID = pid
	m.Name = name
	m.Url = url
	m.Condition = condition
	m.Status = status
	m.AddTime = int(time.Now().Unix())
	db := GetDb()
	result := db.Create(m)
	if result.RowsAffected > 0 {
		return m.ID
	}
	return 0
}

// Edit is AuthRule update by search
func (m *AuthRule) Edit(search map[string]interface{}, data map[string]interface{}) bool {
	db := GetDb()
	result := db.Where(search).Updates(data)
	return result.RowsAffected > 0
}

// Del is AuthRule delete message function
func (m *AuthRule) Del(search map[string]interface{}) bool {
	db := GetDb()
	result := db.Where(search).Delete(m)
	return result.RowsAffected > 0
}

// GetList is authRuel message by list
func (m *AuthRule) GetList(search map[string]interface{}, page int, limit int) (lists *[]AuthRule) {
	db := GetDb()
	for key := range search {
		db.Or(key+" LIKE ?", "%"+fmt.Sprintf("%v", search[key])+"%")
	}
	db.Offset((page - 1) * limit).Limit(page * limit).Find(lists)
	return lists
}

// GetByID is get message by id
func (m *AuthRule) GetByID(ID int) *AuthRule {
	db := GetDb()
	db.Where("id = ?", ID).First(m)
	return m
}
