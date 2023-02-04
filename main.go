package main

import (
	"github.com/bmsandoval/flagr/internal/Entry"

	// Import Handlers
	_ "github.com/bmsandoval/flagr/internal/Transports/GrpcHandlers/ReleaseFlagServer"

	"math/rand"
	"time"
)

func main() {
	// Should seed the randomizer once at start of app
	rand.Seed(time.Now().UnixNano())

	Entry.Entry()
}
