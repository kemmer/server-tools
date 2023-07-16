package functions

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	_  = 0
	KB = 1 << (10 * iota)
	MB
	GB
)

func HelloWorld() string {
	randomNumber := rand.Int() % 10000001

	return fmt.Sprintf("Hello, world! 👋 (ID: %d)", randomNumber)
}

func StressTest(memorySizeGB int, timeSeconds int) string {
	buffer := make([]uint8, memorySizeGB*GB, memorySizeGB*GB)

	// fills up the buffer with data
	c := 0
	for be, _ := range buffer {
		buffer[be] = math.MaxUint8 - uint8(c)
		c++
	}

	d, err := time.ParseDuration(fmt.Sprintf("%ds", timeSeconds))
	if err != nil {
		return fmt.Sprint("Stress test failed to parse duration")
	}
	time.Sleep(d)

	return fmt.Sprintf("Allocated %d GB for %.0f seconds", memorySizeGB, d.Seconds())
}
