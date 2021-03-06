package main

import (
	"fmt"
	"strings"
)

func main() {
	board := &Gameboard{}
	board.Setup()

	for gameNotFinished := true; gameNotFinished; gameNotFinished = !board.IsGameWon() {
		board.Print()
		fmt.Print("Enter cell to fire at (e.g. A1, B1, ...) or q to quit: ")
		var cellReference string
		fmt.Scanln(&cellReference)
		if strings.ToLower(cellReference) == "q" {
			fmt.Println("Bye!")
			return
		}
		var row, column int
		if !GameboardTranslateCellReference(cellReference, &row, &column) {
			fmt.Println("'", cellReference, "'", " is not a valid cell.")
			continue
		}

		result := board.FireMissile(row, column)
		result.PrintResponse()
	}

	board.Print()
	fmt.Println("Congratulations, you Won!")
}
