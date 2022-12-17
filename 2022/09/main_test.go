package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsTouching(t *testing.T) {
	k := Knot{positionX: 1, positionY: 1}
	o := Knot{positionX: 1, positionY: 1}
	t.Log("Same position")
	assert.True(t, k.IsTouching(&o))

	t.Log("k one to the right of o")
	k.positionX = k.positionX + 1
	assert.True(t, k.IsTouching(&o))

	t.Log("k two to the right of o")
	k.positionX = k.positionX + 1
	assert.False(t, k.IsTouching(&o))

	t.Log("k up and to right of o")
	k.positionX = k.positionX - 1
	k.positionY = k.positionY + 1
	assert.True(t, k.IsTouching(&o))

	t.Log("k above o")
	k.positionX = k.positionX - 1
	assert.True(t, k.IsTouching(&o))

	t.Log("k above and to the left of o")
	k.positionX = k.positionX - 1
	assert.True(t, k.IsTouching(&o))

	t.Log("k above and two to the left of o")
	k.positionX = k.positionX - 1
	assert.False(t, k.IsTouching(&o))

}
