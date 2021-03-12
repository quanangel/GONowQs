package controller

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"math/rand"
	"net/http"
	"nowqs/frame/language"
	"reflect"
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

// analysisToken is analysis token function
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

func jsonHandle(data map[string]interface{}) (int, map[string]interface{}) {
	code := http.StatusOK
	switch data["code"] {
	case 0:
		code = http.StatusOK
	default:
		code = http.StatusBadRequest
	}
	errorCode := reflect.ValueOf(data["code"]).Int()
	data["msg"] = errorMsg(int(errorCode))
	return code, data
}
