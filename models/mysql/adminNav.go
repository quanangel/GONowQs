package mysql

import (
	"fmt"
	"reflect"
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
	Url string `gorm:"column:url;type:varchar(100);comment:url"`
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
	for k, v := range data {
		switch k {
		case "pid":
			m.PID = int(reflect.ValueOf(v).Int())
		case "name":
			m.Name = fmt.Sprintf("%v", v)
		case "url":
			m.Url = fmt.Sprintf("%v", v)
		case "status":
			m.Status = int8(reflect.ValueOf(v).Int())
		}
	}
	m.UpdateTime = int(time.Now().Unix())
	result := db.Where(search).Updates(m)
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
func (m *AdminNav) GetList(search map[string]interface{}, page int, limit int) (total int64, list []AdminNav) {

	db := GetDb()
	if nil != search["id"] {
		db.Or("id like ?", "%"+fmt.Sprintf("%v", search["id"])+"%")
	}
	if nil != search["name"] {
		db.Or("name like ?", "%"+fmt.Sprintf("%v", search["name"])+"%")
	}
	if nil != search["url"] {
		db.Or("url like ?", "%"+fmt.Sprintf("%v", search["url"])+"%")
	}
	db.Count(&total)
	total = 1
	if 0 != total {
		db.Offset((page - 1) * limit).Limit(page * limit).Find(&list)
	}

	return total, list
}
