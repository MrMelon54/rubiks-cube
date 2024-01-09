package rubiks_cube

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseCube(t *testing.T) {
	cube, err := ParseCube(`   www
   www
   www
bbbooogggrrr
bbbooogggrrr
bbbooogggrrr
   yyy
   yyy
   yyy
`)
	assert.NoError(t, err)
	assert.Equal(t, NewSolvedCube(3), cube)
}
