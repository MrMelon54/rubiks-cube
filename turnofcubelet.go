package rubiks_cube

//go:generate stringer -type TurnOfCubelet

// TurnOfCubelet defines the rotation of a cubelet relative to its current position.
type TurnOfCubelet byte

const (
	TurnOfUpDown TurnOfCubelet = iota
	TurnOfFrontBack
	TurnOfRightLeft
)
