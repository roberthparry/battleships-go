package main

//Cell target cell of gameboard grid
type Cell struct {
	Ship   *Ship
	Status CellStatus
}
