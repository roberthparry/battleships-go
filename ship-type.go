package main

// ShipType enum representing the type of the ship
type ShipType int

const (
	// BattleShipType ship is a battle-ship
	BattleShipType ShipType = iota
	// DestroyerType ship is a destroyer
	DestroyerType
)

var shipTypes = [...]string{
	"BattleShip",
	"Destroyer",
}

func (st ShipType) String() string {
	return shipTypes[st]
}
