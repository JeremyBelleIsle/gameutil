package gameutil

import (
	"image/color"
	"math"
)

type Platform struct {
	X, Y float64
	W, H float64
	Clr  color.RGBA
}

func HorizontalCollisions(playerX, playerY, playerW, playerH float64, playerVelX *float64, platforms []Platform) float64 {
	newX := playerX

	for _, plat := range platforms {
		if RectColl(playerX, playerY, playerW, playerH, plat.X, plat.Y, plat.W, plat.H) {

			// 1. Calculer les chevauchements de chaque côté
			overlapLeft := (playerX + playerW) - plat.X  // Combien le joueur rentre par la gauche
			overlapRight := (plat.X + plat.W) - playerX  // Combien le joueur rentre par la droite
			overlapTop := (playerY + playerH) - plat.Y   // Combien le joueur rentre par le haut
			overlapBottom := (plat.Y + plat.H) - playerY // Combien le joueur rentre par le bas

			// 2. Trouver le plus petit chevauchement
			minOverlap := math.Min(
				math.Min(overlapLeft, overlapRight),
				math.Min(overlapTop, overlapBottom),
			)

			bounceFactor := 0.7

			// 3. Réagir selon le côté de la collision
			switch minOverlap {
			case overlapLeft: // Collision à gauche du mur
				newX = plat.X - playerW
				*playerVelX = -*playerVelX * bounceFactor

			case overlapRight: // Collision à droite du mur
				newX = plat.X + plat.W
				*playerVelX = -*playerVelX * bounceFactor

			// Pour les collisions verticales, on ne fait rien ici
			// (car c'est HorizontalCollisions)
			case overlapTop:
				// Tu pourrais gérer ça dans une fonction VerticalCollisions
			case overlapBottom:
				// Pareil
			}
		}
	}

	return newX
}
