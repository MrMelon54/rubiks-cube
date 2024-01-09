package rubiks_cube

import (
	"strings"
)

// RubiksCube stores the state of a Rubik's Cube of size NxNxN using the type and
// rotation of each corner and edge cubelet.
//
// The cube is stored in the rotation where for the centers U = white, D =
// yellow, F = orange, B = red, R = green, L = blue
//
// Corners is split into Right and Left each with a size of 4, thus the type is an array with a constant size of 8.
// Corners[Right] is clockwise around the green face starting at the WhiteOrangeGreen corner.
// Corners[Left] is anti-clockwise around the blue face starting at the WhiteOrangeBlue corner.
//
// Edges is split into Right, Left and Middle each with a size of EdgeIndex*4.
// Edges[Right] is clockwise around the green face starting at the top.
// Edges[Left] is anti-clockwise around the blue face starting at the top.
// Edges[Middle] is clockwise around the green face for the final four edge pieces.
//
// Due to the design of the storage type, the solved state of every cubelet has
// the 0 rotation/facing value.
//
// Due to the nature of the wing cubelets they are stored as a naive 1-1
// representation.
type RubiksCube struct {
	N         int
	EdgeIndex int
	WingIndex int
	Corners   [8]CornerCubelet
	Edges     []EdgeCubelet
	Wings     []WingCubelet
}

// NewSolvedCube forms a solved Rubik's cube of size NxNxN. The default facing
// state of each cubelet follows the solved state of the cube so the raw byte
// value is cast directly to the cubelet because of the same lower bits are used
// to store the type.
func NewSolvedCube(N int) RubiksCube {
	if N < 2 {
		panic("Minimum cube size is 2")
	}

	smSize := N - 2
	r := RubiksCube{
		N:         N,
		EdgeIndex: smSize,
		WingIndex: smSize * smSize,
	}

	r.Corners = [8]CornerCubelet{
		CornerCubelet(CornerWhiteOrangeGreen),
		CornerCubelet(CornerWhiteRedGreen),
		CornerCubelet(CornerYellowRedGreen),
		CornerCubelet(CornerYellowOrangeGreen),
		CornerCubelet(CornerWhiteOrangeBlue),
		CornerCubelet(CornerWhiteRedBlue),
		CornerCubelet(CornerYellowRedBlue),
		CornerCubelet(CornerYellowOrangeBlue),
	}
	r.Edges = make([]EdgeCubelet, 12*r.EdgeIndex)
	r.Wings = make([]WingCubelet, 6*r.WingIndex)

	// edges
	for i := 0; i < r.EdgeIndex; i++ {
		// right edges
		r.Edges[i] = EdgeCubelet(EdgeWhiteGreen)
		r.Edges[r.EdgeIndex+i] = EdgeCubelet(EdgeRedGreen)
		r.Edges[r.EdgeIndex*2+i] = EdgeCubelet(EdgeYellowGreen)
		r.Edges[r.EdgeIndex*3+i] = EdgeCubelet(EdgeOrangeGreen)

		// left edges
		r.Edges[r.EdgeIndex*4+i] = EdgeCubelet(EdgeWhiteBlue)
		r.Edges[r.EdgeIndex*5+i] = EdgeCubelet(EdgeRedBlue)
		r.Edges[r.EdgeIndex*6+i] = EdgeCubelet(EdgeYellowBlue)
		r.Edges[r.EdgeIndex*7+i] = EdgeCubelet(EdgeOrangeBlue)

		// middle edges
		r.Edges[r.EdgeIndex*8+i] = EdgeCubelet(EdgeWhiteOrange)
		r.Edges[r.EdgeIndex*9+i] = EdgeCubelet(EdgeWhiteRed)
		r.Edges[r.EdgeIndex*10+i] = EdgeCubelet(EdgeYellowRed)
		r.Edges[r.EdgeIndex*11+i] = EdgeCubelet(EdgeYellowOrange)
	}

	// wing is useless for a 3x3
	if r.N == 3 {
		r.WingIndex = 0
	}

	// wings
	for i := 0; i < r.WingIndex; i++ {
		r.Wings[i] = WingCubelet(WingWhite)
		r.Wings[r.WingIndex+i] = WingCubelet(WingYellow)
		r.Wings[r.WingIndex*2+i] = WingCubelet(WingOrange)
		r.Wings[r.WingIndex*3+i] = WingCubelet(WingGreen)
		r.Wings[r.WingIndex*4+i] = WingCubelet(WingRed)
		r.Wings[r.WingIndex*5+i] = WingCubelet(WingBlue)
	}

	return r
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
	cycleCorners(prime, TurnOfRight, &r.Corners[0], &r.Corners[1], &r.Corners[2], &r.Corners[3])
	cycleItems(prime, &r.Edges[0], &r.Edges[1], &r.Edges[2], &r.Edges[3])
	turnEdge(&r.Edges[0], EdgeTopRight, TurnOfRight)
	turnEdge(&r.Edges[1], EdgeFrontRight, TurnOfRight)
	turnEdge(&r.Edges[2], EdgeTopRight, TurnOfRight)
	turnEdge(&r.Edges[3], EdgeFrontRight, TurnOfRight)
	return r
}

