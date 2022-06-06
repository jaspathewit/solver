package util

import (
	"log"
	"time"
)

// LogDuration logs how long a method took to execute
func LogDuration(start time.Time, name string) {
	elapsed := time.Since(start)
	if elapsed.Nanoseconds() != 0 {
		log.Printf("%s took %s", name, elapsed)
	}
}
