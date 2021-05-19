package mysql

import (
	"fmt"
	"time"
)

type BlogClassify struct {
	// ID
	ID int64 `gorm:"column:id;type:bigint(20) auto_increment;primaryKey;comment:id"`
	// PID
	PID int64 `gorm:"column:pid;type:bigint(20);default(0);index:pid;comment:pid"`
	// UserID
	UserID int64 `gorm:"column:user_id;type:bigint(20);default:0;index:user_id;comment:user id"`
	// Name
	Name string `gorm:"column:name;type:varchar(200);not null;comment:name"`
	// Type
	Type int8 `gorm:"column:type;type:tinyint(1);default(1);comment:type:1markdown/2quill"`
	// Status 0deleted 1public 2 private 3draft
	Status int8 `gorm:"column:status;type:tinyint(1);default(1);comment:status:0deleted/1public/2privarte/3draft"`
	// AddTime
	AddTime int `gorm:"column:add_time;type:int(10);not null;comment:add time"`
	// UpdateTime
	UpdateTime int `gorm:"column:udpate_time;type:int(10);not null;comment:update time"`
	// OrderID
	OrderID int64 `gorm:"column:order_id;type:bigint(20);default(0);comment:order id"`
}

// NewBlogClassify is return BlogClassify struct function
func NewBlogClassify() BlogClassify {
	return BlogClassify{}
}

// GetList is get message list
func (m *BlogClassify) GetList(search map[string]interface{}, page int, limit int, order string) (total int64, lists []BlogClassify) {
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
func (m *BlogClassify) GetByID(id int64, userID int64) *BlogClassify {
	db := GetDb()
	db.Where("id = ?", id)
	db.First(m)
	if m.Status != 1 && m.UserID != userID {
		return nil
	}
	return m
}

// SoftDelete is soft delete message by search
func (m *BlogClassify) SoftDelete(search map[string]interface{}) error {
	db := GetDb()
	tx := db.Begin()
	var listTmp []BlogClassify
	if err := tx.Where(search).Find(&listTmp).Error; err != nil {
		return err
	}
	if len(listTmp) == 0 {
		return nil
	}

	m.Status = 0
	m.UpdateTime = int(time.Now().Unix())
	if err := tx.Where(search).Select("status", "update_time").Updates(m).Error; err != nil {
		tx.Rollback()
		return err
	}

	var classifyID []int64
	for _, val := range listTmp {
		classifyID = append(classifyID, val.ID)
	}
	blogStruct := NewBlog()
	blogStruct.Status = 0
	blogStruct.UpdateTime = int(time.Now().Unix())
	if err := tx.Where("classify_id IN ?", classifyID).Select("status", "update_time").Updates(&blogStruct).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// Edit is edit message function
func (m *BlogClassify) Edit(search map[string]interface{}, data map[string]interface{}) bool {
	db := GetDb()
	result := db.Model(m).Where(search).Updates(data)
	return result.RowsAffected > 0
}

// Add is add message function
func (m *BlogClassify) Add(pid int64, userID int64, name string, typeInt int8, status int8, orderID int64) int64 {
	m.PID = pid
	m.UserID = userID
	m.Name = name
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
