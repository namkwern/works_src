package main

import (
	"github.com/intelfike/rcwindow"
	"image/color"
	"math"
	"math/cmplx"
	"time"
)

var col = []color.Color{
	color.RGBA{0xff, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0xff, 0xff},
	color.RGBA{0xff, 0xff, 0x00, 0xff},
	nil,
}

func main() {
	rc := rcwindow.NewWindow(1, 1, 2000)
	rc.SafeConfig(func() { rc.DotSize = 2 })
	rc.DrawTick(1 << 24)
	v := 1 + 0i
	c := cmplx.Pow(0+1i, 1.0/1000.0)
	count := 0
	rc.FillXc(
		func(x float64) (float64, color.Color) {
			count++
			count %= 2
			v *= c
			if x > 0 {
				count += 2
			}
			switch count {
			case 0:
				return math.Sqrt(imag(v)),
					col[count]
			case 2:
				return math.Sqrt(imag(v)),
					col[count]
			case 1:
				return -math.Sqrt(imag(v)),
					col[count]
			case 3:
				return -math.Sqrt(imag(v)),
					col[count]
			}
			return 0, nil
		},
		func() {
			time.Sleep(1)
		},
	)
	rc.Wait()
}
