package rubiks_cube

//go:generate stringer -type Facing,EdgeFacing,EdgePosition

type Facing byte

const (
	FacingUpDown Facing = iota
	FacingFrontBack
	FacingRightLeft
)

func (f Facing) Valid() bool {
	return f <= FacingRightLeft
}

var rotationTable = [][]Facing{
	FacingUpDown: {
		TurnOfUpDown:    FacingUpDown,
		TurnOfFrontBack: FacingRightLeft,
		TurnOfRightLeft: FacingFrontBack,
	},
	FacingFrontBack: {
		TurnOfUpDown:    FacingRightLeft,
		TurnOfFrontBack: FacingFrontBack,
		TurnOfRightLeft: FacingUpDown,
	},
	FacingRightLeft: {
		TurnOfUpDown:    FacingFrontBack,
		TurnOfFrontBack: FacingUpDown,
		TurnOfRightLeft: FacingRightLeft,
	},
}

func (f Facing) Turn(t TurnOfCubelet) Facing {
	return rotationTable[f][t.Shortened()]
}

// EdgeFacing is the direction which the white/yellow side of an edge will be
// pointing. EdgeNormal represents UP/DOWN if that direction is available,
// otherwise EdgeNormal becomes the FRONT/BACK direction. For edge pieces without
// a white/yellow face, the edge will be in the EdgeNormal direction when the
// cube is in the solved state.
type EdgeFacing byte

const (
	EdgeNormal EdgeFacing = iota
	EdgeOpposite
)

func (f EdgeFacing) Valid() bool { return f <= EdgeOpposite }

var edgeAtFacingTable = [][]Facing{
	EdgeTopFront: {
		EdgeNormal:   FacingUpDown,
		EdgeOpposite: FacingFrontBack,
	},
	EdgeTopRight: {
		EdgeNormal:   FacingUpDown,
		EdgeOpposite: FacingRightLeft,
	},
	EdgeFrontRight: {
		EdgeNormal:   FacingFrontBack,
		EdgeOpposite: FacingRightLeft,
	},
}

func (f EdgeFacing) At(p EdgePosition) Facing {
	return edgeAtFacingTable[p][f]
}

func (f EdgeFacing) Turn(p EdgePosition, t TurnOfCubelet) EdgeFacing {
	return p.Turn(t).StateOf(f.At(p).Turn(t))
}

type EdgePosition byte

const (
	EdgeTopFront EdgePosition = iota
	EdgeTopRight
	EdgeFrontRight
)

var edgeRotationTable = [][]EdgePosition{
	EdgeTopFront: {
		TurnOfUpDown:    EdgeTopRight,
		TurnOfFrontBack: EdgeFrontRight,
		TurnOfRightLeft: EdgeTopFront,
	},
	EdgeTopRight: {
		TurnOfUpDown:    EdgeTopFront,
		TurnOfFrontBack: EdgeTopRight,
		TurnOfRightLeft: EdgeFrontRight,
	},
	EdgeFrontRight: {
		TurnOfUpDown:    EdgeFrontRight,
		TurnOfFrontBack: EdgeTopFront,
		TurnOfRightLeft: EdgeTopRight,
	},
}

func (p EdgePosition) Turn(t TurnOfCubelet) EdgePosition {
	return edgeRotationTable[p][t.Shortened()]
}

var edgeStateOfTable = [][]EdgeFacing{
	EdgeTopFront: {
		FacingUpDown:    EdgeNormal,
		FacingFrontBack: EdgeOpposite,
		FacingRightLeft: 255,
	},
	EdgeTopRight: {
		FacingUpDown:    EdgeNormal,
		FacingFrontBack: 255,
		FacingRightLeft: EdgeOpposite,
	},
	EdgeFrontRight: {
		FacingUpDown:    255,
		FacingFrontBack: EdgeNormal,
		FacingRightLeft: EdgeOpposite,
	},
}

func (p EdgePosition) StateOf(f Facing) EdgeFacing {
	return edgeStateOfTable[p][f]
}

type WingFacing byte

const (
	WingUp WingFacing = iota
	WingDown
	WingFront
	WingBack
	WingRight
	WingLeft
)

func (w WingFacing) Valid() bool { return w <= WingLeft }

var wingRotationTable = [][]WingFacing{
	WingUp: {
		TurnOfUp:    WingUp,
		TurnOfDown:  WingUp,
		TurnOfFront: WingRight,
		TurnOfBack:  WingLeft,
		TurnOfRight: WingBack,
		TurnOfLeft:  WingFront,
	},
	WingDown: {
		TurnOfUp:    WingUp,
		TurnOfDown:  WingUp,
		TurnOfFront: WingLeft,
		TurnOfBack:  WingRight,
		TurnOfRight: WingFront,
		TurnOfLeft:  WingBack,
	},
	WingFront: {
		TurnOfUp:    WingLeft,
		TurnOfDown:  WingRight,
		TurnOfFront: WingFront,
		TurnOfBack:  WingFront,
		TurnOfRight: WingUp,
		TurnOfLeft:  WingDown,
	},
	WingBack: {
		TurnOfUp:    WingRight,
		TurnOfDown:  WingLeft,
		TurnOfFront: WingBack,
		TurnOfBack:  WingBack,
		TurnOfRight: WingDown,
		TurnOfLeft:  WingUp,
	},
	WingRight: {
		TurnOfUp:    WingFront,
		TurnOfDown:  WingBack,
		TurnOfFront: WingDown,
		TurnOfBack:  WingUp,
		TurnOfRight: WingRight,
		TurnOfLeft:  WingRight,
	},
	WingLeft: {
		TurnOfUp:    WingBack,
		TurnOfDown:  WingFront,
		TurnOfFront: WingUp,
		TurnOfBack:  WingDown,
		TurnOfRight: WingLeft,
		TurnOfLeft:  WingLeft,
	},
}

func (w WingFacing) Turn(t TurnOfCubelet) WingFacing {
	return wingRotationTable[w][t]
}
