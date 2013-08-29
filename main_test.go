package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var _ = testing.AllocsPerRun

func TestOpenGame_CheckKeys_Letters_True(t *testing.T) {
	game := &OpenGame{}
	BoardSize := 9

	assert.True(t, game.CheckKey('a', BoardSize))
	assert.True(t, game.CheckKey('f', BoardSize))
	assert.True(t, game.CheckKey('i', BoardSize))
}

func TestOpenGame_CheckKeys_LetteresToBig_False(t *testing.T) {
	game := &OpenGame{}
	BoardSize := 9

	assert.False(t, game.CheckKey('j', BoardSize))
	assert.False(t, game.CheckKey('z', BoardSize))
}

func TestOpenGame_CheckKeys_Numbers_True(t *testing.T) {
	game := &OpenGame{}
	BoardSize := 9

	assert.True(t, game.CheckKey('1', BoardSize))
	assert.True(t, game.CheckKey('2', BoardSize))
}

func TestOpenGame_CheckKeys_NumbersToBig_false(t *testing.T) {
	game := &OpenGame{}
	BoardSize := 9

	assert.False(t, game.CheckKey('0', BoardSize))
}

func TestOpenGame_convertValue_A2_Ok(t *testing.T) {
	game := &OpenGame{}

	x, y := game.convertValue("A2")

	assert.Equal(t, 1, x)
	assert.Equal(t, 2, y)
}

func TestOpenGame_convertValue_a2_Ok(t *testing.T) {
	game := &OpenGame{}

	x, y := game.convertValue("a2")

	assert.Equal(t, 1, x)
	assert.Equal(t, 2, y)
}

func TestOpenGame_convertValue_B3_Ok(t *testing.T) {
	game := &OpenGame{}

	x, y := game.convertValue("B3")

	assert.Equal(t, 2, x)
	assert.Equal(t, 3, y)
}

func TestOpenGame_convertValue_AA23_Ok(t *testing.T) {
	game := &OpenGame{}

	x, y := game.convertValue("AA23")

	assert.Equal(t, 1, x)
	assert.Equal(t, 23, y)
}
