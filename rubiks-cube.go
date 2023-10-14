package rubiks_cube

import (
	"strings"
)

// RubiksCube stores the state of a Rubik's Cube using the type and rotation of each corner and edge cubelet.
//
// The cube is stored in the rotation where for the centers U = white, D = yellow, F = orange, B = red, R = green, L = blue
//
// RightCorners is clockwise around the green face starting at the top right.
// RightEdges is clockwise around the green face starting at the top right.
// LeftCorners is anti-clockwise around the blue face starting at the top left.
// LeftEdges is anti-clockwise around the blue face starting at the top left.
// MiddleEdges is clockwise around the green face for the final four edge pieces.
//
// Due to the design of the storage type, the solved state of every cubelet has
// the 0 rotation/facing value.
type RubiksCube struct {
	RightCorners [4]CornerCubelet
	LeftCorners  [4]CornerCubelet
	RightEdges   [4]EdgeCubelet
	LeftEdges    [4]EdgeCubelet
	MiddleEdges  [4]EdgeCubelet
}

// NewSolvedCube forms a solved Rubik's cube. The default facing state of each
// cubelet follows the solved state of the cube so the raw byte value is cast
// directly to the cubelet because of the same lower bits are used to store the
// type.
func NewSolvedCube() RubiksCube {
	return RubiksCube{
		// corners
		[4]CornerCubelet{
			CornerCubelet(CornerWhiteOrangeGreen),
			CornerCubelet(CornerWhiteRedGreen),
			CornerCubelet(CornerYellowRedGreen),
			CornerCubelet(CornerYellowOrangeGreen),
		},
		[4]CornerCubelet{
			CornerCubelet(CornerWhiteOrangeBlue),
			CornerCubelet(CornerWhiteRedBlue),
			CornerCubelet(CornerYellowRedBlue),
			CornerCubelet(CornerYellowOrangeBlue),
		},

		// edges
		[4]EdgeCubelet{
			EdgeCubelet(EdgeWhiteGreen),
			EdgeCubelet(EdgeRedGreen),
			EdgeCubelet(EdgeYellowGreen),
			EdgeCubelet(EdgeOrangeGreen),
		},
		[4]EdgeCubelet{
			EdgeCubelet(EdgeWhiteBlue),
			EdgeCubelet(EdgeRedBlue),
			EdgeCubelet(EdgeYellowBlue),
			EdgeCubelet(EdgeOrangeBlue),
		},
		[4]EdgeCubelet{
			EdgeCubelet(EdgeWhiteOrange),
			EdgeCubelet(EdgeWhiteRed),
			EdgeCubelet(EdgeYellowRed),
			EdgeCubelet(EdgeYellowOrange),
		},
	}
}

func (r RubiksCube) Move(m Move) RubiksCube {
	switch m {
	case Up, UpPrime:
		return r.RotateUp(m.Prime())
	case Down, DownPrime:
		return r.RotateDown(m.Prime())
	case Front, FrontPrime:
		return r.RotateFront(m.Prime())
	case Back, BackPrime:
		return r.RotateBack(m.Prime())
	case Right, RightPrime:
		return r.RotateRight(m.Prime())
	case Left, LeftPrime:
		return r.RotateLeft(m.Prime())
	}
	return r
}

func (r RubiksCube) RotateUp(prime bool) RubiksCube {
	return r
}

func (r RubiksCube) RotateDown(prime bool) RubiksCube {
	return r
}

func (r RubiksCube) RotateFront(prime bool) RubiksCube {
	return r
}

func (r RubiksCube) RotateBack(prime bool) RubiksCube {
	return r
}

func (r RubiksCube) RotateRight(prime bool) RubiksCube {
	cycleCorners(prime, TurnOfRightLeft, &r.RightCorners[0], &r.RightCorners[1], &r.RightCorners[2], &r.RightCorners[3])
	cycleItems(prime, &r.RightEdges[0], &r.RightEdges[1], &r.RightEdges[2], &r.RightEdges[3])
	turnEdge(&r.RightEdges[0], EdgeTopRight, TurnOfRightLeft)
	turnEdge(&r.RightEdges[1], EdgeFrontRight, TurnOfRightLeft)
	turnEdge(&r.RightEdges[2], EdgeTopRight, TurnOfRightLeft)
	turnEdge(&r.RightEdges[3], EdgeFrontRight, TurnOfRightLeft)
	return r
}

