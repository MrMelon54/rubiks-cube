package rubiks_cube

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWingType_GetColor(t *testing.T) {
	assert.Equal(t, White, WingWhite.GetColor())
	assert.Equal(t, Yellow, WingYellow.GetColor())
	assert.Equal(t, Orange, WingOrange.GetColor())
	assert.Equal(t, Green, WingGreen.GetColor())
	assert.Equal(t, Red, WingRed.GetColor())
	assert.Equal(t, Blue, WingBlue.GetColor())
}

func TestWingType_String(t *testing.T) {
	assert.Equal(t, "WingWhite", WingWhite.String())
	assert.Equal(t, "WingYellow", WingYellow.String())
	assert.Equal(t, "WingOrange", WingOrange.String())
	assert.Equal(t, "WingGreen", WingGreen.String())
	assert.Equal(t, "WingRed", WingRed.String())
	assert.Equal(t, "WingBlue", WingBlue.String())
}
