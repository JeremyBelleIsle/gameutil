package gameutil

import "math"

func Within(px, py, rx, ry, rw, rh float64) bool {
	return px >= rx && px <= rx+rw && py >= ry && py <= ry+rh
}

func CircleRectCollision(cx, cy, cr, rx, ry, rw, rh float64) bool {
	closestX := math.Max(rx, math.Min(cx, rx+rw))
	closestY := math.Max(ry, math.Min(cy, ry+rh))
	dx := cx - closestX
	dy := cy - closestY

	return dx*dx+dy*dy <= cr*cr
}