func (r RubiksCube) RotateLeft(prime bool) RubiksCube {
	cycleCorners(prime, TurnOfRightLeft, &r.LeftCorners[1], &r.LeftCorners[0], &r.LeftCorners[3], &r.LeftCorners[2])
	cycleItems(prime, &r.LeftEdges[1], &r.LeftEdges[0], &r.LeftEdges[3], &r.LeftEdges[2])
	turnEdge(&r.LeftEdges[0], EdgeTopRight, TurnOfRightLeft)
	turnEdge(&r.LeftEdges[1], EdgeFrontRight, TurnOfRightLeft)
	turnEdge(&r.LeftEdges[2], EdgeTopRight, TurnOfRightLeft)
	turnEdge(&r.LeftEdges[3], EdgeFrontRight, TurnOfRightLeft)
	return r
}

func cycleItems[T comparable](reverse bool, a, b, c, d *T) {
	if reverse {
		z := *a
		*a = *b
		*b = *c
		*c = *d
		*d = z
	} else {
		z := *d
		*d = *c
		*c = *b
		*b = *a
		*a = z
	}
}

func cycleCorners(reverse bool, t TurnOfCubelet, a, b, c, d *CornerCubelet) {
	cycleItems(reverse, a, b, c, d)
	*a = a.Turn(t)
	*b = b.Turn(t)
	*c = c.Turn(t)
	*d = d.Turn(t)
}

func turnEdge(edge *EdgeCubelet, p EdgePosition, t TurnOfCubelet) {
	*edge = edge.Turn(p, t)
}

// Face returns the color of each cubelet on a specified face. face[4] will
// always be equal to c.
//
// The face will be returned as FaceData where indexes 0, 1, 2 is the first row
// of the face following the rotation display in the diagram below.
//
// . . . u u u . . . . . .
// . . . u u u . . . . . .
// . . . u u u . . . . . .
// l l l f f f r r r b b b
// l l l f f f r r r b b b
// l l l f f f r r r b b b
// . . . d d d . . . . . .
// . . . d d d . . . . . .
// . . . d d d . . . . . .
func (r RubiksCube) Face(f Face) (face FaceData) {
	face = FaceData{255, 255, 255, 255, 255, 255, 255, 255, 255}
	switch f {
	case FaceUp:
		face[0] = r.LeftCorners[1].GetColor(FacingUpDown)
		face[1] = r.MiddleEdges[1].GetColor(EdgeTopFront, FacingUpDown)
		face[2] = r.RightCorners[1].GetColor(FacingUpDown)
		face[3] = r.LeftEdges[0].GetColor(EdgeTopRight, FacingUpDown)
		face[4] = White
		face[5] = r.RightEdges[0].GetColor(EdgeTopRight, FacingUpDown)
		face[6] = r.LeftCorners[0].GetColor(FacingUpDown)
		face[7] = r.MiddleEdges[0].GetColor(EdgeTopFront, FacingUpDown)
		face[8] = r.RightCorners[0].GetColor(FacingUpDown)
	case FaceDown:
		face[0] = r.LeftCorners[3].GetColor(FacingUpDown)
		face[1] = r.MiddleEdges[3].GetColor(EdgeTopFront, FacingUpDown)
		face[2] = r.RightCorners[3].GetColor(FacingUpDown)
		face[3] = r.LeftEdges[2].GetColor(EdgeTopRight, FacingUpDown)
		face[4] = Yellow
		face[5] = r.RightEdges[2].GetColor(EdgeTopRight, FacingUpDown)
		face[6] = r.LeftCorners[2].GetColor(FacingUpDown)
		face[7] = r.MiddleEdges[2].GetColor(EdgeTopFront, FacingUpDown)
		face[8] = r.RightCorners[2].GetColor(FacingUpDown)
	case FaceFront:
		face[0] = r.LeftCorners[0].GetColor(FacingFrontBack)
		face[1] = r.MiddleEdges[0].GetColor(EdgeTopFront, FacingFrontBack)
		face[2] = r.RightCorners[0].GetColor(FacingFrontBack)
		face[3] = r.LeftEdges[3].GetColor(EdgeFrontRight, FacingFrontBack)
		face[4] = Orange
		face[5] = r.RightEdges[3].GetColor(EdgeFrontRight, FacingFrontBack)
		face[6] = r.LeftCorners[3].GetColor(FacingFrontBack)
		face[7] = r.MiddleEdges[3].GetColor(EdgeTopFront, FacingFrontBack)
		face[8] = r.RightCorners[3].GetColor(FacingFrontBack)
	case FaceBack:
		face[0] = r.RightCorners[1].GetColor(FacingFrontBack)
		face[1] = r.MiddleEdges[1].GetColor(EdgeTopFront, FacingFrontBack)
		face[2] = r.LeftCorners[1].GetColor(FacingFrontBack)
		face[3] = r.RightEdges[1].GetColor(EdgeFrontRight, FacingFrontBack)
		face[4] = Red
		face[5] = r.LeftEdges[1].GetColor(EdgeFrontRight, FacingFrontBack)
		face[6] = r.RightCorners[2].GetColor(FacingFrontBack)
		face[7] = r.MiddleEdges[2].GetColor(EdgeTopFront, FacingFrontBack)
		face[8] = r.LeftCorners[2].GetColor(FacingFrontBack)
	case FaceRight:
		face[0] = r.RightCorners[0].GetColor(FacingRightLeft)
		face[1] = r.RightEdges[0].GetColor(EdgeTopRight, FacingRightLeft)
		face[2] = r.RightCorners[1].GetColor(FacingRightLeft)
		face[3] = r.RightEdges[3].GetColor(EdgeFrontRight, FacingRightLeft)
		face[4] = Green
		face[5] = r.RightEdges[1].GetColor(EdgeFrontRight, FacingRightLeft)
		face[6] = r.RightCorners[3].GetColor(FacingRightLeft)
		face[7] = r.RightEdges[2].GetColor(EdgeTopRight, FacingRightLeft)
		face[8] = r.RightCorners[2].GetColor(FacingRightLeft)
	case FaceLeft:
		face[0] = r.LeftCorners[1].GetColor(FacingRightLeft)
		face[1] = r.LeftEdges[3].GetColor(EdgeFrontRight, FacingRightLeft)
		face[2] = r.LeftCorners[0].GetColor(FacingRightLeft)
		face[3] = r.LeftEdges[0].GetColor(EdgeTopRight, FacingRightLeft)
		face[4] = Blue
		face[5] = r.LeftEdges[2].GetColor(EdgeTopRight, FacingRightLeft)
		face[6] = r.LeftCorners[2].GetColor(FacingRightLeft)
		face[7] = r.LeftEdges[1].GetColor(EdgeFrontRight, FacingRightLeft)
		face[8] = r.LeftCorners[3].GetColor(FacingRightLeft)
	}

	return
}

