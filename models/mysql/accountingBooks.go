package mysql

import (
	"fmt"
	"time"
)

type AccountingBooks struct {
	// ID
	ID int64 `gorm:"column:id;type:bigint(20) auto_increment;primaryKey;comment:id"`
	// UserID
	UserID int64 `gorm:"column:user_id;type:bigint(20);default:0;index:user_id;comment:user id"`
	// Name
	Name string `gorm:"column:name;type:varchar(200);not null;comment:name"`
	// Cover
	Cover string `gorm:"column:cover;type:varchar(100);comment:cover"`
	// Status 0deleted 1normal 2disable
	Status int8 `gorm:"column:status;type:tinyint(1);default(1);comment:status:0deleted/1normal/2disable"`
	// Type 0deleted 1normal 2disable
	Type int8 `gorm:"column:type;type:tinyint(1);default(2);comment:status:1public/2private"`
	// OrderID
	OrderID int64 `gorm:"column:order_id;type:bigint(20);default(0);comment:order id"`
	// BaseTimeModel
	BaseTimeModel
}

// NewAccountingBooks is return AccountingBooks struct
func NewAccountingBooks() AccountingBooks {
	return AccountingBooks{}
}

// GetList is get message list
func (m *AccountingBooks) GetList(search map[string]interface{}, page int, limit int, order string) (total int64, lists []AccountingBooks) {
	db := GetDb()
	whereOrStr := ""
	for key, val := range search {
		switch key {
		case "user_id":
			db.Where("user_id = ?", val)
		case "id":
			db.Where("id = ?", val)
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
func (m *AccountingBooks) GetByID(id int64, userID int64) *AccountingBooks {
	db := GetDb()
	db.Where("id = ?", id)
	db.First(m)
	if m.Status != 1 && m.UserID != userID {
		return nil
	}
	return m
}

// SoftDelete is soft delete message by search
func (m *AccountingClassify) SoftDelete(search map[string]interface{}) bool {
	db := GetDb()
	m.Status = 0
	m.UpdateTime = int(time.Now().Unix())
	result := db.Where(search).Select("status", "update_time").Updates(m)
	return result.RowsAffected > 0
}

// Edit is edit message function
func (m *AccountingBooks) Edit(search map[string]interface{}, data map[string]interface{}) bool {
	db := GetDb()
	result := db.Model(m).Where(search).Updates(data)
	return result.RowsAffected > 0
}

// Add is add message function
func (m *AccountingBooks) Add(userID int64, name string, cover string, typeInt int8, status int8, orderID int64) int64 {
	m.UserID = userID
	m.Name = name
	m.Cover = cover
	m.Type = typeInt
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
