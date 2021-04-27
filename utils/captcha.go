package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
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
	// Image is captcha image
	Image string
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
	fontFile += string(os.PathSeparator) + "utils" + string(os.PathSeparator) + "ttfs" + string(os.PathSeparator) + config.CharType + string(os.PathSeparator) + strconv.Itoa(rand.Intn(2)+1) + ".ttf"
	freeTypeByte, err := ioutil.ReadFile(fontFile)
	if err != nil {
		return err
	}
	captchaFontByte, err = freetype.ParseFont(freeTypeByte)

	return err
}

func (config *CaptchaOptions) New() (*CaptchaData, error) {
	data := &CaptchaData{}
	// step 0 init config
	if config.Width == 0 {
		config.Width = int(float64(config.CodeLength)*float64(config.FontSize)*1.5 + float64(config.CodeLength)*float64(config.FontSize)/2)
	}
	if config.Height == 0 {
		config.Height = int(float64(config.FontSize) * 2.5)
	}
	// step 1 random char
	data.Code = config.randomChar()
	// step 2 draw backgroud
	// TODO: draw color
	img := image.NewNRGBA(image.Rect(0, 0, config.Width, config.Height))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.Transparent}, image.ZP, draw.Src)

	// step 3 draw code
	err := config.getFontByte()
	if err != nil {
		return nil, err
	}
	err = config.drawCode(data.Code, img)

	if err != nil {
		return nil, err
	}

	// step 4 draw curve
	for i := 0; i < config.CurveNumber; i++ {
		config.drawCurve(img)
	}
	// step 5 draw noise
	config.drawNoise(img)

	pngTemp := bytes.NewBuffer(nil)
	err = png.Encode(pngTemp, img)
	if err != nil {
		return nil, err
	}
	data.ExpireTime = int(time.Now().Unix()) + config.Expire
	data.Image = "data:image/png;base64," + base64.StdEncoding.EncodeToString(pngTemp.Bytes())
	data.CodeKey = fmt.Sprintf("%x", md5.Sum(pngTemp.Bytes()))
	return data, err
}

// randomChar is build cpatcha code
func (config *CaptchaOptions) randomChar() (code string) {
	charLength := len(config.CharPreset)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < config.CodeLength; i++ {
		code += string(config.CharPreset[rand.Intn(charLength)])
	}
	return code
}

// drawCureve is draw curve
func (config *CaptchaOptions) drawCurve(img *image.NRGBA) {
	var xStart, xEnd int
	if config.Width < 40 {
		xStart, xEnd = 1, config.Width-1
	} else {
		xStart = rand.Intn(config.Width/10) + 1
		xEnd = config.Width - rand.Intn(config.Width/10) - 1
	}
	curveHeight := float64(rand.Intn(config.Height/6) + config.Height/6)
	yStart := rand.Intn(config.Height*2/3) + config.Height/6
	angle := 1.0 + rand.Float64()
	yFlip := 1.0
	if rand.Intn(2) == 0 {
		yFlip = -1.0
	}
	for x1 := xStart; x1 <= xEnd; x1++ {
		y := math.Sin(math.Pi*angle*float64(x1)/float64(config.Width)) * curveHeight * yFlip
		img.Set(x1, int(y)+yStart, config.randomColor(255, 0))
	}
}

// drawNoise is draw noise
func (config *CaptchaOptions) drawNoise(img *image.NRGBA) {
	noiseCount := (config.Width * config.Height) / int(28.0/config.NoiseNumber)
	for i := 0; i < noiseCount; i++ {
		x := rand.Intn(config.Width)
		y := rand.Intn(config.Height)
		img.Set(x, y, config.randomColor(255, 0))
	}
}

// drawCode is draw code
func (config *CaptchaOptions) drawCode(code string, img *image.NRGBA) error {
	ftc := freetype.NewContext()
	ftc.SetDPI(config.FontDPI)
	ftc.SetClip(img.Bounds())
	ftc.SetDst(img)
	ftc.SetHinting(font.HintingFull)
	ftc.SetFont(captchaFontByte)
	fontSpacing := config.Width / len(code)
	fontOffset := rand.Intn(fontSpacing / 2)

	for index, val := range code {
		fontScale := 0.8 + rand.Float64()*0.4
		fontSize := float64(config.Height) / fontScale * config.FontScale
		ftc.SetFontSize(fontSize)
		ftc.SetSrc(image.NewUniform(config.randomColor(125, 125)))
		x := fontSpacing*index + fontOffset
		y := config.Height/6 + rand.Intn(config.Height/3) + int(fontSize/2)
		pt := freetype.Pt(x, y)
		if _, err := ftc.DrawString(string(val), pt); err != nil {
			return err
		}
	}
	return nil
}

// randomColor is random color.Color message
func (config *CaptchaOptions) randomColor(limit int, area int) color.Color {
	red := rand.Intn(limit) + area
	green := rand.Intn(limit) + area
	blue := rand.Intn(limit) + area
	return color.RGBA{R: uint8(red), G: uint8(green), B: uint8(blue), A: uint8(255)}
}
