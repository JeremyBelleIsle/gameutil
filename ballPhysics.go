package gameutil

import (
	"image/color"
)

type Platform struct {
	X, Y float64
	W, H float64
	Clr  color.RGBA
}

func PlatformCollisions(playerX, playerY, playerR float64, playerVelX, playerVelY *float64, platforms []Platform) (float64, float64) {

	newX := playerX + *playerVelX
	newY := playerY + *playerVelY

	// ===== X AXIS =====
	for _, p := range platforms {
		if CircleRectCollision(newX, playerY, playerR, p.X, p.Y, p.W, p.H) {
			if *playerVelX > 0 {
				newX = p.X - playerR
			} else if *playerVelX < 0 {
				newX = p.X + p.W
			}
			*playerVelX = -*playerVelX * 0.6
		}
	}

	// ===== Y AXIS =====
	for _, p := range platforms {
		if CircleRectCollision(newX, newY, playerR, p.X, p.Y, p.W, p.H) {
			if *playerVelY > 0 {
				newY = p.Y - playerR
			} else if *playerVelY < 0 {
				newY = p.Y + p.H
			}
			*playerVelY = -*playerVelY * 0.6
		}
	}

	return newX, newY
}
