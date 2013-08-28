package main

import (
	"strconv"
	// "fmt"
	"github.com/DeyV/go_game/board"
	"github.com/nsf/termbox-go"
	"time"
)

func main() {
	b := board.NewBoard()

	// fmt.Println("START")

	// b.Put(2, 2, board.BLACK)
	// b.Put(2, 3, board.WHITE)
	// fmt.Println(b)

	// fmt.Println("END")

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	defer time.Sleep(2 * time.Second)

	b.Put(3, 3, board.BLACK)
	b.Put(3, 4, board.BLACK)
	b.Put(5, 4, board.WHITE)

	showBoardFrame(b)

	termbox.Flush()
}

var FirstRow = 2
var FirstCol = 5

func showBoardFrame(b *board.PlayBoard) {
	size := b.Size()

	// top label
	for i := 1; i <= size; i++ {
		nr := strconv.Itoa(i)
		showField(FirstCol+i*4, FirstRow, rune(nr[0]), termbox.ColorWhite, termbox.ColorBlue)
	}

	// left label
	for i := 1; i <= size; i++ {
		nr := strconv.Itoa(i)
		showField(FirstCol, FirstRow+2+(i-1)*2, rune(nr[0]), termbox.ColorWhite, termbox.ColorBlue)
	}

	showBoard(b)

}

func showBoard(b *board.PlayBoard) {
	size := b.Size()

	fg, bg := termbox.ColorWhite, termbox.ColorBlack

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			field := b.Get(x, y)

			xx := FirstCol + (x+1)*4
			yy := FirstRow + 2 + y*2

			if field.IsEmpty() {
				showField(xx, yy, ' ', fg, bg)
			} else {
				r := rune(field.String()[0])
				showField(xx, yy, r, fg, bg)
			}

			fg, bg = bg, fg
		}
	}
}

func showField(x, y int, ch rune, fg, bg termbox.Attribute) {
	termbox.SetCell(x, y, ' ', fg, bg)
	termbox.SetCell(x+1, y, ch, fg, bg)
	termbox.SetCell(x+2, y, ' ', fg, bg)
}
