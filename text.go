package gameutil

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func DrawText(textV string, LetterSize, screenWidth int, StartX, StartY float64, lineJump float64, screen *ebiten.Image, clr color.RGBA, font *text.GoTextFaceSource) {
	x := StartX
	y := StartY

	if lineJump <= float64(LetterSize)-5 {
		lineJump = float64(LetterSize) + 3
	}

	words := []string{}
	currentWord := ""

	for _, r := range textV {
		switch r {
		case ' ':
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}

		case ';':
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}
			words = append(words, ";")

		default:
			currentWord += string(r)
		}
	}

	if currentWord != "" {
		words = append(words, currentWord)
	}

	if len(words) == 0 {
		return
	}

	face := &text.GoTextFace{
		Source: font,
		Size:   float64(LetterSize),
	}

	for i := 0; i < len(words); i++ {
		if words[i] == ";" {
			x = StartX
			y += lineJump
			continue
		}

		wordWidth, _ := text.Measure(words[i], face, 0)

		if x+wordWidth > float64(screenWidth) && x != StartX {
			x = StartX
			y += lineJump
		}

		op := &text.DrawOptions{}
		op.GeoM.Translate(x, y)
		op.ColorScale.ScaleWithColor(clr)

		text.Draw(screen, words[i], face, op)

		x += wordWidth + 10
	}
}
