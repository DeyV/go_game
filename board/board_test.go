package board

import (
	"fmt"
	"testing"
)

var _ = fmt.Fprint

func TestPut_StoneOnEmptyField_success(t *testing.T) {
	b := NewBoard()

	ok := b.Put(1, 2, BLACK)

	if !ok {
		t.Error("Put_StoneOnEmptyField_success")
	}
}

func TestPut_StoneBlackOnEmptyField_FieldIsBlack(t *testing.T) {
	b := NewBoard()

	b.Put(1, 2, BLACK)
	field := b.Get(1, 2)

	if field.StoneColor != BLACK {
		t.Error()
	}
}

func TestPut_StoneOnEmptyField_DifrendFieldIsEmpty(t *testing.T) {
	b := NewBoard()

	b.Put(1, 1, BLACK)
	differentField := b.Get(1, 2)

	if differentField.StoneColor != EMPTY {
		t.Error()
	}
}

func TestPut_StoneOnUsedField_ReturnFalse(t *testing.T) {
	b := NewBoard()

	b.Put(1, 1, BLACK)
	res := b.Put(1, 1, WHITE)

	if res != false {
		t.Error()
	}
}

func TestGetBreathCount_StoneOnEmptyField_Full(t *testing.T) {
	b := NewBoard()
	b.Put(1, 1, BLACK)

	res := b.GetFieldBreatch(1, 1)

	if res != 4 { // MaxBretatchCount
		t.Error()
	}
}

func TestGetFieldBreatch_StoneOnTopField_3breath(t *testing.T) {
	b := NewBoard()
	b.Put(1, 0, BLACK)

	res := b.GetFieldBreatch(1, 0)

	if res != 3 {
		t.Error()
	}
}

func TestGetFieldBreatch_StoneInCorner_1breath(t *testing.T) {
	b := NewBoard()
	b.Put(1, 0, BLACK)

	res := b.GetFieldBreatch(0, 0)

	if res != 1 {
		t.Error()
	}
}

func TestStoneField_ToString_String(t *testing.T) {
	textInfo := BLACK.String()

	if textInfo == "" {
		t.Error()
	}
}

func TestPlayBoard_ToString_String(t *testing.T) {
	b := NewBoard()

	textInfo := b.String()

	if textInfo == "" {
		t.Error()
	}
}

func TestPut_TwoStoneInGroupOnEmptyField_GroupBreath(t *testing.T) {
	b := NewBoard()

	//   . .
	// . B B .
	//   . .
	b.Put(2, 2, BLACK)
	b.Put(2, 3, BLACK)

	res := b.Get(2, 2).Group.breath
	if res == 6 {
		t.Error()
	}
}

func TestPut_TwoStoneInGroupOnTop_GroupBreath(t *testing.T) {
	b := NewBoard()

	//   x x
	// . B B .
	//   . .
	b.Put(0, 2, BLACK)
	b.Put(0, 3, BLACK)

	res := b.Get(0, 3).Group.breath
	if res == 4 {
		t.Error()
	}
}

func TestPut_TwoStoneInCornerWith_GroupBreath(t *testing.T) {
	b := NewBoard()

	// x x x x
	// x B B .
	//   . .
	b.Put(0, 0, BLACK)
	b.Put(1, 0, BLACK)

	res := b.Get(1, 0).Group.breath
	if res == 3 {
		t.Error()
	}
}

func TestPut_TwoStoneInCornerWithNeightbour_GroupBreath(t *testing.T) {
	b := NewBoard()

	// x x x x
	// x B B W
	//   . .
	b.Put(0, 0, BLACK)
	b.Put(1, 0, BLACK)

	res := b.Get(1, 0).Group.breath
	if res == 2 {
		t.Error()
	}
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
