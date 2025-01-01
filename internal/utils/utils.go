package utils

import "fmt"

// clear screen
func ClearScreen() {
	for i := 0; i < 100; i++ {
		fmt.Println()
	}
}
