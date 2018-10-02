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
func (gameboard *Gameboard) Setup() {
	gameboard.cellGrid = make([][]Cell, GameboardGridSize)
	for row := 0; row < GameboardGridSize; row++ {
		gameboard.cellGrid[row] = make([]Cell, GameboardGridSize)
		for column := 0; column < GameboardGridSize; column++ {
			gameboard.cellGrid[row][column].Status = CellStatusUnshelled
		}
	}
	gameboard.battleShip = NewBattleship()
	gameboard.destroyer1 = NewDestroyer()
	gameboard.destroyer2 = NewDestroyer()

	RandomInit()

	gameboard.placeShip(&gameboard.battleShip.Ship)
	gameboard.placeShip(&gameboard.destroyer1.Ship)
	gameboard.placeShip(&gameboard.destroyer2.Ship)
}

func (gameboard *Gameboard) placeShip(ship *Ship) {
	var direction Direction
	var position1 int
	var position2 int

	for cannotPlace := true; cannotPlace; cannotPlace = !gameboard.canPlaceShip(ship, position1, position2, direction) {
		direction = RandomDirection()
		position1 = RandomPosition(ship.Length)
		position2 = RandomPosition(1)
	}

	gameboard.doPlaceShip(ship, position1, position2, direction)
}

func (gameboard *Gameboard) canPlaceShip(ship *Ship, position1 int, position2 int, direction Direction) bool {
	if direction == DirectionDown {
		for i := 0; i < ship.Length; i++ {
			if gameboard.cellGrid[position1+i][position2].Ship != nil {
				return false
			}
		}
	} else {
		for i := 0; i < ship.Length; i++ {
			if gameboard.cellGrid[position2][position1+i].Ship != nil {
				return false
			}
		}
	}
	return true
}

func (gameboard *Gameboard) doPlaceShip(ship *Ship, position1 int, position2 int, direction Direction) {
	if direction == DirectionDown {
		for i := 0; i < ship.Length; i++ {
			gameboard.cellGrid[position1+i][position2].Ship = ship
		}
	} else {
		for i := 0; i < ship.Length; i++ {
			gameboard.cellGrid[position2][position1+i].Ship = ship
		}
	}
}

//Print print the battleships game-board
func (gameboard *Gameboard) Print() {
	fmt.Println("     ╔═══╤═══╤═══╤═══╤═══╤═══╤═══╤═══╤═══╤═══╗")
	fmt.Println("     ║ A │ B │ C │ D │ E │ F │ G │ H │ I │ J ║")
	fmt.Println("╔════╬═══╪═══╪═══╪═══╪═══╪═══╪═══╪═══╪═══╪═══╣")
	for row := 0; row < GameboardGridSize; row++ {
		gameboard.printRow(row)
	}
}

func (gameboard *Gameboard) printRow(row int) {
	if row == GameboardGridSize-1 {
		fmt.Print("║ ", row+1, " ║")
	} else {
		fmt.Print("║  ", row+1, " ║")
	}

	for column := 0; column < GameboardGridSize; column++ {
		fmt.Print(" ", gameboard.symbolForCell(&gameboard.cellGrid[row][column]), " ")
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

func (gameboard *Gameboard) symbolForCell(cell *Cell) string {
	if cell.Status != CellStatusShelled {
		return " "
	}
	if cell.Ship != nil {
		return "#"
	}

	return "x"
}

//IsGameWon returns true if all the ships are destroyed
func (gameboard *Gameboard) IsGameWon() bool {
	return gameboard.battleShip.IsDestroyed() && gameboard.destroyer1.IsDestroyed() && gameboard.destroyer2.IsDestroyed()
}

//FireMissile fire a missile at the given row and column
func (gameboard *Gameboard) FireMissile(row int, column int) FiringResult {
	cell := &gameboard.cellGrid[row][column]
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
	if err != nil {
		return false
	}

	*row = number - 1
	return true
}
