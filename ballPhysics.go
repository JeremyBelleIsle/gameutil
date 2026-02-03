package gameutil

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Platform struct {
	x, y float64
	w, h float64
	clr  color.RGBA
}

func HorizontalCollisions(playerX, playerY, playerW, playerH float64, platforms []Platform) {
	for _, plat := range platforms {
		if RectColl(playerX, playerY, playerW, playerH, plat.x, plat.y, plat.w, plat.h) {
			if ebiten.IsKeyPressed(ebiten.KeyLeft) {
				playerX = plat.x + plat.w
			}
			if ebiten.IsKeyPressed(ebiten.KeyRight) {
				playerX = plat.x - playerW
			}
		}
	}
}
