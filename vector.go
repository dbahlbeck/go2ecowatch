package main

import (
	"errors"
	"fmt"
)

type IntV struct {
	x, y, z int
}

type V struct {
	x, y, z float64
}

func (source *V) Sub(target *V) *V {
	return &V{
		x: source.x - target.x,
		y: source.y - target.y,
		z: source.z - target.z,
	}
}

func (source *V) Mult(f float64) *V {
	return &V{
		x: source.x * f,
		y: source.y * f,
		z: source.z * f,
	}
}

func (source *V) Add(target *V) *V {
	return &V{
		x: source.x + target.x,
		y: source.y + target.y,
		z: source.z + target.z,
	}
}

func (source *V) ToInt() *IntV {
	return &IntV{
		x: int(source.x),
		y: int(source.y),
		z: int(source.z),
	}
}

func (source *IntV) ToHex() string {
	return fmt.Sprintf("#%02X%02X%02X", source.x, source.y, source.z)
}

// IntInterpolate takes a position in 3D space and produces step_count points along the line
// from (x1,y1,z1) to (x2, y2, z2) include the start and end.
func (source *V) IntInterpolate(target *V, stepCount int) ([]*IntV, error) {
	if stepCount < 2 {
		return nil, errors.New("need at least 2 steps")
	}

	// get the direction from start to end
	d := target.Sub(source)
	// interpolate
	points := make([]*IntV, stepCount)

	for i := range stepCount {
		step := float64(i) * float64(1.0/float64(stepCount-1))
		//fmt.Println(step)
		points[i] = (source.Add(d.Mult(step))).ToInt()
	}

	return points, nil
}
