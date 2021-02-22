package config

import "os"

// Config is the system config struct
type Config struct {
	AppName  string `json:"app_name" bson:"app_name"`
	Language string `json:"language" bson:"language"`
}

var (
	// AppConfig is the system config
	AppConfig *Config
	// PathSeparator is / or \
	PathSeparator string
)

var (
	rootPath   string
	logPath    string
	uploadPath string
	assetsPath string
)

func init() {
	PathSeparator = string(os.PathSeparator)
	rootPath, _ = os.Getwd()
	logPath = rootPath + PathSeparator + "log"
	uploadPath = rootPath + PathSeparator + "upload"
	assetsPath = rootPath + PathSeparator + "assets"
}

// GetRootPath is get root path
func GetRootPath() string {
	return rootPath
}

// GetLogPath is get log path
func GetLogPath() string {
	return logPath
}

// GetUploadPath is get upload path
func GetUploadPath() string {
	return uploadPath
}
