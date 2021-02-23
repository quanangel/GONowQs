package language

import (
	"nowqs/frame/config"
	"nowqs/frame/errorcode"
)

// GetMsg is get language message by text
func GetMsg(text string) string {
	switch config.AppConfig.Language {
	case "zh-cn":
		return zhCn[text]
	default:
		return text
	}
}

// GetErrorMsg is get error message by code
func GetErrorMsg(code int) string {
	return GetMsg(errorcode.GetMsg(code))
}
