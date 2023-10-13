package rubiks_cube

// EdgeCubelet stores the type and rotation of a cubelet. Only 5 bits are used.
type EdgeCubelet byte

func MakeEdgeCubelet(eType EdgeType, eFace EdgeFacing) EdgeCubelet {
	return EdgeCubelet(eType&0b1111) | EdgeCubelet(eFace&0b1)<<4
}

func (e EdgeCubelet) Piece() EdgeType {
	return EdgeType(e & 0b1111)
}

func (e EdgeCubelet) Rotation() EdgeFacing {
	return EdgeFacing((e >> 4) & 0b1)
}

func (e EdgeCubelet) Valid() bool {
	return e.Piece().Valid() && e.Rotation().Valid()
}

func (e EdgeCubelet) Turn(p EdgePosition, t TurnOfCubelet) EdgeCubelet {
	z := e.Rotation().Turn(p, t)
	return MakeEdgeCubelet(e.Piece(), z)
}

func (e EdgeCubelet) GetColor(p EdgePosition, f Facing) Color {
	z := e.Piece()
	r := e.Rotation()
	r2 := p.StateOf(f)
	switch r {
	case EdgeNormal:
		return z.GetColor(r2)
	case EdgeOpposite:
		return z.GetColor((r2 + 1) % 2)
	}
	return UnknownColor
}

type EdgeType byte

const (
	EdgeWhiteOrange EdgeType = iota
	EdgeWhiteGreen
	EdgeWhiteRed
	EdgeWhiteBlue

	EdgeYellowOrange
	EdgeYellowGreen
	EdgeYellowRed
	EdgeYellowBlue

	EdgeOrangeGreen
	EdgeRedGreen
	EdgeRedBlue
	EdgeOrangeBlue
)

var edgeColorTable = [][2]Color{
	{White, Orange},
	{White, Green},
	{White, Red},
	{White, Blue},

	{Yellow, Orange},
	{Yellow, Green},
	{Yellow, Red},
	{Yellow, Blue},

	{Orange, Green},
	{Red, Green},
	{Red, Blue},
	{Orange, Blue},
}

func (t EdgeType) Valid() bool {
	return t <= EdgeOrangeBlue
}

func (t EdgeType) GetColor(f EdgeFacing) Color {
	return edgeColorTable[t][f]
}

func (t EdgeType) String() string {
	z := edgeColorTable[t]
	return "Edge" + z[0].String() + z[1].String()
}

func DetectEdge(a, b Color) EdgeCubelet {
	modeA := [2]Color{a, b}
	modeB := [2]Color{b, a}
	for i, j := range edgeColorTable {
		switch j {
		case modeA:
			return MakeEdgeCubelet(EdgeType(i), EdgeNormal)
		case modeB:
			return MakeEdgeCubelet(EdgeType(i), EdgeOpposite)
		}
	}
	return 255
}
