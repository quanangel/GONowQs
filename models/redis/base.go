package redis

import (
	"strconv"
	"time"

	"nowqs/frame/config"

	"github.com/gomodule/redigo/redis"
)

var (
	// Pool is redis pool
	Pool redis.Pool
)

// InitPool is pool is init function
func InitPool() {
	if config.AppConfig.Redis.Status {
		Pool = redis.Pool{
			MaxIdle:     config.AppConfig.Redis.MaxIdle,
			MaxActive:   config.AppConfig.Redis.MaxActice,
			IdleTimeout: (100 * time.Second),
			Wait:        config.AppConfig.Redis.Wait,
			Dial: func() (redis.Conn, error) {
				con, err := redis.Dial(config.AppConfig.Redis.Protocol, config.AppConfig.Redis.Host+":"+strconv.Itoa(config.AppConfig.Redis.Port))
				if err != nil {
					return nil, err
				}
				if _, err := con.Do("AUTH", config.AppConfig.Redis.Password); err != nil {
					con.Close()
					return nil, err
				}
				if _, err := con.Do("SELECT", config.AppConfig.Redis.Db); err != nil {
					con.Close()
					return nil, err
				}
				return con, nil
			},
		}
	}
}
