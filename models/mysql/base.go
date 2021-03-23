package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"nowqs/frame/config"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	// _db is GORM.DB example
	pool *gorm.DB
	// logger is log example
	logger *log.Logger
	// prefix is table prefix
	prefix string
)

func init() {
	prefix = config.AppConfig.Db.Prefix

}

// InitDb is initialization DB example function
func InitDb() {
	if config.AppConfig.Db.Status {
		logFile, err := os.OpenFile(getLogFileName(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Fprintln(logFile, err)
		}
		mysqlLogger := gormLogger.New(
			log.New(logFile, "\r\n", log.LstdFlags|log.Lshortfile),
			gormLogger.Config{
				SlowThreshold: 200 * time.Millisecond,
				LogLevel:      gormLogger.Info,
				Colorful:      false,
			},
		)

		pool, err = gorm.Open(mysql.New(mysql.Config{
			DSN: config.AppConfig.Db.DSN,
		}), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				// 表名前缀，`User` 的表名应该是 `t_users`
				TablePrefix: prefix,
				// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
				SingularTable: true,
			},
			// Logger: gormLogger.Default.LogMode(gormLogger.Info),
			Logger: mysqlLogger,
		})
		checkError(err, true)
		sqlDb, err := pool.DB()
		checkError(err, true)
		sqlDb.SetMaxIdleConns(config.AppConfig.Db.SetMaxIdleConns)
		sqlDb.SetMaxOpenConns(config.AppConfig.Db.SetMaxOpenConns)

		db := GetDb()
		db.AutoMigrate(&AuthRule{}, &AuthGroup{}, &AuthGroupAccess{}, &AdminNav{}, &Member{}, &MemberToken{})

	} else {
		logWrite("db status is close")
	}
}

// getLogPath is get mysql path name
func getLogPath() string {
	now := time.Now()
	logDir := config.GetLogPath() + config.PathSeparator + now.Format("200601") + config.PathSeparator + "mysql"
	os.Mkdir(logDir, 0666)
	return logDir
}

// getLogFileName is get log file name have path
func getLogFileName() string {
	now := time.Now()
	return getLogPath() + config.PathSeparator + strconv.Itoa(now.Day()) + ".log"
}

// logWrite is make log
func logWrite(content string) {
	logfile, _ := os.OpenFile(getLogFileName(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer logfile.Close()
	logger = log.New(logfile, "", log.LstdFlags|log.Lshortfile)
	logger.Println(content)
	if config.AppConfig.Debug {
		println(content)
	}
}

// checkError is handle error function
func checkError(err error, isPanic bool) {
	switch {
	case err == sql.ErrNoRows:
		break
	case err != nil:
		if config.AppConfig.Debug {
			fmt.Println(err.Error())
		}
		logWrite("db error:" + err.Error())
		if isPanic {
			panic(err.Error())
		}
	}
}

// GetDb is get db pool function
func GetDb() *gorm.DB {
	return pool
}
