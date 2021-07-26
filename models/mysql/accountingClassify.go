package mysql

import (
	"fmt"
	"time"
)

type AccountingClassify struct {
	// ID
	ID int64 `gorm:"column:id;type:bigint(20) auto_increment;primaryKey;comment:id"`
	// PID
	PID int64 `gorm:"column:pid;type:bigint(20);default(0);index:pid;comment:pid"`
	// UserID
	UserID int64 `gorm:"column:user_id;type:bigint(20);default:0;index:user_id;comment:user id"`
	// Name
	Name string `gorm:"column:name;type:varchar(200);not null;comment:name"`
	// Icon
	Icon string `gorm:"column:icon;type:varchar(20);comment:icon"`
	// Status 0deleted 1normal 2disable
	Status int8 `gorm:"column:status;type:tinyint(1);default(1);comment:status:0deleted/1normal/2disable"`
	// AddTime
	AddTime int `gorm:"column:add_time;type:int(10);not null;comment:add time"`
	// UpdateTime
	UpdateTime int `gorm:"column:udpate_time;type:int(10);not null;comment:update time"`
	// OrderID
	OrderID int64 `gorm:"column:order_id;type:bigint(20);default(0);comment:order id"`
}

// NewAccountingClassify is return AccountingClassify struct function
func NewAccountingClassify() AccountingClassify {
	return AccountingClassify{}
}

// GetList is get message list
func (m *AccountingClassify) GetList(search map[string]interface{}, page int, limit int, order string) (total int64, lists []BlogClassify) {
	db := GetDb()
	whereOrStr := ""
	for key, val := range search {
		switch key {
		case "user_id":
			db.Where("user_id = ?", val)
		case "id":
			db.Where("id = ?", val)
		case "pid":
			db.Where("pid = ?", val)
		case "type":
			db.Where("type = ?", val)
		case "status":
			db.Where("status = ?", val)
		default:
			if whereOrStr == "" {
				whereOrStr = fmt.Sprintf("( %v LIKE '%%%v%%'", key, val)
			} else {
				whereOrStr += fmt.Sprintf(" or %v LIKE '%%%v%%'", key, val)
			}
		}
	}
	if whereOrStr != "" {
		whereOrStr += ")"
		db.Where(whereOrStr)
	}
	db.Count(&total)
	if total != 0 {
		if order == "" {
			order = "order_id asc, id desc"
		}
		db.Order(order).Offset((page - 1) * limit).Limit(page * limit).Find(&lists)
	}
	return total, lists
}

// GetByID is one single message where search
func (m *AccountingClassify) GetByID(id int64, userID int64) *AccountingClassify {
	db := GetDb()
	db.Where("id = ?", id)
	db.First(m)
	if m.Status != 1 && m.UserID != userID {
		return nil
	}
	return m
}

// Edit is edit message function
func (m *AccountingClassify) Edit(search map[string]interface{}, data map[string]interface{}) bool {
	db := GetDb()
	result := db.Model(m).Where(search).Updates(data)
	return result.RowsAffected > 0
}

// Add is add message function
func (m *AccountingClassify) Add(pid int64, userID int64, name string, icon string, status int8, orderID int64) int64 {
	m.PID = pid
	m.UserID = userID
	m.Name = name
	m.Icon = icon
	m.Status = status
	m.OrderID = orderID
	m.AddTime = int(time.Now().Unix())
	m.UpdateTime = m.AddTime
	db := GetDb()
	result := db.Create(m)
	if result.RowsAffected > 0 {
		return m.ID
	}
	return 0
}
