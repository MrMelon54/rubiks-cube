package rubiks_cube

// CornerCubelet stores the type and rotation of a cubelet. Only 5 bits are used.
type CornerCubelet byte

func MakeCornerCubelet(cType CornerType, cFace Facing) CornerCubelet {
	return CornerCubelet(cType&0b111) | CornerCubelet(cFace&0b11)<<3
}

func (c CornerCubelet) Piece() CornerType {
	return CornerType(c & 0b111)
}

func (c CornerCubelet) Rotation() Facing {
	return Facing((c >> 3) & 0b11)
}

func (c CornerCubelet) Valid() bool {
	return c.Piece().Valid() && c.Rotation().Valid()
}

func (c CornerCubelet) Turn(t TurnOfCubelet) CornerCubelet {
	z := c.Rotation().Turn(t)
	return MakeCornerCubelet(c.Piece(), z)
}

func (c CornerCubelet) GetColor(f Facing) Color {
	z := c.Piece()
	switch c.Rotation() {
	case FacingUpDown:
		return z.GetColor(f)
	case FacingFrontBack:
		switch f {
		case FacingUpDown:
			return z.GetColor(FacingFrontBack)
		case FacingFrontBack:
			return z.GetColor(FacingUpDown)
		case FacingRightLeft:
			return z.GetColor(FacingRightLeft)
		}
	case FacingRightLeft:
		switch f {
		case FacingUpDown:
			return z.GetColor(FacingFrontBack)
		case FacingFrontBack:
			return z.GetColor(FacingRightLeft)
		case FacingRightLeft:
			return z.GetColor(FacingUpDown)
		}
	}
	return 255
}

type CornerType byte

const (
	CornerWhiteOrangeGreen CornerType = iota
	CornerWhiteRedGreen
	CornerWhiteRedBlue
	CornerWhiteOrangeBlue
	CornerYellowOrangeGreen
	CornerYellowRedGreen
	CornerYellowRedBlue
	CornerYellowOrangeBlue
)

var cornerColorTable = [][3]Color{
	{White, Orange, Green},
	{White, Red, Green},
	{White, Red, Blue},
	{White, Orange, Blue},
	{Yellow, Orange, Green},
	{Yellow, Red, Green},
	{Yellow, Red, Blue},
	{Yellow, Orange, Blue},
}

func (t CornerType) Valid() bool {
	return t <= CornerYellowOrangeBlue
}

func (t CornerType) GetColor(f Facing) Color {
	return cornerColorTable[t][f]
}

func (t CornerType) String() string {
	z := cornerColorTable[t]
	return "Corner" + z[0].String() + z[1].String() + z[2].String()
}

func DetectCorner(up, front, right Color) CornerCubelet {
	upFrontRight := [3]Color{up, front, right}
	frontRightUp := [3]Color{front, right, up}
	rightFrontUp := [3]Color{right, up, front}
	for i, j := range cornerColorTable {
		switch j {
		case upFrontRight:
			return MakeCornerCubelet(CornerType(i), FacingUpDown)
		case frontRightUp:
			return MakeCornerCubelet(CornerType(i), FacingFrontBack)
		case rightFrontUp:
			return MakeCornerCubelet(CornerType(i), FacingRightLeft)
		}
	}
	return 255
}
