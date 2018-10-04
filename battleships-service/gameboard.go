package service

import (
	"fmt"
	"strconv"
	"strings"
)

//GameboardGridSize the size of the gameboard
const GameboardGridSize = 10

const columnHeaders = "abcdefghij"

//Gameboard type representing the battleships board
type Gameboard struct {
	cellGrid   [][]Cell
	destroyer1 *Destroyer
	destroyer2 *Destroyer
	battleShip *Battleship
}

//Setup set up the game board
func (board *Gameboard) Setup() {
	board.cellGrid = make([][]Cell, GameboardGridSize)
	for row := 0; row < GameboardGridSize; row++ {
		board.cellGrid[row] = make([]Cell, GameboardGridSize)
		for column := 0; column < GameboardGridSize; column++ {
			board.cellGrid[row][column].Status = CellStatusUnshelled
		}
	}
	board.battleShip = NewBattleship()
	board.destroyer1 = NewDestroyer()
	board.destroyer2 = NewDestroyer()

	RandomInit()

	board.placeShip(&board.battleShip.Ship)
	board.placeShip(&board.destroyer1.Ship)
	board.placeShip(&board.destroyer2.Ship)
}

func (board *Gameboard) placeShip(ship *Ship) {
	var direction Direction
	var position1 int
	var position2 int

	for cannotPlace := true; cannotPlace; cannotPlace = !board.canPlaceShip(ship, position1, position2, direction) {
		direction = RandomDirection()
		position1 = RandomPosition(ship.Length)
		position2 = RandomPosition(1)
	}

	board.doPlaceShip(ship, position1, position2, direction)
}

func (board *Gameboard) canPlaceShip(ship *Ship, position1 int, position2 int, direction Direction) bool {
	if direction == DirectionDown {
		for i := 0; i < ship.Length; i++ {
			if board.cellGrid[position1+i][position2].Ship != nil {
				return false
			}
		}
	} else {
		for i := 0; i < ship.Length; i++ {
			if board.cellGrid[position2][position1+i].Ship != nil {
				return false
			}
		}
	}
	return true
}

func (board *Gameboard) doPlaceShip(ship *Ship, position1 int, position2 int, direction Direction) {
	if direction == DirectionDown {
		for i := 0; i < ship.Length; i++ {
			board.cellGrid[position1+i][position2].Ship = ship
		}
	} else {
		for i := 0; i < ship.Length; i++ {
			board.cellGrid[position2][position1+i].Ship = ship
		}
	}
}

//Print print the battleships game-board
func (board *Gameboard) Print() {
	fmt.Println("     ╔═══╤═══╤═══╤═══╤═══╤═══╤═══╤═══╤═══╤═══╗")
	fmt.Println("     ║ A │ B │ C │ D │ E │ F │ G │ H │ I │ J ║")
	fmt.Println("╔════╬═══╪═══╪═══╪═══╪═══╪═══╪═══╪═══╪═══╪═══╣")
	for row := 0; row < GameboardGridSize; row++ {
		board.printRow(row)
	}
}

func (board *Gameboard) printRow(row int) {
	if row == GameboardGridSize-1 {
		fmt.Print("║ ", row+1, " ║")
	} else {
		fmt.Print("║  ", row+1, " ║")
	}

	for column := 0; column < GameboardGridSize; column++ {
		fmt.Print(" ", board.symbolForCell(&board.cellGrid[row][column]), " ")
		if column == GameboardGridSize-1 {
			fmt.Println("║")
		} else {
			fmt.Print("│")
		}
	}

	if row == GameboardGridSize-1 {
		fmt.Println("╚════╩═══╧═══╧═══╧═══╧═══╧═══╧═══╧═══╧═══╧═══╝")
	} else {
		fmt.Println("╟────╫───┼───┼───┼───┼───┼───┼───┼───┼───┼───╢")
	}
}

func (board *Gameboard) symbolForCell(cell *Cell) string {
	if cell.Status != CellStatusShelled {
		return " "
	}
	if cell.Ship != nil {
		return "#"
	}

	return "x"
}

//IsGameWon returns true if all the ships are destroyed
func (board *Gameboard) IsGameWon() bool {
	return board.battleShip.IsDestroyed() && board.destroyer1.IsDestroyed() && board.destroyer2.IsDestroyed()
}

//FireMissile fire a missile at the given row and column
func (board *Gameboard) FireMissile(row int, column int) FiringResult {
	cell := &board.cellGrid[row][column]
	if cell.Status == CellStatusShelled {
		return FiringResultRepeat
	}

	cell.Status = CellStatusShelled

	if cell.Ship == nil {
		return FiringResultMissed
	}

	cell.Ship.Hits++
	return FiringResultHit
}

//GameboardTranslateCellReference translate A1, B1, etc. to a row and column
func GameboardTranslateCellReference(cellReference string, row *int, column *int) bool {
	if cellReference == "" {
		return false
	}
	cellReference = strings.ToLower(cellReference)
	index := strings.IndexByte(columnHeaders, cellReference[0])
	if index < 0 {
		return false
	}
	*column = index

	number, err := strconv.Atoi(cellReference[1:len(cellReference)])
	if err != nil || number > 10 || number < 1 {
		return false
	}

	*row = number - 1
	return true
}
