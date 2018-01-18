package rect

import (
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"testing"
)

var img image.Image

func init() {
	fd, err := os.Open("gosea.jpg")
	if err != nil {
		panic(err)
	}
	img, _ = jpeg.Decode(fd)
}

func TestDraw(t *testing.T) {
	r := NewRectangle(50, 110, 130, 190, color.RGBA{
		A: 255,
		R: 255,
		G: 1,
		B: 1,
	}, 5)

	imgNRGBA := DrawRectangle(img, r)

	ICatPngNRGBA(imgNRGBA)
}

func TestRect(t *testing.T) {
	imgNRGBA := SetRectangle(img,
		NewRecter(50, 15, 90, 45, color.RGBA{
			A: 255,
			R: 255,
		}, 1),
		NewRecter(50, 110, 130, 190, color.RGBA{
			A: 255,
			R: 255,
		}, 5))

	ICatPngNRGBA(imgNRGBA)
}

func TestRect2(t *testing.T) {
	imgNRGBA := SetRectangle(img,
		NewRecter(90, 15, 50, 45, color.RGBA{
			A: 255,
			R: 255,
		}, 1),
		NewRecter(50, 190, 130, 110, color.RGBA{
			A: 255,
			R: 255,
		}, 5))
	ICatPngNRGBA(imgNRGBA)
}
