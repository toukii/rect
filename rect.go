package rect

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"

	"github.com/toukii/bytes"
	"github.com/toukii/icat"
)

func leftTop(x0, y0, x1, y1 *int) {
	if *x0 > *x1 {
		*x0, *x1 = *x1, *x0
	}
	if *y0 > *y1 {
		*y0, *y1 = *y1, *y0
	}
}

func Rect(imgNRGBA *image.NRGBA, x0, y0, x1, y1 int, c color.Color) {
	leftTop(&x0, &y0, &x1, &y1)

	X := []int{x0, x1}
	xx := (X[1] - X[0]) / 3
	Y := []int{y0, y1}
	yy := (Y[1] - Y[0]) / 3

	for _, x := range X {
		for y := Y[0]; y < Y[1]; y++ {
			if y-Y[0] > yy && y-Y[0] < yy<<1 {
				continue
				// c = &color.RGBA{}
			}
			imgNRGBA.Set(x, y, c)
		}
	}

	for _, y := range Y {
		for x := X[0]; x < X[1]; x++ {
			if x-X[0] > xx && x-X[0] < xx<<1 {
				continue
				// c = &color.RGBA{}
			}
			imgNRGBA.Set(x, y, c)
		}
	}

	wr := bytes.NewWriter(make([]byte, 0, 10240))

	err := jpeg.Encode(wr, imgNRGBA, &jpeg.Options{90})
	if err != nil {
		panic(err)
	}

	img, err := jpeg.Decode(bytes.NewReader(wr.Bytes()))
	if err != nil {
		panic(err)
	}

	err = icat.ICat(img, os.Stdout)
	if err != nil {
		fmt.Println(err)
	}
}
