package rubiks_cube

//go:generate stringer -type Face

// Face defines the different faces of a cube.
type Face byte

const (
	FaceUp Face = iota
	FaceDown
	FaceFront
	FaceBack
	FaceRight
	FaceLeft
)

func (f Face) Valid() bool {
	return f <= FaceLeft
}

var faceToFacing = []Facing{
	FacingUpDown,
	FacingUpDown,
	FacingFrontBack,
	FacingFrontBack,
	FacingRightLeft,
	FacingRightLeft,
}

func (f Face) Facing() Facing {
	return faceToFacing[f]
}
