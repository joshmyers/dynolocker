package util

import (
	"math/rand"
	"time"
)

func GetRandomTime(lowerBound, upperBound time.Duration) time.Duration {
	if lowerBound < 0 {
		lowerBound = -1 * lowerBound
	}

	if upperBound < 0 {
		upperBound = -1 * upperBound
	}

	if lowerBound > upperBound {
		return upperBound
	}

	if lowerBound == upperBound {
		return lowerBound
	}

	lowerBoundMs := lowerBound.Seconds() * 1000
	upperBoundMs := upperBound.Seconds() * 1000

	lowerBoundMsInt := int(lowerBoundMs)
	upperBoundMsInt := int(upperBoundMs)

	randTimeInt := random(lowerBoundMsInt, upperBoundMsInt)
	return time.Duration(randTimeInt) * time.Millisecond
}

// Generate a random int between min and max, inclusive
func random(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
