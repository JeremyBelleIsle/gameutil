package gameutil

import (
	"image/color"
)

type Platform struct {
	X, Y float64
	W, H float64
	Clr  color.RGBA
}

func PlatformCollisions(playerX, playerY, playerW, playerH float64, playerVelX, playerVelY *float64, platforms []Platform) (float64, float64) {

	newX := playerX + *playerVelX
	newY := playerY + *playerVelY

	// ===== X AXIS =====
	for _, p := range platforms {
		if RectColl(newX, playerY, playerW, playerH, p.X, p.Y, p.W, p.H) {
			if *playerVelX > 0 {
				newX = p.X - playerW
			} else if *playerVelX < 0 {
				newX = p.X + p.W
			}
			*playerVelX = -*playerVelX * 0.7
		}
	}

	// ===== Y AXIS =====
	for _, p := range platforms {
		if RectColl(newX, newY, playerW, playerH, p.X, p.Y, p.W, p.H) {
			if *playerVelY > 0 {
				newY = p.Y - playerH
			} else if *playerVelY < 0 {
				newY = p.Y + p.H
			}
			*playerVelY = -*playerVelY * 0.3
		}
	}

	return newX, newY
}
