package rect

import (
	"image"
	"image/color"
)

var defaultColor = &color.RGBA{}
var _empty empty

type empty struct {
}

type Rectangle struct {
	*Recter
	xm map[int]empty
	ym map[int]empty
}

func NewRectangle(x0, y0, x1, y1 int, c color.Color, bold int) *Rectangle {
	if bold < 1 {
		bold = 1
	}
	r := &Rectangle{
		Recter: NewRecter(x0, y0, x1, y1, c, bold),
		xm:     make(map[int]empty, bold),
		ym:     make(map[int]empty, bold),
	}
	for _, x := range r.X {
		r.xm[x] = _empty
	}
	for _, y := range r.Y {
		r.ym[y] = _empty
	}
	return r
}

func (r *Rectangle) ColorModel() color.Model {
	return color.RGBAModel
}

func (r *Rectangle) Bounds() image.Rectangle {
	r.Init()
	return image.Rect(r.x0, r.y0, r.x1+r.bold, r.y1+r.bold)
}

func (r *Rectangle) AtA(x, y int) color.Color {
	if _, ex := r.xm[x]; ex &&
		(y-r.Y[0] <= r.yy || y-r.Y[0] >= r.yy*4+r.bold-2) {
		return r.c
	}
	if _, ex := r.ym[y]; ex &&
		(x-r.X[0] <= r.xx || x-r.X[0] >= r.xx*4+r.bold-2) {
		return r.c
	}
	return defaultColor
}

func (r *Rectangle) At(x, y int) color.Color {
	if _, ex := r.xm[x]; ex &&
		y >= r.Y[0] && y <= r.Y[1]+r.bold-2 {
		return r.c
	}
	if _, ex := r.ym[y]; ex &&
		(x-r.X[0] <= r.xx || x-r.X[0] >= r.xx*4+r.bold-2) {
		return r.c
	}
	return defaultColor
}

func (r *Rectangle) AtArrows(x, y int) color.Color {
	if y-r.Y[0] > r.yy && y-r.Y[0] < r.yy<<1 {
		return color.Alpha{0}
	}
	if x-r.X[0] > r.xx && x-r.X[0] < r.xx<<1 {
		return color.Alpha{0}
	}
	return color.Alpha{255}
}
