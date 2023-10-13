package rubiks_cube

//go:generate stringer -type Color

type Color byte

const (
	White Color = iota
	Yellow
	Orange
	Green
	Red
	Blue
	UnknownColor Color = 255
)

var colorByteTable = [6]byte{'w', 'y', 'o', 'g', 'r', 'b'}

func (c Color) Valid() bool {
	return c <= Blue
}

func (c Color) Byte() byte {
	if c == UnknownColor {
		return '?'
	}
	return colorByteTable[c]
}

func ParseColor(c byte) Color {
	switch c {
	case 'w':
		return White
	case 'y':
		return Yellow
	case 'o':
		return Orange
	case 'g':
		return Green
	case 'r':
		return Red
	case 'b':
		return Blue
	}
	return UnknownColor
}