func (r RubiksCube) RotateLeft(prime bool) RubiksCube {
	cycleCorners(prime, TurnOfLeft, &r.Corners[1], &r.Corners[0], &r.Corners[3], &r.Corners[2])
	cycleItems(prime, &r.Edges[1], &r.Edges[0], &r.Edges[3], &r.Edges[2])
	turnEdge(&r.Edges[0], EdgeTopRight, TurnOfLeft)
	turnEdge(&r.Edges[1], EdgeFrontRight, TurnOfLeft)
	turnEdge(&r.Edges[2], EdgeTopRight, TurnOfLeft)
	turnEdge(&r.Edges[3], EdgeFrontRight, TurnOfLeft)
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

// faceCornerIndexes stores the indexes of the corners for each face
var faceCornerIndexes = [6 * 4]byte{
	5, 1, 4, 0,
	7, 3, 6, 2,
	4, 0, 7, 3,
	1, 5, 2, 6,
	0, 1, 3, 2,
	5, 4, 6, 7,
}

// faceEdgeIndexes stores the indexes of the edges for each face
// - right edges = 0 - 3
// - left edges = 4 - 7
// - middle edges = 8 - 11
var faceEdgeIndexes = [6 * 4]byte{
	0x09, 0x00, 0x08, 0x04,
	0x0b, 0x02, 0x0a, 0x06,
	0x08, 0x03, 0x0b, 0x07,
	0x09, 0x05, 0x0a, 0x01,
	0x00, 0x01, 0x02, 0x03,
	0x00, 0x03, 0x02, 0x01,
}

// faceEdgePosition stores the position data for the edges for each face
var faceEdgePosition = [3 * 4]EdgePosition{
	EdgeTopFront, EdgeTopRight, EdgeTopFront, EdgeTopRight,
	EdgeTopFront, EdgeFrontRight, EdgeTopFront, EdgeFrontRight,
	EdgeTopRight, EdgeFrontRight, EdgeTopRight, EdgeFrontRight,
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
	face = make(FaceData, r.N*r.N)
	facing := f.Facing()
	subN := r.N - 1
	bottomRow := r.N * subN

	// get indexes
	cornerIdx := faceCornerIndexes[4*f : 4*f+3]
	edgeIdx := faceEdgeIndexes[4*f : 4*f+3]
	edgePos := faceEdgePosition[4*facing : 4*facing+3]
	wingIdx := r.Wings[r.WingIndex*int(f) : r.WingIndex*int(f+1)]

	// find corner colors
	face[0] = r.Corners[cornerIdx[0]].GetColor(facing)
	face[subN] = r.Corners[cornerIdx[1]].GetColor(facing)
	face[bottomRow] = r.Corners[cornerIdx[2]].GetColor(facing)
	face[bottomRow+subN] = r.Corners[cornerIdx[3]].GetColor(facing)

	// find edge colors
	for i := 0; i < r.EdgeIndex; i++ {
		face[1+i] = r.Edges[edgeIdx[0]].GetColor(edgePos[0], facing)
		face[r.N+subN+r.N*i] = r.Edges[edgeIdx[1]].GetColor(edgePos[1], facing)
		face[bottomRow+1+i] = r.Edges[edgeIdx[2]].GetColor(edgePos[2], facing)
		face[r.N+r.N*i] = r.Edges[edgeIdx[3]].GetColor(edgePos[3], facing)
	}

	// find wing colors
	for i := 0; i < subN; i++ {
		for j := 0; j < subN; j++ {
			face[1+i+r.N*(j+1)] = wingIdx[i+subN*j].GetColor()
		}
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
