package board

import (
	"fmt"
	"strconv"
)

type StoneGroup struct {
	size   int
	breath int8
}

type StoneField struct {
	StoneColor
	breath int8
	Group  *StoneGroup
}

func NewStoneField(c StoneColor, breathCount int8, group *StoneGroup) StoneField {
	return StoneField{StoneColor: c, breath: breathCount, Group: group}
}

func (s *StoneField) ChangeBreath(val int8) bool {
	s.breath += val

	s.Group.breath += val

	fmt.Println(val, s.breath, s.Group.breath)

	if s.Group.breath <= 0 {
		return false
	}

	return true
}

func (s *StoneField) String() string {
	return s.StoneColor.String() + strconv.Itoa(int(s.breath))
}

type StoneColor int

const (
	EMPTY StoneColor = iota
	BLACK
	WHITE
)

func (c StoneColor) IsEmpty() bool {
	return c == EMPTY
}

func (c StoneColor) String() string {
	if c == WHITE {
		return "W"
	}

	if c == BLACK {
		return "B"
	}

	return "-"
}
