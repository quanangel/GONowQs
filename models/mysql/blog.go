package mysql

import (
	"fmt"
	"time"
)

// Blog is blog table struct
type Blog struct {
	// ID
	ID int64 `gorm:"column:id;type:bigint(20) auto_increment;primaryKey;comment:id"`
	// ClassifyID
	ClassifyID int64 `gorm:"column:id;type:bigint(20);index:classify_id;not null;comment:classify id"`
	// UserID
	UserID int64 `gorm:"column:user_id;type:bigint(20);index:user_id;not null;comment:user id"`
	// Cover
	Cover string `gorm:"column:cover;type:text;comment:cover"`
	// Title
	Title string `gorm:"column:title;type:varchar(255);not null;comment:title"`
	// Content
	Content string `gorm:"column:content;type:text;not null;comment:content"`
	// Status 0deleted 1public 2 private 3draft
	Status int8 `gorm:"column:status;type:tinyint(1);default(1);comment:status:0deleted/1public/2private/3draft"`
	// Type 1markdown 2quill
	Type int8 `gorm:"column:type;type:tinyint(1);default(1);comment:1markdown/2quill"`
	// IsPush 1yes 2no
	IsPush int8 `gorm:"column:is_push;type:tinyint(1);default(2);comment:is push:1yes/2no"`
	// AddTime
	AddTime int `gorm:"column:add_time;type:int(10);not null;comment:add time"`
	// UpdateTime
	UpdateTime int `gorm:"column:update_time:type:int(10);not null;comment:update time"`
	// ReadNum
	ReadNum int64 `gorm:"column:read_num;type:bigint(20);default(0);comment:read num"`
}

// NewBlog is return Blog struct
func NewBlog() Blog {
	return Blog{}
}

// GetList is authRuel message by list
func (m *Blog) GetList(search map[string]interface{}, page int, limit int, order string) (total int64, lists []Blog) {
	db := GetDb()

	whereOrStr := ""
	for key := range search {
		switch key {
		case "classify_id":
			db.Where("classify_id = ?", search[key])
		case "user_id":
			db.Where("user_id = ?", search[key])
		case "status":
			db.Where("status = ?", search[key])
		case "type":
			db.Where("type = ?", search[key])
		case "is_push":
			db.Where("is_push = ?", search[key])
		default:
			if whereOrStr == "" {
				whereOrStr = fmt.Sprintf("( %v LIKE '%%v%'", key, search[key])
			} else {
				whereOrStr += fmt.Sprintf(" or %v LIKE '%%v%'", key, search[key])
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
			order = "is_push asc, id desc"
		}
		db.Order(order).Offset((page - 1) * limit).Limit(page * limit).Find(&lists)
	}
	return total, lists
}

// GetByID is get message by id
func (m *Blog) GetByID(ID int) *Blog {
	db := GetDb()
	db.Where("id = ?", ID).First(m)
	return m
}

// Edit is edit message by search
func (m *Blog) Edit(search map[string]interface{}, data map[string]interface{}) bool {
	db := GetDb()
	db.Where(search).Updates(m)
	return db.RowsAffected > 0
}

// SoftDelete is soft delete message by search
func (m *Blog) SoftDelete(search map[string]interface{}) bool {
	db := GetDb()
	m.Status = 0
	m.UpdateTime = int(time.Now().Unix())
	db.Where(search).Select("status", "update_time").Updates(m)
	return db.RowsAffected > 0
}

// Add is add message function
func (m *Blog) Add(classifyID int64, userID int64, cover string, title string, content string, status int8, typeint int8, isPush int8) int64 {
	m.ClassifyID = classifyID
	m.UserID = userID
	m.Cover = cover
	m.Title = title
	m.Content = content
	m.Status = status
	m.Type = typeint
	m.IsPush = isPush
	m.AddTime = int(time.Now().Unix())
	m.UpdateTime = m.AddTime

	db := GetDb()
	result := db.Create(m)
	if result.RowsAffected > 0 {
		return m.ID
	}
	return 0
}