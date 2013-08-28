package main

import (
	"strconv"
	"strings"
	// "fmt"
	"github.com/DeyV/go_game/board"
	"github.com/nsf/termbox-go"
)

func main() {
	b := board.NewBoard()

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	// b.Put(3, 3, board.BLACK)
	// b.Put(3, 4, board.BLACK)
	// b.Put(5, 4, board.WHITE)

	showBoardFrame(b)
	termbox.Flush()

	termbox.SetInputMode(termbox.InputEsc)

	runGame(b)
}

var FirstRow = 3
var FirstCol = 5

func showBoardFrame(b *board.PlayBoard) {
	size := b.Size()

	// top label
	for i := 1; i <= size; i++ {
		// nr := strconv.Itoa(i)
		showField(FirstCol+i*4, FirstRow, rune(i+64), termbox.ColorWhite, termbox.ColorBlue)
	}

	// left label
	for i := 1; i <= size; i++ {
		nr := strconv.Itoa(i)
		showField(FirstCol, FirstRow+2+(i-1)*2, rune(nr[0]), termbox.ColorWhite, termbox.ColorBlue)
	}

	showBoard(b)
	showLabels()
	showInput("")
	showAlert("Write field address", 0)
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
				showField(xx, yy, r, fg|termbox.AttrBold, bg)
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

func showAlert(val string, col int) {
	x, y := termbox.Size()

	valLengh := len(val)
	spaces := strings.Repeat(" ", x-valLengh-col-1)

	printText(col, y-1, termbox.ColorWhite|termbox.AttrBold, termbox.ColorWhite, " "+val+spaces)
}

func runGame(b *board.PlayBoard) {
	// value := ""

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:

			switch ev.Key {
			case termbox.KeyEsc:
				break loop

			case termbox.KeyEnter:
				// if len(value) >= 2 {
				showAlert("Pressed enter", 30)
				// }

			case termbox.KeyBackspace:
				showAlert("Pressed backspace", 30)
			}

			showBoard(b)

			// dispatch_press(&ev)
			// pretty_print_press(&ev)
			termbox.Flush()

		case termbox.EventResize:
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			showBoardFrame(b)
			termbox.Flush()
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}

var nextMoveColor = board.BLACK

func showLabels() {
	printText(10, 1, termbox.ColorBlue|termbox.AttrBold, termbox.ColorDefault, "WELCOME IN \"GO\" GAME")
	printText(40, 1, termbox.ColorGreen, termbox.ColorDefault, "(Esc to close game)")
}

func showInput(val string) {
	valLengh := len(val)
	spaces := strings.Repeat(" ", 5-valLengh)

	printText(5, 23, termbox.ColorBlue|termbox.AttrBold, termbox.ColorYellow, val+spaces)
	termbox.SetCursor(5+valLengh, 23)

	if valLengh >= 2 {
		printText(12, 23, termbox.ColorYellow, termbox.ColorDefault, "(enter)")
	} else {
		printText(12, 23, termbox.ColorYellow, termbox.ColorDefault, "      ")
	}

	printText(24, 23, termbox.ColorBlue, termbox.ColorDefault, "move:")

	var userColor string
	if nextMoveColor == board.BLACK {
		userColor = "BLACK"
	} else {
		userColor = "WHITE"
	}

	printText(30, 23, termbox.ColorBlue|termbox.AttrBold, termbox.ColorDefault, userColor)
}

func printText(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

/*
func dispatch_press(ev *termbox.Event) {
	// if ev.Mod & termbox.ModAlt != 0 {
	// 	draw_key(K_LALT, termbox.ColorWhite, termbox.ColorRed);
	// 	draw_key(K_RALT, termbox.ColorWhite, termbox.ColorRed);
	// }

	var k *combo
	if ev.Key >= termbox.KeyArrowRight {
		k = &func_combos[0xFFFF-ev.Key]
	} else if ev.Ch < 128 {
		if ev.Ch == 0 && ev.Key < 128 {
			k = &combos[ev.Key]
		} else {
			k = &combos[ev.Ch]
		}
	}
	if k == nil {
		return
	}

	keys := k.keys
	for _, k := range keys {
		draw_key(k, termbox.ColorWhite, termbox.ColorRed)
	}
}
*/
