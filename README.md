# rect

golang image draw rect

## usage

```		
r := NewRectangle(50, 110, 130, 190, color.RGBA{
		A: 255,
		R: 255,
		G: 1,
		B: 1,
	}, 5)
DrawRectangle(img, r)


SetRectangle(img,
	NewRecter(50, 15, 90, 45, color.RGBA{
		A: 255,
		R: 255,
	}, 1),
	NewRecter(50, 110, 130, 190, color.RGBA{
		A: 255,
		R: 255,
	}, 5))
```		

## test

![](https://raw.githubusercontent.com/toukii/rect/master/gosea.jpg)

![](https://raw.githubusercontent.com/toukii/rect/master/test.png)
