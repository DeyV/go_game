package board

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var _ = fmt.Fprint

func TestPut_StoneOnEmptyField_success(t *testing.T) {
	b := NewBoard()

	ok := b.Put(1, 2, BLACK)

	assert.True(t, ok)
}

func TestPut_StoneBlackOnEmptyField_FieldIsBlack(t *testing.T) {
	b := NewBoard()

	b.Put(1, 2, BLACK)
	field := b.Get(1, 2)

	assert.Equal(t, field.StoneColor, BLACK)
}

func TestPut_StoneOnEmptyField_DifrendFieldIsEmpty(t *testing.T) {
	b := NewBoard()

	b.Put(1, 1, BLACK)
	differentField := b.Get(1, 2)

	assert.Equal(t, differentField.StoneColor, EMPTY)
}

func TestPut_StoneOnUsedField_Panic(t *testing.T) {
	b := NewBoard()

	b.Put(1, 1, BLACK)

	assert.Panics(t, func() {
		_ = b.Put(1, 1, WHITE)
	})

}

func TestGetBreathCount_StoneOnEmptyField_Full(t *testing.T) {
	b := NewBoard()
	b.Put(1, 1, BLACK)

	res := b.GetFieldBreatch(1, 1)

	assert.Equal(t, 4, res)
}

func TestGetFieldBreatch_StoneOnTopField_3breath(t *testing.T) {
	b := NewBoard()
	b.Put(1, 0, BLACK)

	res := b.GetFieldBreatch(1, 0)

	assert.Equal(t, 3, res)
}

func TestGetFieldBreatch_StoneInCorner_1breath(t *testing.T) {
	b := NewBoard()
	b.Put(1, 0, BLACK)

	res := b.GetFieldBreatch(0, 0)

	assert.Equal(t, 1, res)
}

func TestStoneField_ToString_String(t *testing.T) {
	textInfo := BLACK.String()

	assert.NotEmpty(t, textInfo)
}

func TestPlayBoard_ToString_String(t *testing.T) {
	b := NewBoard()

	textInfo := b.String()

	assert.NotEmpty(t, textInfo)
}

func TestPut_OneStoneInGroupOnEmptyField_GroupBreath(t *testing.T) {
	b := NewBoard()

	//   .
	// . B .
	//   .
	b.Put(2, 2, BLACK)

	res := b.Get(2, 2).Group.breath

	assert.Equal(t, 4, res)
}

func TestPut_TwoStoneNextTo_OneGroup(t *testing.T) {
	b := NewBoard()

	//   . .
	// . B B .
	//   . .
	b.Put(2, 2, BLACK)
	b.Put(3, 2, BLACK)

	g1 := b.Get(2, 2).Group
	g2 := b.Get(3, 2).Group

	assert.Equal(t, g1, g2)
}

func TestPut_TwoStoneInGroupOnEmptyField_GroupBreath(t *testing.T) {
	b := NewBoard()

	//   . .
	// . B B .
	//   . .
	b.Put(2, 2, BLACK)
	b.Put(2, 3, BLACK)

	res := b.Get(2, 2).Group.breath

	assert.Equal(t, 6, res)
}

func TestPut_TwoStoneInGroupOnTop_GroupBreath(t *testing.T) {
	b := NewBoard()

	//   x x
	// . B B .
	//   . .
	b.Put(0, 2, BLACK)
	b.Put(0, 3, BLACK)

	res := b.Get(0, 3).Group.breath

	assert.Equal(t, 4, res)
}

/*
func TestPut_TwoStoneInCornerWith_GroupBreath(t *testing.T) {
	b := NewBoard()

	// x x x x
	// x B B .
	//   . .
	b.Put(0, 0, BLACK)
	b.Put(1, 0, BLACK)

	res := b.Get(1, 0).Group.breath

	assert.Equal(t, 3, res)
}

func TestPut_TwoStoneInCornerWithNeightbour_GroupBreath(t *testing.T) {
	b := NewBoard()

	// x x x x
	// x B B W
	//   . .
	b.Put(0, 0, BLACK)
	b.Put(1, 0, BLACK)
	b.Put(2, 0, WHITE)

	res := b.Get(1, 0).Group.breath

	assert.Equal(t, 2, res)
}

func TestPut_ThreeStoneInCornerWithNeightbour_GroupBreath(t *testing.T) {
	b := NewBoard()

	// x x x x x
	// x B B B W
	//   . . .
	b.Put(0, 0, BLACK)
	b.Put(1, 0, BLACK)
	b.Put(2, 0, BLACK)
	b.Put(3, 0, WHITE)

	res := b.Get(1, 0).Group.breath

	assert.Equal(t, 3, res)
}

/*
func TestPut_PutStoneAroundOtherInCorner_OtherDisaper(t *testing.T) {
	b := NewBoard()

	// W B
	// B
	b.Put(0, 0, WHITE)
	b.Put(1, 0, BLACK)
	b.Put(0, 1, BLACK)

	res := b.Get(0, 0)
	// fmt.Print(b.String())

	if !res.IsEmpty() {
		t.Error()
	}
}

*/
