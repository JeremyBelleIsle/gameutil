package gameutil

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func GenerateBackgroundWithDegraded(screen *ebiten.Image, principalColor color.RGBA64, screenWidth, screenHeight float64, VOrH string) {

	currentColor := color.RGBA64{principalColor.R / 2, principalColor.G / 2, principalColor.B / 2, principalColor.A / 2}

	switch VOrH {
	case "h":
		for i := 0; i < int(screenWidth); i++ {
			vector.DrawFilledRect(screen, float32(i), 0, 1, float32(screenHeight), currentColor, false)
			r := uint16((principalColor.R / 2) / uint16(screenWidth))
			g := uint16((principalColor.G / 2) / uint16(screenWidth))
			b := uint16((principalColor.B / 2) / uint16(screenWidth))
			a := uint16((principalColor.A / 2) / uint16(screenWidth))

			currentColor = color.RGBA64{R: currentColor.R + r, G: currentColor.G + g, B: currentColor.B + b, A: currentColor.A + a}
		}
	case "v":
		for i := 0; i < int(screenHeight); i++ {
			vector.DrawFilledRect(screen, 0, float32(i), float32(screenWidth), 1, currentColor, false)
			r := uint16((principalColor.R / 2) / uint16(screenHeight))
			g := uint16((principalColor.G / 2) / uint16(screenHeight))
			b := uint16((principalColor.B / 2) / uint16(screenHeight))
			a := uint16((principalColor.A / 2) / uint16(screenHeight))

			currentColor = color.RGBA64{R: currentColor.R + r, G: currentColor.G + g, B: currentColor.B + b, A: currentColor.A + a}
		}
	}
}

func GenerateBackgroundWithDegradedCircle(screen *ebiten.Image, principalColor color.RGBA64, screenWidth, screenHeight float64) {
	maxRadius := screenHeight
	if screenWidth < screenHeight {
		maxRadius = screenWidth
	}

	// Dessiner du plus grand cercle au plus petit pour voir le dégradé
	for i := int(maxRadius); i >= 0; i-- {
		// Calculer un ratio de progression (0 à 1)
		ratio := float64(i) / maxRadius

		// Créer un dégradé de couleur en partant de la couleur principale
		currentColor := color.RGBA64{
			R: uint16(float64(principalColor.R) * ratio),
			G: uint16(float64(principalColor.G) * ratio),
			B: uint16(float64(principalColor.B) * ratio),
			A: principalColor.A,
		}

		vector.DrawFilledCircle(screen, float32(screenWidth)/2, float32(screenHeight)/2, float32(i), currentColor, false)
	}
}
