package config

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"
)

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
	// Version is system version
	Version string = "0.0.1"
	// AppConfig is the system config
	AppConfig *Config
	// PathSeparator is / or \
	PathSeparator string
)

var (
	rootPath   string
	configPath string
	logPath    string
	uploadPath string
	assetsPath string
)

func init() {
	PathSeparator = string(os.PathSeparator)
	rootPath, _ = os.Getwd()
	configPath = rootPath + PathSeparator + "config"
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

	configFile := configPath + PathSeparator + "config.json"
	config := new(Config)
	_, errFile := os.Stat(configFile)
	if errFile == nil {
		configFile, errOpen := os.Open(configFile)
		defer configFile.Close()
		whiteLog(errOpen)
		decode := json.NewDecoder(configFile)
		errDecode := decode.Decode(&config)
		whiteLog(errDecode)
		if config.Db.DSN == "" {
			config.Db.DSN = config.Db.User + ":" + config.Db.Password + "@tcp(" + config.Db.Host + ":" + strconv.Itoa(config.Db.Port) + ")/" + config.Db.Db + "?charset=" + config.Db.Charset + "&parseTime=true&loc=Local"
		}
	} else {
		config = newConfigMsg()
		if config.Db.DSN == "" {
			config.Db.DSN = config.Db.User + ":" + config.Db.Password + "@tcp(" + config.Db.Host + ":" + strconv.Itoa(config.Db.Port) + ")/" + config.Db.Db + "?charset=" + config.Db.Charset + "&parseTime=true&loc=Local"
		}
		encode, errConfig := json.Marshal(config)
		whiteLog(errConfig)
		createConfigFile, _ := os.OpenFile(configFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777)
		defer createConfigFile.Close()
		_, errWhite := createConfigFile.Write(encode)
		whiteLog(errWhite)
	}
	return config
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

func whiteLog(err error) {
	now := time.Now()
	logDir := logPath + PathSeparator + now.Format("200601")
	os.Mkdir(logDir, 0777)
	logFileName := logDir + PathSeparator + strconv.Itoa(now.Day()) + "_config.log"
	logFile, _ := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	defer logFile.Close()
	logger := log.New(logFile, "", log.LstdFlags|log.Lshortfile)

	if err != nil {
		logger.Panicln(now.Format("2006-01-02 15:04:05") + ": " + err.Error())
	}
}
