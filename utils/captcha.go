package utils

import (
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

// CaptchaOptions is captcha options struct
type CaptchaOptions struct {
	// Key is encryption key
	Key string
	// CharType is preset char type
	CharType string
	// NumCode 2345678abcdefhijkmnpqrstuvwxyzABCDEFGHJKLMNPQRTUVWXY
	CharPreset string
	// Expire is the captcha expire time
	Expire int
	// FontSize
	FontSize int
	// FontDPI
	FontDPI float64
	// FontScale
	FontScale float64
	// CurveUse is use curve status
	CurveUse bool
	// CurveNumber is whether use curve
	CurveNumber int
	// NoiseUse is whether use noise
	NoiseUse bool
	// NoiseNumber is noise number
	NoiseNumber float64

	// CodeLength is text length
	CodeLength int
	// Width
	Width int
	// Height
	Height int
}

// CaptchaData is return data struct
type CaptchaData struct {
	// Code is captcha code
	Code string
	// CodeKey is captcha code key
	CodeKey string
	// ExpireTime is captcha code expire time
	ExpireTime int
}

var captchaFontByte *truetype.Font

// NewDefaultOptions is get captcha default config
func NewDefaultOptions() *CaptchaOptions {
	return &CaptchaOptions{
		Key:         "NowQs",
		CharType:    "en",
		CharPreset:  "2345678abcdefhijkmnpqrstuvwxyzABCDEFGHJKLMNPQRTUVWXY",
		Expire:      180,
		FontSize:    25,
		FontDPI:     72.0,
		FontScale:   1.0,
		CurveUse:    true,
		CurveNumber: 4,
		NoiseUse:    true,
		NoiseNumber: 1.0,
		CodeLength:  5,
		Width:       0,
		Height:      0,
	}
}

func (config *CaptchaOptions) getFontByte() (err error) {
	fontFile, _ := os.Getwd()
	rand.Seed(time.Now().UnixNano())
	fontFile += string(os.PathSeparator) + "utils" + string(os.PathSeparator) + "ttfs" + config.CharType + strconv.Itoa(rand.Intn(2)+1) + ".ttf"
	freeTypeByte, err := ioutil.ReadFile(fontFile)
	if err != nil {
		return err
	}
	captchaFontByte, err = freetype.ParseFont(freeTypeByte)

	return err
}

func (config *CaptchaOptions) New() *CaptchaData {
	data := &CaptchaData{}
	// step 1 random char
	data.Code = config.randomChar()
	// step 2 draw backgroud
	// TODO:
	// step 3 draw curve
	// step 4 draw noise
	return data
}

func (config *CaptchaOptions) randomChar() (code string) {
	charLength := len(config.CharPreset)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < config.CodeLength; i++ {
		code += string(config.CharPreset[rand.Intn(charLength)])
	}
	return code
}
