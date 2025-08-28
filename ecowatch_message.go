package main

import (
	"fmt"
	"log"
)

type Pixel map[string]string

type EcowatchMessage struct {
	Inner []Pixel `json:"inner"`
}

// MakePixelSlice Produces a slice of vectors filled with a single vector
func MakePixelSlice(color *V, length int) []Pixel {
	slice := make([]Pixel, 0)
	for i := range length {
		pixel := Pixel{fmt.Sprintf("%02d", i): color.ToInt().ToHex()}
		slice = append(slice, pixel)
	}
	return slice
}

func MakeGradientPixelSlice(startColour *V, endColour *V, length int) ([]Pixel, error) {
	slice := make([]Pixel, 0)
	gradient, err := startColour.IntInterpolate(endColour, length)
	if err != nil {
		return nil, fmt.Errorf("could not generate gradient: %w", err)
	}
	for i := range length {
		colour := gradient[i].ToHex()
		pixel := Pixel{fmt.Sprintf("%02d", i): colour}
		slice = append(slice, pixel)
	}
	return slice, nil
}

func MakeGradientProgressBar(startColour *V, endColour *V, length int, percent int) ([]Pixel, error) {
	slice, err := MakeGradientPixelSlice(startColour, endColour, length)
	if err != nil {
		log.Println("couldn't build gradient: %w", err)
		return nil, err
	}

	numberOfPixelsToBlank := (float64(percent) / float64(100.0)) * float64(length)

	blankV := IntV{0, 0, 0}
	for i := 0; i < int(numberOfPixelsToBlank); i++ {
		blankPixel := Pixel{fmt.Sprintf("%02d", i): blankV.ToHex()}
		slice[i] = blankPixel
	}

	return slice, nil
}
