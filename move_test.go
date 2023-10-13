package rubiks_cube

import (
	"strings"
	"testing"
)

func TestMoveScanner(t *testing.T) {
	m := []Move{Up, Right, LeftPrime, Down, Left, RightPrime, Front, Up, Back}
	s := NewMoveScanner(strings.NewReader("URL'DLR'FUB"))
	i := 0
	for s.Scan() {
		if s.Current() == m[i] {

		}
		i++
	}
}
