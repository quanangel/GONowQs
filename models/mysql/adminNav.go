package mysql

import (
	"fmt"
	"time"
)

// AdminNav is table admin_nav table
type AdminNav struct {
	// id
	ID int `gorm:"column:id;type:int(10) auto_increment;primaryKey;comment:id"`
	// pid
	PID int `gorm:"column:pid;type:int(10);default:0;index:pid;comment:pid"`
	// name
	Name string `gorm:"column:name;type:varchar(10);not null;comment:name"`
	// url
	Url string `gorm:"column:name;type:varchar(100);comment:url"`
	// status: 1normal、2disable
	Status int8 `gorm:"column:status;type:tinyint(1);default:1;comment:status:1normal、2disable"`
	// add time
	AddTime int `gorm:"column:add_time;type:int(10);not null;comment:add time"`
	// udpate time
	UpdateTime int `gorm:"column:update_time;type:int(10);not null;comment:update time"`
}

// NewAdminNav is return AdminNav struct
func NewAdminNav() AdminNav {
	return AdminNav{}
}

// Add is add nav message
func (m *AdminNav) Add(name string, pid int, url string, status int8) int {
	m.Name = name
	m.PID = pid
	m.Url = url
	m.Status = status
	m.AddTime = int(time.Now().Unix())
	m.UpdateTime = int(time.Now().Unix())
	db := GetDb()
	result := db.Create(m)
	if result.RowsAffected > 0 {
		return m.ID
	}
	return 0
}

// Edit is edit message by search
func (m *AdminNav) Edit(search map[string]interface{}, data map[string]interface{}) bool {
	db := GetDb()
	result := db.Where(search).Updates(data)
	if result.RowsAffected > 0 {
		return true
	}
	return false
}

// Del is delete message by search
func (m *AdminNav) Del(search map[string]interface{}) bool {
	db := GetDb()
	result := db.Where(search).Delete(m)
	if result.RowsAffected > 0 {
		return true
	}
	return false
}

// GetOne is get one single message by search
func (m *AdminNav) GetOne(search map[string]interface{}) *AdminNav {
	db := GetDb()
	db.Where(search)
	db.First(m)
	return m
}

// GetList is get list by search
func (m *AdminNav) GetList(search map[string]interface{}, page int, limit int) (list *[]AdminNav) {
	db := GetDb()
	db.Where(search)
	if nil != search["name"] {
		db.Or("name like ?", "%"+fmt.Sprintf("%v", search["name"])+"%")
	}
	db.Offset((page - 1) * limit).Limit(page * limit).Find(list)
	return list
}
