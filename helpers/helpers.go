package helpers

import (
	"math/rand"
	"time"
)

// Generate a pseudo-random number consisting of as many numbers as specified (default is 32)
func GenerateRandomNumber(maxNumbers int) int {
	if maxNumbers == 0 {
		maxNumbers = 32
	}
	// Seed the random number generator with unixtime nanoseconds
	rand.Seed(time.Now().Unix())

	var generatedNumber int

	for i := 0; i < maxNumbers; i++ {

		generatedNumber += rand.Intn(9)
	}

	return generatedNumber
}
