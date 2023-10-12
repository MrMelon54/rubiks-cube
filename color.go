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

func (c Color) Byte() byte {
	if c == UnknownColor {
		return '?'
	}
	return colorByteTable[c]
}
