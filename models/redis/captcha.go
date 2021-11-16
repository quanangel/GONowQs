package redis

import (
	"nowqs/frame/utils"

	"github.com/gomodule/redigo/redis"
)

// GetCaptcha is get captcha message
func GetCaptcha() (*utils.CaptchaData, error) {
	config := utils.NewDefaultOptions()
	data, err := config.New()
	if err != nil {
		return nil, err
	}
	rc := Pool.Get()
	defer rc.Close()
	_, err = rc.Do("SETEX", data.CodeKey, config.Expire, data.Code)
	if err != nil {
		return nil, err
	}
	return data, err
}

// CheckCaptcha is check captcha by key/value
func CheckCaptcha(key string, value string) bool {
	rc := Pool.Get()
	defer rc.Close()
	captcha, err := redis.String(rc.Do("GET", key))
	if err != nil {
		return false
	}
	return captcha == value
}
