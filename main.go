package main

import (
	"fmt"
	"mentax.it/TDD/board"
)

func main() {
	b := board.NewBoard()

	fmt.Println("START")

	b.Put(2, 2, board.BLACK)
	b.Put(2, 3, board.WHITE)
	fmt.Println(b)

	fmt.Println("END")
}
