package helpers

import (
	log "github.com/Sirupsen/logrus"
	"math/rand"
	"strconv"
	"time"
)

// Generate a pseudo-random number consisting of as many numbers as passed into the function (default is 32)
func GenerateRandomNumber(maxNumbers int) int {
	if maxNumbers == 0 {
		maxNumbers = 32
	}
	// Seed the random number generator with unixtime nanoseconds
	rand.Seed(time.Now().Unix())

	var generatedNumber string

	for i := 0; i < maxNumbers; i++ {

		generatedNumber += strconv.Itoa(rand.Intn(10))
	}
	// Convert string back to number
	number, err := strconv.Atoi(generatedNumber)
	if err != nil {
		// Newbie conversion, but should be sufficient for our purposes
		return rand.Intn(100000)
	}

	return number
}

// Helper function to catch and log errors in their appropriate level through the Logrus package
func CheckAndLogError(error interface{}, level string) {
	if error != nil {
		switch level {
		case "warning":
			log.Warning(error)
		case "error":
			log.Error(error)
		case "fatal":
			log.Fatal(error)
		case "panic":
			log.Panic(error)
		}
	}
}
