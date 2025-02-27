package pixarray

import (
	"github.com/ckzdev/ledctl/rpi"
)

type LEDStrip interface {
	RPi() *rpi.RPi
	MaxPerChannel() int
	GetPixel(i int) Pixel
	SetPixel(i int, p Pixel)
	Write() error
}
