package mysql

// BlogTags is
type BlogTags struct {
	// ID
	ID int64 `gorm:"column:id;type:bigint(20) auto_increment;primaryKey;comment:id"`
	// Name
	Name string `gorm:"column:name;type:varchar(100);not null;comment:tag name"`
	// BaseTimeModel
	BaseTimeModel
}

// NewBlogTags is return blog tags struct
func NewBlogTags() BlogTags {
	return BlogTags{}
}
