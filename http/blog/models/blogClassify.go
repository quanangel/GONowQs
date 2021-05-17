package models

import (
	"nowqs/frame/models/mysql"
)

// BlogClassify is struct
type BlogClassify struct {
	mysql.BlogClassify
}

// NewBlogClassify is return mysql BlogClassify struct
func NewBlogClassify() BlogClassify {
	return BlogClassify{}
}
