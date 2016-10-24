package main

import (
	"golang.org/x/tour/pic"
	"image/color"
	"image"
)

type Image struct{}
func (i Image)Bounds() image.Rectangle{
	return image.Rect(0,0,300,300)
}
func (i Image)ColorModel() color.Model{
	return color.RGBAModel
}
func (i Image)At(x, y int) color.Color{
	return color.RGBA{0, 0, 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}