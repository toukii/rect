package rect

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"testing"

	"github.com/toukii/bytes"
	"github.com/toukii/icat"
)

func TestRect(t *testing.T) {
	rect := image.Rect(0, 0, 600, 400)
	bg := image.NewNRGBA(rect)

	Rect(bg, 100, 100, 300, 300, color.RGBA{
		A: 255,
		R: 255,
	})

	wr := bytes.NewWriter(make([]byte, 0, 10240))

	err := jpeg.Encode(wr, bg, &jpeg.Options{90})
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
