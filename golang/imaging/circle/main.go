package main

import (
	"image"
	"image/color"
	"log"

	"github.com/disintegration/imaging"
)

//代码来源 https://github.com/disintegration/imaging/issues/72
func main() {
	img, err := imaging.Open("1.jpeg")
	if err != nil {
		log.Fatal(err)
	}

	img = makeCircle(img)

	err = imaging.Save(img, "c1.png")
	if err != nil {
		log.Fatal(err)
	}
}

func makeCircle(src image.Image) image.Image {
	d := src.Bounds().Dx()
	if src.Bounds().Dy() < d {
		d = src.Bounds().Dy()
	}
	dst := imaging.CropCenter(src, d, d)
	r := d / 2
	for x := 0; x < d; x++ {
		for y := 0; y < d; y++ {
			if (x-r)*(x-r)+(y-r)*(y-r) > r*r {
				dst.SetNRGBA(x, y, color.NRGBA{0, 0, 0, 0})
			}
		}
	}
	return dst
}
