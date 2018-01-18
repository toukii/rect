package rect

import (
	"image"
	"image/draw"
	"image/png"
	"os"

	"github.com/toukii/icat"
)

func DrawRectangle(imgsrc image.Image, r *Rectangle) *image.NRGBA {
	dst := image.NewNRGBA(imgsrc.Bounds())
	draw.Draw(dst, imgsrc.Bounds(), imgsrc, image.ZP, draw.Over)
	draw.Draw(dst, r.Bounds(), r, image.ZP.Add(image.Pt(r.x0, r.y0)), draw.Over)
	return dst
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

func SetRectangle(srcimg image.Image, rcts ...*Recter) *image.NRGBA {
	r := srcimg.Bounds()
	imgNRGBA := image.NewNRGBA(r)
	draw.Draw(imgNRGBA, r, srcimg, image.ZP, draw.Src)

	for _, rct := range rcts {
		SetNRGBARectangle(imgNRGBA, rct)
	}
	return imgNRGBA
}

func SetNRGBARectangle(imgNRGBA *image.NRGBA, rct *Recter) {
	X, Y, xx, yy, c := rct.X, rct.Y, rct.xx, rct.yy, rct.c

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
}
