package gameutil

import "math"

func DirigePointToPoint(speed float32, x1, y1, x2, y2 float32) (float32, float32) {

	dx := x2 - x1
	dy := y2 - y1

	distance := float32(math.Sqrt(float64(dx*dx + dy*dy)))

	// Si déjà arrivé
	if distance <= speed || distance == 0 {
		return x2, y2
	}

	// Direction normalisée
	dx /= distance
	dy /= distance

	// Avance proportionnelle
	x1 += dx * speed
	y1 += dy * speed

	return x1, y1
}