func (r RubiksCube) String() string {
	var z CubeFaceData
	for i := 0; i < 6; i++ {
		z[i] = r.Face(Face(i))
	}

	var s strings.Builder
	s.Grow(13 * 9)
	for i := 0; i < 9; i += 3 {
		s.WriteString("   ")
		s.WriteByte(z[FaceUp][i].Byte())
		s.WriteByte(z[FaceUp][i+1].Byte())
		s.WriteByte(z[FaceUp][i+2].Byte())
		s.WriteString("      \n")
	}
	for i := 0; i < 9; i += 3 {
		s.WriteByte(z[FaceLeft][i].Byte())
		s.WriteByte(z[FaceLeft][i+1].Byte())
		s.WriteByte(z[FaceLeft][i+2].Byte())
		s.WriteByte(z[FaceFront][i].Byte())
		s.WriteByte(z[FaceFront][i+1].Byte())
		s.WriteByte(z[FaceFront][i+2].Byte())
		s.WriteByte(z[FaceRight][i].Byte())
		s.WriteByte(z[FaceRight][i+1].Byte())
		s.WriteByte(z[FaceRight][i+2].Byte())
		s.WriteByte(z[FaceBack][i].Byte())
		s.WriteByte(z[FaceBack][i+1].Byte())
		s.WriteByte(z[FaceBack][i+2].Byte())
		s.WriteByte('\n')
	}
	for i := 0; i < 9; i += 3 {
		s.WriteString("   ")
		s.WriteByte(z[FaceDown][i].Byte())
		s.WriteByte(z[FaceDown][i+1].Byte())
		s.WriteByte(z[FaceDown][i+2].Byte())
		s.WriteString("      \n")
	}

	return s.String()
}
