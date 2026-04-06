package gameutil

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	bgRectImage   *ebiten.Image
	bgCircleImage *ebiten.Image
)

func InitBackgrounds(principalColor color.RGBA64, screenWidth, screenHeight int) {
	bgRectImage = ebiten.NewImage(screenWidth, screenHeight)
	drawRectDegraded(bgRectImage, principalColor, float64(screenWidth), float64(screenHeight))

	bgCircleImage = ebiten.NewImage(screenWidth, screenHeight)
	drawCircleDegraded(bgCircleImage, principalColor, float64(screenWidth), float64(screenHeight))
}

func DrawBackground(screen *ebiten.Image) {
	if bgRectImage != nil {
		screen.DrawImage(bgRectImage, nil)
	}
}

func DrawCircleBackground(screen *ebiten.Image) {
	if bgCircleImage != nil {
		screen.DrawImage(bgCircleImage, nil)
	}
}

func drawRectDegraded(dest *ebiten.Image, principalColor color.RGBA64, w, h float64) {
	currentColor := color.RGBA64{principalColor.R / 2, principalColor.G / 2, principalColor.B / 2, principalColor.A / 2}

	r := uint16((principalColor.R / 2) / uint16(w))
	g := uint16((principalColor.G / 2) / uint16(w))
	b := uint16((principalColor.B / 2) / uint16(w))
	a := uint16((principalColor.A / 2) / uint16(w))

	for i := 0; i < int(w); i++ {
		vector.DrawFilledRect(dest, float32(i), 0, 1, float32(h), currentColor, false)
		currentColor.R += r
		currentColor.G += g
		currentColor.B += b
		currentColor.A += a
	}
}

func drawCircleDegraded(dest *ebiten.Image, principalColor color.RGBA64, w, h float64) {
	maxRadius := h
	if w < h {
		maxRadius = w
	}

	for i := int(maxRadius); i >= 0; i-- {
		ratio := float64(i) / maxRadius
		currentColor := color.RGBA64{
			R: uint16(float64(principalColor.R) * ratio),
			G: uint16(float64(principalColor.G) * ratio),
			B: uint16(float64(principalColor.B) * ratio),
			A: principalColor.A,
		}
		vector.DrawFilledCircle(dest, float32(w)/2, float32(h)/2, float32(i), currentColor, false)
	}
}
