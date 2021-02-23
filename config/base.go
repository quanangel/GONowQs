package config

import "os"

// Config is the system config struct
type Config struct {
	AppName  string `json:"app_name" bson:"app_name"`
	Language string `json:"language" bson:"language"`
	Debug    bool   `json:"debug" bson:"debug"`
	Host     string `json:"host" bson:"host"`
	Port     int    `json:"port" bson:"port"`
	Db       *dbConfig
	Redis    *redisConfig
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
	AppConfig = getConfig()
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

// getConfig is get Config message
func getConfig() *Config {
	// TODO:
	panic("TODO")
}

// newConfigMsg is creater config struct
func newConfigMsg() *Config {
	return &Config{
		AppName:  "NowQs",
		Debug:    true,
		Host:     "localhost",
		Port:     8080,
		Language: "zh-cn",
		Db: &dbConfig{
			Status:          true,
			TableCheck:      false,
			Type:            "mysql",
			Host:            "localhost",
			Port:            3306,
			Db:              "nowqs",
			User:            "root",
			Password:        "root",
			Charset:         "utf8mb4,utf8",
			DSN:             "",
			Prefix:          "now_",
			SetMaxIdleConns: 10,
			SetMaxOpenConns: 100,
		},
		Redis: &redisConfig{
			Status:      false,
			Protocol:    "tcp",
			Host:        "localhost",
			Port:        6379,
			Password:    "root",
			Db:          0,
			MaxIdle:     10,
			MaxActice:   10,
			IdleTimeOut: 240,
			Wait:        true,
		}}
}
