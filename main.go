package main

import (
	"github.com/bmsandoval/flagr/internal/Entry"

	// Import Handlers

	"math/rand"
	"time"
)

func main() {
	// Should seed the randomizer once at start of app
	rand.Seed(time.Now().UnixNano())

	Entry.Entry()
}
