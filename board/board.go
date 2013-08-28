package board

import (
	"bytes"
	"strconv"
)

type playBoard struct {
	fields [9][9]StoneField
	size   int
}

func NewBoard() *playBoard {
	return &playBoard{size: 9}
}

func (b *playBoard) Put(x, y int, color StoneColor) bool {
	if !b.fields[y][x].IsEmpty() {
		return false
	}

	fieldBreath := b.GetFieldBreatch(x, y)
	fieldGroup := b.GetFieldGroup(x, y)
	field := NewStoneField(color, fieldBreath, fieldGroup)

	b.fields[y][x] = field
	return true
}

func (b *playBoard) Get(x, y int) StoneField {
	return b.fields[y][x]
}

func (b *playBoard) String() string {
	var buff bytes.Buffer

	for i, row := range b.fields {
		buff.WriteString(strconv.Itoa(i) + ".\t")

		for _, field := range row {
			buff.WriteString(field.String() + "\t")
		}

		buff.WriteString("\n")
	}
	return buff.String()
}

var neighbors = [...][2]int{
	{0, -1},
	{-1, 0},
	{0, 1},
	{1, 0},
}

func (b *playBoard) visitNeighbors(x, y int, do func(StoneField) bool) bool {

	for _, row := range neighbors {
		if b.onBoard(row[0]+y, row[1]+x) {
			field := b.Get(row[0]+y, row[1]+x)

			toBeCon := do(field)

			if toBeCon != true {
				return false
			}
		}
	}

	return true
}

func (b *playBoard) onBoard(x, y int) bool {
	if x < 0 || x >= b.size {
		return false
	}
	if y < 0 || y >= b.size {
		return false
	}

	return true
}

func (b *playBoard) GetFieldBreatch(x, y int) int8 {
	var result int8 = 0

	b.visitNeighbors(x, y, func(f StoneField) bool {
		if f.IsEmpty() {
			result++
		}
		return true
	})

	return result
}

func (b *playBoard) GetFieldGroup(x, y int) *StoneGroup {
	var group *StoneGroup
	myColor := b.Get(x, y).StoneColor

	b.visitNeighbors(x, y, func(f StoneField) bool {
		if !f.IsEmpty() && f.StoneColor == myColor {
			group = f.Group
		}
		return true
	})

	if group == nil {
		group = &StoneGroup{}
	}

	return group
}
