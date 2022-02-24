package system

import (
	"bytes"
	"sync/atomic"

	"git.kicoe.com/blog/write/modules/utils"
	"github.com/fogleman/gg"
)

type Captcha struct {
	index int
	gg    *gg.Context
}

func NewCaptcha() *Captcha {
	return &Captcha{
		index: 0,
		gg:    gg.NewContext(100, 20),
	}
}

func (c *Captcha) Dot() {
	x := float64((c.index * 20) + 5)
	y := float64(utils.Rand(5) + 10)
	c.gg.DrawCircle(x, y, float64(utils.Rand(3)+2))
	c.gg.SetRGB255(utils.Rand(255), utils.Rand(255), utils.Rand(255))
	c.gg.Fill()
}

func (c *Captcha) Dash() {
	x := (c.index * 20) + 5
	y := utils.Rand(5) + 1
	c.gg.DrawLine(float64(utils.RandR(x, x+20)), float64(utils.RandR(y, y+10)), float64(utils.RandR(x+10, x+20)), float64(utils.RandR(y, y+20)))
	c.gg.SetRGB255(utils.Rand(255), utils.Rand(255), utils.Rand(255))
	c.gg.Stroke()
	c.gg.Fill()
}

func (c *Captcha) Draw() {
	c.gg.SavePNG("test.png")
}

var CaptchaCode atomic.Value

func (c *Captcha) Generate() *bytes.Buffer {
	var code []byte
	for i := 0; i < 4; i++ {
		r := utils.Rand(2)
		if r == 0 {
			c.Dot()
			code = append(code, '.')
		} else {
			c.Dash()
			code = append(code, '-')
		}
		c.index++
	}
	CaptchaCode.Store(code)
	buffer := bytes.NewBuffer(nil)
	c.gg.EncodePNG(buffer)
	return buffer
}
