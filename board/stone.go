package board

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
