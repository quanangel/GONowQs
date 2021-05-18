package controller

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"math/rand"
	"net/http"
	"nowqs/frame/config"
	models "nowqs/frame/http/blog/models"
	"nowqs/frame/language"
	redis "nowqs/frame/models/redis"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// @title NowQS Farme Blog Api
// @version 0.0.1
// @description This is Api Document
// @contact.name Qs
// @contact.email quanangel@outlook.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// errorMsg is get error message by error code
func errorMsg(code int) string {
	return language.GetErrorMsg(code)
}

// buildToken is build token string
func buildToken(userID int64) string {
	key := strconv.FormatInt(time.Now().Unix(), 10)
	rand.Seed(time.Now().UnixNano())
	key += strconv.Itoa(rand.Intn(10000))
	md5Key := md5.New()
	md5Key.Write([]byte(key))
	md5String := hex.EncodeToString(md5Key.Sum(nil))
	base64Key := base64.StdEncoding.EncodeToString([]byte(md5String + "_" + strconv.FormatInt(userID, 10)))

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

// userAuth is user authenticity function
func userAuth(token string) int64 {
	userID := analysisToken(token)
	if userID == 0 {
		return 0
	}
	serverToken := tokenGet(userID)

	if serverToken != token {
		return 0
	}

	return userID
}

// tokenSave is save token function
func tokenSave(userID int64, token string) error {
	var err error = nil
	if !config.AppConfig.Redis.Status {
		modelMemberToken := models.NewUsersToken()
		result := modelMemberToken.Add(userID, token)
		if !result {
			err = errors.New(errorMsg(5))
		}
	} else {
		err = redis.SetLoginToken(userID, token)
	}

	return err
}

// tokenGet is get token by function
func tokenGet(userID int64) string {
	var token string = ""
	if !config.AppConfig.Redis.Status {
		modelMemberToken := models.NewUsersToken()
		token = modelMemberToken.GetTokenByID(userID)
	} else {
		token = redis.GetLoginToken(userID)
	}
	return token
}

// jsonHandle is handle returnData message
func jsonHandle(c *gin.Context, data map[string]interface{}) {
	code := http.StatusOK
	switch data["code"] {
	case 0:
		code = http.StatusOK
	default:
		code = http.StatusBadRequest
	}
	if data["msg"] == "" || data["msg"] == nil {
		errorCode := reflect.ValueOf(data["code"]).Int()
		data["msg"] = errorMsg(int(errorCode))
	}
	c.JSON(code, data)
}
