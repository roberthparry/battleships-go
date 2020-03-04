package main

import (
	"math/rand"
	"time"
)

var seeded = false

func randomInit() {
	rand.Seed(time.Now().UnixNano())
	seeded = true
}

//RandomDirection returns DirectionDown or DirectionRight
func RandomDirection() Direction {
	if !seeded {
		randomInit()
	}
	return Direction(rand.Intn(2))
}

//RandomPosition returns a random position to start laying out the ship of the given length
func RandomPosition(length int) int {
	if !seeded {
		randomInit()
	}
	return rand.Intn(GameboardGridSize - length + 1)
}
