package utils

import (
	"os"
	"os/exec"
	"runtime"
)

// clear screen
func ClearScreen() {
	// switch based on the OS
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
