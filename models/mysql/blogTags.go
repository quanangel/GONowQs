package mysql

import (
	"time"
)

// BlogTags is
type BlogTags struct {
	// ID
	ID int64 `gorm:"column:id;type:bigint(20) auto_increment;primaryKey;comment:id"`
	//
	BlogID int64 `gorm:"column:blog_id;type:bigint(20);not null;comment:blog id"`
	// Name
	Name string `gorm:"column:name;type:varchar(100);not null;comment:tag name"`
	// BaseTimeModel
	BaseTimeModel
}

// NewBlogTags is return blog tags struct
func NewBlogTags() BlogTags {
	return BlogTags{}
}

// Add is add message function
func (m *BlogTags) Add(blogID int64, Name string) int64 {
	m.BlogID = blogID
	m.Name = Name
	m.AddTime = int(time.Now().Unix())
	m.UpdateTime = m.AddTime

	db := GetDb()
	result := db.Create(m)
	if result.RowsAffected > 0 {
		return m.ID
	}
	return 0
}

func (m *BlogTags) IsRepeat(blogID int64, Name string) bool {
	db := GetDb()
	db.Where("id = ? and name = '?'", blogID, Name).First(m)
	return m.ID == 0
}
