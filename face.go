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
