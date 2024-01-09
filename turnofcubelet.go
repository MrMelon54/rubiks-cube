package rubiks_cube

//go:generate stringer -type TurnOfCubelet

// TurnOfCubelet defines the rotation of a cubelet relative to its current position.
type TurnOfCubelet byte

const (
	TurnOfUp TurnOfCubelet = iota
	TurnOfDown
	TurnOfFront
	TurnOfBack
	TurnOfRight
	TurnOfLeft
)

func (t TurnOfCubelet) Shortened() TurnOfCubeletShort {
	return TurnOfCubeletShort(t / 2)
}

func (t TurnOfCubelet) Opposite() TurnOfCubelet {
	return t/2 + (1 - t%2)
}

// TurnOfCubeletShort defines the rotation of a cubelet relative to its current position but shortened
type TurnOfCubeletShort byte

const (
	TurnOfUpDown TurnOfCubeletShort = iota
	TurnOfFrontBack
	TurnOfRightLeft
)
