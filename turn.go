package rubiks_cube

// TurnOfCubelet defines the rotation of a cubelet relative to its current position.
type TurnOfCubelet byte

const (
	TurnOfUpDown TurnOfCubelet = iota
	TurnOfFrontBack
	TurnOfRightLeft
)
