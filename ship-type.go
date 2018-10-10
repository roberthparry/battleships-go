package main

const (
	//ShipTypeBattleShip ship is a battle-ship
	ShipTypeBattleShip = 0
	//ShipTypeDestroyer ship is a destroyer
	ShipTypeDestroyer = 1
)

//ShipType enum representing the type of the ship
type ShipType int

func (shipType ShipType) String() string {
	var description string
	switch shipType {
	case ShipTypeBattleShip:
		description = "BattleShip"
	case ShipTypeDestroyer:
		description = "Destroyer"
	}
	return description
}
