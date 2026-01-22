package gameutil

import "math"

func DirigePointToPoint(speed float32, x1, y1, x2, y2 float32) (float32, float32) {
	if x1-float32(x2) < 0 {
		x1 += speed
	}
	if x1-float32(x2) > 0 {
		x1 -= speed
	}

	if math.Abs(float64(x1)-float64(x2)) <= float64(speed) {
		x1 = float32(x2)
	}
	// y
	if y1-float32(y2) < 0 {
		y1 += speed
	}
	if y1-float32(y2) > 0 {
		y1 -= speed
	}

	if math.Abs(float64(y1)-float64(y2)) <= float64(speed) {
		y1 = float32(y2)
	}

	return x1, y1
}
