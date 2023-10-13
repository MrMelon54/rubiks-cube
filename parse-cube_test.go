package rubiks_cube

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestParseCube(t *testing.T) {
	cube, err := ParseCube(strings.NewReader(`   www
   www
   www
bbbooogggrrr
bbbooogggrrr
bbbooogggrrr
   yyy
   yyy
   yyy
`))
	assert.NoError(t, err)
	assert.Equal(t, NewSolvedCube(), cube)
}
