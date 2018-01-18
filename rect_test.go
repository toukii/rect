package rect

import (
	"image/color"
	"image/jpeg"
	"os"
	"testing"
)

func TestRect(t *testing.T) {
	fd, err := os.Open("gosea.jpg")
	if err != nil {
		panic(err)
	}
	bg, _ := jpeg.Decode(fd)

	// Draw(bg, 50, 50, 100, 100)
	// return
	SetRectangle(bg,
		NewRecter(50, 15, 90, 45, color.RGBA{
			A: 255,
			R: 255,
		}, 1),
		NewRecter(50, 110, 130, 190, color.RGBA{
			A: 255,
			R: 255,
		}, 5))
}
