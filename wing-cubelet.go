package rubiks_cube

// WingCubelet stores the type and rotation of a cubelet. Only 6 bits are used.
type WingCubelet byte

func MakeWingCubelet(wType WingType, wFace WingFacing) WingCubelet {
	return WingCubelet(wType&0b111) | WingCubelet(wFace&0b111)<<3
}

func (w WingCubelet) Piece() WingType {
	return WingType(w & 0b111)
}

func (w WingCubelet) Rotation() WingFacing {
	return WingFacing((w >> 3) & 0b111)
}

func (w WingCubelet) Valid() bool {
	return w.Piece().Valid() && w.Rotation().Valid()
}

func (w WingCubelet) Turn(t TurnOfCubelet) WingCubelet {
	z := w.Rotation().Turn(t)
	return MakeWingCubelet(w.Piece(), z)
}

func (w WingCubelet) GetColor() Color {
	return w.Piece().GetColor()
}

type WingType byte

const (
	WingWhite WingType = iota
	WingYellow
	WingOrange
	WingGreen
	WingRed
	WingBlue
)

func (t WingType) Valid() bool {
	return t <= WingBlue
}

func (t WingType) GetColor() Color {
	return Color(t)
}

func (t WingType) String() string {
	return "Wing" + t.GetColor().String()
}
