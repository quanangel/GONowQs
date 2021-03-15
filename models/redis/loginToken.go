package redis

import "github.com/gomodule/redigo/redis"

// SetLoginToken is save the login token function
func SetLoginToken(userID int64, token string) error {
	rc := Pool.Get()
	defer rc.Close()
	_, err := rc.Do("SET", userID, token)
	if err != nil {
		return err
	}
	return nil
}

// GetLoginToken is get the token by userID
func GetLoginToken(userID int64) string {
	rc := Pool.Get()
	defer rc.Close()
	token, _ := redis.String(rc.Do("GET", userID))
	return token
}
