package board

import (
	"bytes"
	"fmt"
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
	if !b.Get(x, y).IsEmpty() {
		// return false
		panic(fmt.Sprintf("this field ( %d, %d ) is not empty", x, y))
	}

	fieldBreath := b.GetFieldBreatch(x, y)
	fieldGroup := b.GetNearGroup(x, y, color)
	field := NewStoneField(color, fieldBreath, fieldGroup)

	fieldGroup.breath += fieldBreath

	if fieldGroup.breath <= 0 {
		panic(fmt.Sprintf("You cant put stone on field ( %d, %d ) ", x, y))
	}

	fieldGroup.size++

	b.fields[y][x] = field

	b.RemoveNeighborBreath(x, y)

	return true
}

func (b *playBoard) Get(x, y int) *StoneField {
	return &b.fields[y][x]
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

func (b *playBoard) visitNeighbors(x, y int, do func(*StoneField) bool) bool {
	for _, row := range neighbors {
		if b.onBoard(row[0]+x, row[1]+y) {
			field := b.Get(row[0]+x, row[1]+y)

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

	b.visitNeighbors(x, y, func(f *StoneField) bool {
		if f.IsEmpty() {
			result++
		}
		return true
	})

	return result
}

func (b *playBoard) GetNearGroup(x, y int, color StoneColor) *StoneGroup {
	var group *StoneGroup

	b.visitNeighbors(x, y, func(f *StoneField) bool {
		if !f.IsEmpty() && f.StoneColor == color {
			group = f.Group
			return false
		}
		return true
	})

	if group == nil {
		group = &StoneGroup{}
	}

	return group
}

func (b *playBoard) RemoveGroup(gr *StoneGroup) int {
	i := 0
	for y, row := range b.fields {
		for x, field := range row {
			if field.Group == gr {
				b.clearField(x, y)
				i++
			}
		}
	}
	return i
}

func (b *playBoard) clearField(x, y int) bool {
	field := b.Get(x, y)
	field.StoneColor = EMPTY

	group := field.Group

	b.visitNeighbors(x, y, func(f *StoneField) bool {
		if f.IsEmpty() {
			return true
		}

		if f.Group == group {
			return true // its group removed at this moment
		}

		f.ChangeBreath(1)
		return true
	})
	return true
}

func (b *playBoard) RemoveNeighborBreath(x, y int) bool {
	groupLive := true
	b.visitNeighbors(x, y, func(f *StoneField) bool {
		if f.IsEmpty() {
			return true
		}

		groupLive = f.ChangeBreath(-1)

		if !groupLive {
			b.RemoveGroup(f.Group)
		}

		return true
	})
	return true
}
