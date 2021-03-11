package controller

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"math/rand"
	"nowqs/frame/language"
	"strconv"
	"strings"
	"time"
)

// errorMsg is get error message by error code
func errorMsg(code int) string {
	return language.GetErrorMsg(code)
}

// buildToken is build token string
func buildToken(userID int) string {
	key := strconv.FormatInt(time.Now().Unix(), 10)
	rand.Seed(time.Now().UnixNano())
	key += strconv.Itoa(rand.Intn(10000))
	md5Key := md5.New()
	md5Key.Write([]byte(key))
	md5String := hex.EncodeToString(md5Key.Sum(nil))
	base64Key := base64.StdEncoding.EncodeToString([]byte(md5String + "_" + strconv.Itoa(userID)))

	return base64Key
}

func analysisToken(token string) int64 {
	tokenByte, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return 0
	}
	tokenArray := strings.Split(string(tokenByte), "_")
	if len(tokenArray) != 2 {
		return 0
	}
	userID, err := strconv.ParseInt(tokenArray[1], 10, 64)
	if err != nil {
		return 0
	}
	return userID
}
