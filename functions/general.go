package functions

import (
	"fmt"
	"math/rand"
)

func HelloWorld() string {
	randomNumber := rand.Int() % 10000001

	return fmt.Sprintf("Hello, world! 👋 (ID: %d)", randomNumber)
}
