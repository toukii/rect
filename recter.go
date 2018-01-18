package rect

import (
	"image/color"
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
