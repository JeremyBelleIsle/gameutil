package gameutil

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

var audioContext = audio.NewContext(44100)

func LoadSound(path string) (*audio.Player, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot read %q: %w", path, err)
	}

	ext := strings.ToLower(filepath.Ext(path))

	switch ext {

	case ".wav":
		stream, err := wav.DecodeWithSampleRate(44000, bytes.NewReader(data))
		// stream, err := wav.Decode(audioContext, bytes.NewReader(data))
		if err != nil {
			return nil, fmt.Errorf("cannot decode wav %q: %w", path, err)
		}
		fmt.Printf("Sound load with success!")
		return audioContext.NewPlayer(stream)
		// return audio.NewPlayer(audioContext, stream)

	case ".mp3":
		stream, err := mp3.DecodeWithSampleRate(44100, bytes.NewReader(data))
		if err != nil {
			return nil, fmt.Errorf("cannot decode mp3 %q: %w", path, err)
		}
		return audioContext.NewPlayer(stream)

	default:
		return nil, fmt.Errorf("unsupported audio extension %q", ext)
	}
}

func main() {
	audioContext = audio.NewContext(44100)
}
