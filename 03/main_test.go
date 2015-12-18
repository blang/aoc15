package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestMove(t *testing.T) {
	assert := assert.New(t)
	assertPos := func(x1, y1, x2, y2 int) {
		assert.Equal(x1, x2, "X not right")
		assert.Equal(y1, y2, "Y not right")
	}
	var x, y int
	s := &Santa{}
	assertPos(0, 0, s.x, s.y)
	x, y = s.move('^')
	assertPos(0, 1, s.x, s.y)
	assertPos(0, 1, x, y)
	x, y = s.move('>')
	assertPos(1, 1, s.x, s.y)
	assertPos(1, 1, x, y)
	x, y = s.move('v')
	assertPos(1, 0, s.x, s.y)
	assertPos(1, 0, x, y)
	x, y = s.move('<')
	assertPos(0, 0, s.x, s.y)
	assertPos(0, 0, x, y)
}
