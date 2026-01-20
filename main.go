package gameutil

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func DrawText(textV string, LetterSize int, StartX, StartY float64, lineJump float64, screen *ebiten.Image, clr color.RGBA) {
	x := StartX
	y := StartY

	if lineJump <= float64(LetterSize)-5 {
		lineJump = float64(LetterSize) + 3
	}

	// Découper le texte en mots ET garder les points
	words := []string{}
	currentWord := ""

	for i := 0; i < len(textV); i++ {
		if textV[i] == ' ' {
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}
		} else if textV[i] == ';' {
			// Ajouter le mot actuel s'il existe
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}
			// Ajouter le point comme un marqueur spécial
			words = append(words, ";")
		} else {
			currentWord += string(textV[i])
		}
	}

	if currentWord != "" {
		words = append(words, currentWord)
	}

	if len(words) == 0 {
		return
	}

	for i := 0; i < len(words); i++ {
		// Si c'est un point, faire un saut de ligne
		if words[i] == ";" {
			x = StartX
			y += lineJump
			continue
		}

		wordWidth := float64(len(words[i]) * LetterSize)

		if x+wordWidth > float64(screenWidth) && x != StartX {
			x = StartX
			y += lineJump
		}

		op := &text.DrawOptions{}
		op.GeoM.Translate(x, y)
		op.ColorScale.ScaleWithColor(clr)
		text.Draw(screen, words[i], &text.GoTextFace{
			Source: mplusFaceSource,
			Size:   float64(LetterSize),
		}, op)

		x += wordWidth + 10
	}
}
