package main

import (
	"fmt"
	"strings"

	"github.com/roberthparry/battleships-go/battleships-service"
)

func main() {
	board := &service.Gameboard{}
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
		if !service.GameboardTranslateCellReference(cellReference, &row, &column) {
			fmt.Println("'", cellReference, "'", " is not a valid cell.")
			continue
		}

		result := board.FireMissile(row, column)
		switch result {
		case service.FiringResultHit:
			fmt.Println("Hit!")
		case service.FiringResultMissed:
			fmt.Println("Missed!")
		default:
			fmt.Println("You've already been there!")
		}
	}

	board.Print()
	fmt.Println("Congratulations, you Won!")
}
