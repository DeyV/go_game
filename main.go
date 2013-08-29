package main

import (
	"fmt"
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

	var game = &OpenGame{NextMove: board.BLACK}

	showBoardFrame(b, game)
	termbox.Flush()

	termbox.SetInputMode(termbox.InputEsc)

	runGame(b, game)
}

/******************************************/
type OpenGame struct {
	NextMove   board.StoneColor
	InputValue string
}

func (g *OpenGame) Move(b *board.PlayBoard) (bool, error) {
	defer func() { g.InputValue = "" }()

	return false, fmt.Errorf("wrong move")
}

func (g *OpenGame) PressKey(r rune) {
	if len(g.InputValue) >= 2 {
		return
	}

	g.InputValue += string(r)
}

func (g *OpenGame) PressBackspace() {
	if g.InputValue != "" {
		g.InputValue = g.InputValue[:len(g.InputValue)-1]
	}
}

func (g *OpenGame) CheckKey(r rune, s int) bool {
	size := rune(s)

	if r >= 'a' && r < 'a'+size {
		return true
	}

	if r >= 'A' && r < 'A'+size {
		return true
	}

	if r >= '1' && (r < '1'+size || r <= '9') {
		return true
	}

	return false
}

/*******************************************/

var FirstRow = 3
var FirstCol = 5

func showBoardFrame(b *board.PlayBoard, game *OpenGame) {
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
	showGameInput(game)
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

func runGame(b *board.PlayBoard, game *OpenGame) {

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:

			switch ev.Key {
			case termbox.KeyEsc:
				break loop

			case termbox.KeyEnter:
				if len(game.InputValue) >= 2 {
					showAlert("Pressed enter", 30)

					_, err := game.Move(b)

					if err != nil {
						showAlert(err.Error(), 50)
					}

					showGameInput(game)

				} else {
					showAlert("To fast, write correct address first", 30)
				}

			case termbox.KeyBackspace:
				showAlert("Pressed backspace", 30)
				game.PressBackspace()
				showGameInput(game)

			default:
				if !game.CheckKey(ev.Ch, b.Size()) {
					showAlert("Wrong key", 30)
				} else {
					game.PressKey(ev.Ch)

					showGameInput(game)
					showBoard(b)
				}
			}

			// dispatch_press(&ev)
			// pretty_print_press(&ev)
			termbox.Flush()

		case termbox.EventResize:
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			showBoardFrame(b, game)
			termbox.Flush()
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}

func showLabels() {
	printText(10, 1, termbox.ColorBlue|termbox.AttrBold, termbox.ColorDefault, "WELCOME IN \"GO\" GAME")
	printText(40, 1, termbox.ColorGreen, termbox.ColorDefault, "(Esc to close game)")
}

func showGameInput(game *OpenGame) {
	showInput(game.InputValue)

	printText(24, 23, termbox.ColorBlue, termbox.ColorDefault, "move:")

	var userColor string

	if game.NextMove == board.BLACK {
		userColor = "BLACK"
	} else {
		userColor = "WHITE"
	}

	printText(30, 23, termbox.ColorBlue|termbox.AttrBold, termbox.ColorDefault, userColor)
}

func showInput(val string) {
	valLengh := len(val)
	inputLenght := 4

	spaces := strings.Repeat(" ", inputLenght-valLengh-1)

	printText(inputLenght, 23, termbox.ColorBlue|termbox.AttrBold, termbox.ColorYellow, " "+val+spaces)
	termbox.SetCursor(inputLenght+valLengh+1, 23)

	if valLengh >= 2 {
		printText(12, 23, termbox.ColorYellow, termbox.ColorDefault, "(enter)")
	} else {
		printText(12, 23, termbox.ColorYellow, termbox.ColorDefault, "       ")
	}

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
