package models

import (
	"nowqs/frame/models/mysql"
)

// Blog is struct
type Blog struct {
	mysql.Blog
}

// NewBlog is return mysql Blog struct
func NewBlog() Blog {
	return Blog{}
}
