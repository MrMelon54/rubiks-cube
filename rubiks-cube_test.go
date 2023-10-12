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

func TestHi2(t *testing.T) {
	r := NewSolvedCube()
	r.RightEdges[0] = MakeEdgeCubelet(EdgeWhiteOrange, EdgeOpposite)
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
