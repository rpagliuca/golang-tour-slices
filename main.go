package main

import (
	"math"
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	var y [][]uint8
	for iy := 0; iy < dy; iy++ {
		var x []uint8
		for ix := 0; ix < dx; ix++ {
			x = append(x, color(ix, iy, dx, dy))
		}
		y = append(y, x)
	}
	return y
}

func color(x, y, xMax, yMax int) uint8 {
	val := uint8(255)
	if x == y {
		val = 0
	}
	if x == yMax - y {
		val = 0
	}
	if y == yMax/2 {
		val = 0
	}
	if x == xMax/2 {
		val = 0
	}
	if math.Pow(float64(x - xMax/2), 2) + math.Pow(float64(y - yMax/2), 2) > math.Pow(float64(xMax)/2.2 - 1, 2) && math.Pow(float64(x - xMax/2), 2) + math.Pow(float64(y - yMax/2), 2) < math.Pow(float64(xMax)/2.2 + 1, 2)  {
	//if math.Pow(float64(x - xMax/2), 2) + math.Pow(float64(y - yMax/2), 2) == math.Pow(float64(xMax/3), 2)  {
		val = uint8((math.Pow(float64(x - xMax/2), 2) + math.Pow(float64(y - yMax/2), 2))/math.Pow(float64(xMax/3), 2) * 255)
		val = 240
	}
	if math.Pow(float64(x - xMax/2), 2) + math.Pow(float64(y - yMax/2), 2) < math.Pow(float64(xMax/3) - 1, 2) {
		//val = 255
		val = uint8((math.Pow(float64(x - xMax/2), 2) + math.Pow(float64(y - yMax/2), 2)) / math.Pow(float64(xMax/3) - 1, 2) * 255)
	}
	if x == xMax - 1 {
		val = 0
	}
	if x == 0 {
		val = 0
	}
	if y == 0 {
		val = 0
	}
	if y == yMax - 1 {
		val = 0
	}
	return val
}

func main() {
	pic.Show(Pic)
}
