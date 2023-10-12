package rubiks_cube

import "testing"

func TestHi(t *testing.T) {
	r := NewSolvedCube()
	println(r.String())
	r = r.RotateRight(false)
	println(r.String())
	r = r.RotateRight(false)
	println(r.String())
	r = r.RotateRight(false)
	println(r.String())
	r = r.RotateRight(false)
	println(r.String())
}
