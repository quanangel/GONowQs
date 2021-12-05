package mysql

// BlogTagLink is
type BlogTagLink struct {
	// ID
	ID int64 `gorm:"column:id;type:bigint(20) auto_increment;primaryKey;comment:id"`
	// TagID
	TagID int64 `gorm:"column:tag_id;bigint(20);index:tag_id;not null;comment:tag id"`
	// BlogID
	BlogID int64 `gorm:"column:blog_id;bigint(20);index:blog_id;not null;comment:blog id"`
	// BaseTimeModel
	BaseTimeModel
}

// NewBlogTagLink is return blog tag link struct
func NewBlogTagLink() BlogTagLink {
	return BlogTagLink{}
}
