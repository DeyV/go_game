package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var _ = testing.AllocsPerRun

func TestCheckKeys_Letters_True(t *testing.T) {
	game := &OpenGame{}
	BoardSize := 9

	assert.True(t, game.CheckKey('a', BoardSize))
	assert.True(t, game.CheckKey('f', BoardSize))
	assert.True(t, game.CheckKey('i', BoardSize))
}

func TestCheckKeys_LetteresToBig_False(t *testing.T) {
	game := &OpenGame{}
	BoardSize := 9

	assert.False(t, game.CheckKey('j', BoardSize))
	assert.False(t, game.CheckKey('z', BoardSize))
}

func TestCheckKeys_Numbers_True(t *testing.T) {
	game := &OpenGame{}
	BoardSize := 9

	assert.True(t, game.CheckKey('1', BoardSize))
	assert.True(t, game.CheckKey('2', BoardSize))
}

func TestCheckKeys_NumbersToBig_false(t *testing.T) {
	game := &OpenGame{}
	BoardSize := 9

	assert.False(t, game.CheckKey('0', BoardSize))
}
