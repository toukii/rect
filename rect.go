package rect

import (
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
	if bold < 1 {
		bold = 1
	}
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

func (r *Recter) leftTop() {
	if r.x0 > r.x1 {
		r.x0, r.x1 = r.x1, r.x0
	}
	if r.y0 > r.y1 {
		r.y0, r.y1 = r.y1, r.y0
	}
}

func (r *Recter) Init() {
	r.leftTop()
	r.X = []int{r.x0, r.x1}
	r.xx = (r.X[1] - r.X[0]) / 5
	r.Y = []int{r.y0, r.y1}
	r.yy = (r.Y[1] - r.Y[0]) / 5

	for i := 1; i < r.bold; i++ {
		r.X = append(r.X, r.x0+i)
		r.X = append(r.X, r.x1+i)
		r.Y = append(r.Y, r.y0+i)
		r.Y = append(r.Y, r.y1+i)
	}
}

//  p := image.Pt(0, 20)
// draw.DrawMask(dst, dst.Bounds(), src, image.ZP, &circle{p, r}, image.ZP, draw.Over)

func DrawRectangle(imgsrc image.Image, r *Rectangle) {
	dst := image.NewNRGBA(imgsrc.Bounds())
	draw.Draw(dst, imgsrc.Bounds(), imgsrc, image.ZP, draw.Over)
	// draw.DrawMask(dst, imgsrc.Bounds(), imgsrc, image.ZP, tar, image.ZP, draw.Over)
	// draw.DrawMask(dst, imgsrc.Bounds(), imgsrc, image.ZP, &circle{image.Pt(100, 200), 100}, image.ZP, draw.Src)
	draw.Draw(dst, r.Bounds(), r, image.ZP.Add(image.Pt(r.x0, r.y0)), draw.Over)
	// draw.DrawMask(dst, imgsrc.Bounds(), imgsrc, image.ZP, r, image.ZP, draw.Src)

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

func SetNRGBARectangle(imgNRGBA *image.NRGBA, rct *Recter) *image.NRGBA {
	// x0, y0, x1, y1, c := rct.x0, rct.y0, rct.x1, rct.y1, rct.c
	c := rct.c
	X, Y, xx, yy := rct.X, rct.Y, rct.xx, rct.yy

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
