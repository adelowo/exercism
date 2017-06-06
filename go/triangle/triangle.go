package triangle

import (
	"math"
	"sort"
)

const testVersion = 3

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

type Kind int

func KindFromSides(a, b, c float64) Kind {

	data := []float64{a, b, c}

	sort.Float64s(data)

	if data[0]+data[1] < data[2] || math.IsNaN(data[0]) ||
		data[0] <= 0 || math.Inf(1) == data[2] {
		return NaT
	}

	if data[0] == data[1] && data[1] == data[2] {
		return Equ
	}

	if data[0] == data[1] || data[1] == data[2] {
		return Iso
	}

	return Sca
}
