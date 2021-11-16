package mysql

import (
	"nowqs/frame/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// Captcha struct
type Captcha struct {
	// ID
	ID int64 `gorm:"column:id;type:bigint(20) auto_increment;primaryKey;comment:id"`
	// CaptchaKey
	CaptchaKey string `gorm:"column:captcha_key;type:varchar(100);not null;comment:captcha key"`
	// CaptchaValue
	CaptchaValue string `gorm:"column:captcha_value;type:varchar(20);not null;comment:captcha value"`
	// ExpireTime
	ExpireTime int `gorm:"column:expire_time;type:int(10);not null;comment:expire time"`
	// BaseTimeModel
	BaseTimeModel
}

// NewCaptcha is return catpcha struct
func NewCaptcha() Captcha {
	return Captcha{}
}

// Add is add message
func (m *Captcha) Add() gin.H {
	res := gin.H{}
	config := utils.NewDefaultOptions()
	data, err := config.New()
	if err != nil {
		return res
	}
	m.CaptchaKey = data.CodeKey
	m.CaptchaValue = data.Code
	m.AddTime = int(time.Now().Unix())
	m.UpdateTime = m.AddTime
	m.ExpireTime = m.AddTime + data.ExpireTime
	db := GetDb()
	result := db.Create(m)
	if result.RowsAffected > 0 {
		res = gin.H{
			"key":         data.CodeKey,
			"expire_time": m.ExpireTime,
			"image":       data.Image,
		}
	}
	return res
}

func (m *Captcha) Check(key string, value string) bool {
	db := GetDb()
	nowTime := int(time.Now().Unix())
	db.Where("captcha_key = ? and expire_time >= ?", key, nowTime).Find(&m)
	if db.RowsAffected > 0 {
		db.Delete(m)
	}
	return m.CaptchaValue == value
}
