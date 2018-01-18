package rect

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"github.com/toukii/icat"
)

type Recter struct {
	x0, y0, x1, y1 int
	xx, yy         int
	bold           int
	X, Y           []int
	c              color.Color
}

func NewRecter(x0, y0, x1, y1 int, c color.Color, bold int) *Recter {
	r := &Recter{
		x0:   x0,
		y0:   y0,
		x1:   x1,
		y1:   y1,
		c:    c,
		bold: bold,
	}
	r.Init()
	return r
}

func (r *Recter) Init() {
	r.leftTop()
	r.X = []int{r.x0, r.x1}
	r.xx = (r.X[1] - r.X[0]) / 3
	r.Y = []int{r.y0, r.y1}
	r.yy = (r.Y[1] - r.Y[0]) / 3
}

func (r *Recter) ColorModel() color.Model {
	return color.AlphaModel
}

func (r *Recter) Bounds() image.Rectangle {
	r.Init()
	fmt.Println(r.X, r.Y)
	fmt.Println(r.x0, r.x1, r.y0, r.y1)
	return image.Rect(r.x0, r.y0, r.x1, r.y1)
}

func (r *Recter) At(x, y int) color.Color {
	if x == r.x1 {
		fmt.Println(x)
	}
	if (x == r.x0 || x == r.x1) &&
		y > r.y0 && y < r.y1 {
		// fmt.Print(x, "-", y, " ")
		return color.Alpha{0}
	}
	if (y == r.y0 || y == r.y1) &&
		x > r.x0 && x < r.x1 {
		return color.Alpha{0}
	}
	return color.Alpha{255}
}

func (r *Recter) AtUp(x, y int) color.Color {
	if (x == r.X[0] || x == r.X[1]) &&
		y-r.Y[0] > r.yy && y-r.Y[0] < r.yy<<1 {
		fmt.Print(x, "-", y, " ")
		return color.Alpha{0}
	}
	if (y == r.Y[0] || y == r.Y[1]) &&
		x-r.X[0] > r.xx && x-r.X[0] < r.xx<<1 {
		return color.Alpha{0}
	}
	return color.Alpha{255}
}

func (r *Recter) AtArrows(x, y int) color.Color {
	if y-r.Y[0] > r.yy && y-r.Y[0] < r.yy<<1 {
		return color.Alpha{0}
	}
	if x-r.X[0] > r.xx && x-r.X[0] < r.xx<<1 {
		return color.Alpha{0}
	}
	return color.Alpha{255}
}

func (r *Recter) leftTop() {
	if r.x0 > r.x1 {
		r.x0, r.x1 = r.x1, r.x0
	}
	if r.y0 > r.y1 {
		r.y0, r.y1 = r.y1, r.y0
	}
}

//  p := image.Pt(0, 20)
// draw.DrawMask(dst, dst.Bounds(), src, image.ZP, &circle{p, r}, image.ZP, draw.Over)

func Draw(imgsrc image.Image, x0, y0, x1, y1 int) {
	dst := image.NewNRGBA(imgsrc.Bounds())
	// draw.Draw(dst, imgsrc.Bounds(), imgsrc, image.ZP, draw.Src)

	// tar := Rect(imgsrc, x0, y0, x1, y1, color.RGBA{
	// 	A: 255,
	// 	R: 255,
	// })
	// icat.ICat(tar, os.Stdout)
	r := &Recter{
		x0: x0,
		y0: y0,
		x1: x1,
		y1: y1,
	}

	draw.Draw(dst, imgsrc.Bounds(), imgsrc, image.ZP, draw.Src)
	// draw.DrawMask(dst, imgsrc.Bounds(), imgsrc, image.ZP, tar, image.ZP, draw.Over)
	// draw.DrawMask(dst, imgsrc.Bounds(), imgsrc, image.ZP, &circle{image.Pt(100, 200), 100}, image.ZP, draw.Src)
	draw.DrawMask(dst, imgsrc.Bounds(), imgsrc, image.ZP, r, image.Point{0, 0}, draw.Src)

	ICatPngNRGBA(dst)
}

func ICatPngNRGBA(imgNRGBA *image.NRGBA) image.Image {
	w := icat.NewEncodeWr(os.Stdout, nil)
	err := png.Encode(w, imgNRGBA)
	if err != nil {
		panic(err)
	}
	w.FlushStdout()

	return nil
}

func SetRectangle(srcimg image.Image, rcts ...*Recter) image.Image {
	r := srcimg.Bounds()
	imgNRGBA := image.NewNRGBA(r)
	draw.Draw(imgNRGBA, r, srcimg, image.ZP, draw.Src)

	for _, rct := range rcts {
		SetNRGBARectangle(imgNRGBA, rct)
	}
	ICatPngNRGBA(imgNRGBA)
	return nil
}

func leftTop(x0, y0, x1, y1 *int) {
	if *x0 > *x1 {
		*x0, *x1 = *x1, *x0
	}
	if *y0 > *y1 {
		*y0, *y1 = *y1, *y0
	}
}

func SetNRGBARectangle(imgNRGBA *image.NRGBA, rct *Recter) *image.NRGBA {
	// leftTop(&x0, &y0, &x1, &y1)
	x0, y0, x1, y1, c := rct.x0, rct.y0, rct.x1, rct.y1, rct.c

	X := []int{x0, x1}
	xx := (X[1] - X[0]) / 5
	Y := []int{y0, y1}
	yy := (Y[1] - Y[0]) / 5

	for i := 1; i < rct.bold; i++ {
		X = append(X, x0+i)
		X = append(X, x1+i)
		Y = append(Y, y0+i)
		Y = append(Y, y1+i)
	}

	for _, x := range X {
		for y := Y[0]; y <= Y[len(Y)-1]; y++ {
			if y-Y[0] > yy && y-Y[0] < yy*4 {
				continue
			}
			imgNRGBA.Set(x, y, c)
		}
	}

	for _, y := range Y {
		for x := X[0]; x <= X[len(X)-1]; x++ {
			if x-X[0] > xx && x-X[0] < xx*4 {
				continue
			}
			imgNRGBA.Set(x, y, c)
		}
	}
	return imgNRGBA
}
