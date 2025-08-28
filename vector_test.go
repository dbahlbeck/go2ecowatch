package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	a := &V{x: 1, y: 1, z: 1}
	b := &V{x: 1, y: 2, z: 3}

	assert.Equal(t, a.Add(b), &V{x: 2, y: 3, z: 4})
}

func TestSub(t *testing.T) {
	a := &V{x: 1, y: 1, z: 1}
	b := &V{x: 1, y: 2, z: 3}

	assert.Equal(t, a.Sub(b), &V{x: 0, y: -1, z: -2})
}

func TestInterpolateNormalCase(t *testing.T) {
	// Arrange
	start := &V{x: 255, y: 0, z: 0}
	end := &V{x: 0, y: 255, z: 0}

	// Act
	result, err := start.IntInterpolate(end, 3)

	// Assert
	assert.Nil(t, err)
	expected := []*IntV{{x: 255, y: 0, z: 0}, {x: 127, y: 127, z: 0}, {x: 0, y: 255, z: 0}}
	assert.Equal(t, expected, result)
}

func TestInterpolateTwoSteps(t *testing.T) {
	// Arrange
	start := &V{x: 255, y: 0, z: 0}
	end := &V{x: 0, y: 255, z: 0}

	// Act
	result, err := start.IntInterpolate(end, 2)

	// Assert
	assert.Nil(t, err)
	expected := []*IntV{start.ToInt(), end.ToInt()}
	assert.Equal(t, expected, result)
}

func TestInterpolateOneStep(t *testing.T) {
	// Arrange
	start := &V{x: 255, y: 0, z: 0}
	end := &V{x: 0, y: 255, z: 0}

	// Act
	result, err := start.IntInterpolate(end, 1)

	// Assert
	assert.Nil(t, result)
	if assert.Error(t, err) {
		assert.Equal(t, "need at least 2 steps", err.Error())
	}
}

// func TestMakeFilledVectorSlice(t *testing.T) {
// 	// Act
// 	v := &V{255, 0, 0}
// 	vSlice := MakeFilledRGBSlice(v, 10)

// 	// Assert
// 	assert.Equal(t, 10, len(vSlice))
// 	for i := range vSlice {
// 		log.Println(vSlice[i])
// 	}

// }

func TestVToHex(t *testing.T) {
	v := IntV{x: 0, y: 0, z: 0}
	assert.Equal(t, "#000000", v.ToHex())
}
