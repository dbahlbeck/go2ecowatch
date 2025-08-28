package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalEcowatchMessage(t *testing.T) {

	message := EcowatchMessage{
		Inner: []Pixel{
			{"00": "#000000"},
		},
	}

	jsonMessage, _ := json.Marshal(message)

	fmt.Println(jsonMessage)
}

func TestMakePixelSlice(t *testing.T) {
	pSlice := MakePixelSlice(&V{255, 0, 0}, 5)
	print(json.Marshal(pSlice))
}

func TestMakeEcowatchMessage(t *testing.T) {
	m := EcowatchMessage{
		Inner: MakePixelSlice(&V{255, 0, 0}, 5),
	}

	marshalledMessage, _ := json.Marshal(m)
	actual := string(marshalledMessage)
	expected := `{"inner":[{"00":"#FF0000"},{"01":"#FF0000"},{"02":"#FF0000"},{"03":"#FF0000"},{"04":"#FF0000"}]}`
	assert.Equal(t, expected, actual)

}

func TestMakeGradientPixelSliceWithTwoColours(t *testing.T) {
	g, err := MakeGradientPixelSlice(&V{x: 0, y: 0, z: 0}, &V{x: 0, y: 0, z: 1}, 2)
	assert.Nil(t, err)
	assert.Equal(t, []Pixel{{"00": "#000000"}, {"01": "#000001"}}, g)
}

func TestMakeGradientPixelSliceWithManyColours(t *testing.T) {
	g, err := MakeGradientPixelSlice(&V{x: 0, y: 0, z: 0}, &V{x: 0, y: 0, z: 5}, 5)
	assert.Nil(t, err)
	assert.Equal(t, []Pixel{
		{"00": "#000000"},
		{"01": "#000001"},
		{"02": "#000002"},
		{"03": "#000003"},
		{"04": "#000005"},
	}, g)
}

func TestMakeGradientProgressBar(t *testing.T) {
	g, err := MakeGradientProgressBar(&V{x: 0, y: 0, z: 0}, &V{x: 0, y: 0, z: 5}, 5, 50)
	assert.Nil(t, err)
	assert.Equal(t, []Pixel{
		{"00": "#000000"},
		{"01": "#000000"},
		{"02": "#000002"},
		{"03": "#000003"},
		{"04": "#000005"},
	}, g)
}

func TestMakeGradientProgressBarError(t *testing.T) {
	_, err := MakeGradientProgressBar(&V{x: 0, y: 0, z: 0}, &V{x: 0, y: 0, z: 5}, 1, 50)
	assert.ErrorContains(t, err, "need at least 2 steps")

}
