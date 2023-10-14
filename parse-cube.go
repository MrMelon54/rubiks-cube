package rubiks_cube

import (
	"errors"
)

var ErrInvalidCubeState = errors.New("invalid cube state")

func ParseCube(v string) (RubiksCube, error) {
	faces, err := ParseFaces(v)
	if err != nil {
		return RubiksCube{}, err
	}

	cube := RubiksCube{
		// corners
		[4]CornerCubelet{
			DetectCorner(faces[FaceUp][8], faces[FaceFront][2], faces[FaceRight][0]),
			DetectCorner(faces[FaceUp][2], faces[FaceBack][0], faces[FaceRight][2]),
			DetectCorner(faces[FaceDown][8], faces[FaceBack][6], faces[FaceRight][8]),
			DetectCorner(faces[FaceDown][2], faces[FaceFront][8], faces[FaceRight][6]),
		},
		[4]CornerCubelet{
			DetectCorner(faces[FaceUp][6], faces[FaceFront][0], faces[FaceLeft][2]),
			DetectCorner(faces[FaceUp][0], faces[FaceBack][2], faces[FaceLeft][0]),
			DetectCorner(faces[FaceDown][6], faces[FaceBack][8], faces[FaceLeft][6]),
			DetectCorner(faces[FaceDown][0], faces[FaceFront][6], faces[FaceLeft][8]),
		},

		// edges
		[4]EdgeCubelet{
			DetectEdge(faces[FaceUp][5], faces[FaceRight][1]),
			DetectEdge(faces[FaceBack][3], faces[FaceRight][5]),
			DetectEdge(faces[FaceDown][5], faces[FaceRight][7]),
			DetectEdge(faces[FaceFront][5], faces[FaceRight][3]),
		},
		[4]EdgeCubelet{
			DetectEdge(faces[FaceUp][3], faces[FaceLeft][1]),
			DetectEdge(faces[FaceBack][5], faces[FaceLeft][3]),
			DetectEdge(faces[FaceDown][3], faces[FaceLeft][7]),
			DetectEdge(faces[FaceFront][3], faces[FaceLeft][5]),
		},
		[4]EdgeCubelet{
			DetectEdge(faces[FaceUp][7], faces[FaceFront][1]),
			DetectEdge(faces[FaceUp][1], faces[FaceBack][1]),
			DetectEdge(faces[FaceDown][7], faces[FaceBack][7]),
			DetectEdge(faces[FaceDown][1], faces[FaceFront][7]),
		},
	}
	for _, i := range cube.RightCorners {
		if i == 255 {
			return RubiksCube{}, ErrInvalidCubeState
		}
	}
	for _, i := range cube.LeftCorners {
		if i == 255 {
			return RubiksCube{}, ErrInvalidCubeState
		}
	}
	for _, i := range cube.RightEdges {
		if i == 255 {
			return RubiksCube{}, ErrInvalidCubeState
		}
	}
	for _, i := range cube.LeftEdges {
		if i == 255 {
			return RubiksCube{}, ErrInvalidCubeState
		}
	}
	for _, i := range cube.MiddleEdges {
		if i == 255 {
			return RubiksCube{}, ErrInvalidCubeState
		}
	}

	return cube, nil
}
