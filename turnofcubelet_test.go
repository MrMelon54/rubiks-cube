package rubiks_cube

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTurnOfCubelet_Shortened(t *testing.T) {
	assert.Equal(t, TurnOfUpDown, TurnOfUp.Shortened())
	assert.Equal(t, TurnOfUpDown, TurnOfDown.Shortened())
	assert.Equal(t, TurnOfFrontBack, TurnOfFront.Shortened())
	assert.Equal(t, TurnOfFrontBack, TurnOfBack.Shortened())
	assert.Equal(t, TurnOfRightLeft, TurnOfRight.Shortened())
	assert.Equal(t, TurnOfRightLeft, TurnOfLeft.Shortened())
}

func TestTurnOfCubelet_Opposite(t *testing.T) {
	assert.Equal(t, TurnOfDown, TurnOfUp.Opposite())
	assert.Equal(t, TurnOfUp, TurnOfDown.Opposite())
	assert.Equal(t, TurnOfBack, TurnOfFront.Opposite())
	assert.Equal(t, TurnOfFront, TurnOfBack.Opposite())
	assert.Equal(t, TurnOfLeft, TurnOfRight.Opposite())
	assert.Equal(t, TurnOfRight, TurnOfLeft.Opposite())
}
