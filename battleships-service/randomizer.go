package service

import (
	"math/rand"
	"time"
)

//RandomInit initialise randomizer
func RandomInit() {
	rand.Seed(time.Now().UnixNano())
}

//RandomDirection returns DirectionDown or DirectionRight
func RandomDirection() Direction {
	return Direction(rand.Intn(2))
}

//RandomPosition returns a random position to start laying out the ship of the given length
func RandomPosition(length int) int {
	return rand.Intn(GameboardGridSize - length + 1)
}
