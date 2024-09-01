package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func RunDiscCleanup() {
	// TODO: Implement Disk Cleanup functionality
}

func RunUpdateAllPrograms() {
	clearScreen()

	if !IsWingetInstalled() {
		fmt.Println("Winget is not installed! Packages can't be updated.")
		reset()
		return
	}

	fmt.Println("Updating all packages...")
	if output, err := runCommand("winget", "upgrade", "--all"); err != nil {
		fmt.Println("Error updating packages:", err)
		fmt.Println("Output:", output)
	} else {
		fmt.Println("Successfully updated all packages!")
	}

	reset()
}

func IsWingetInstalled() bool {
	output, err := runCommand("winget", "--version")
	if err != nil {
		return false
	}
	return strings.Contains(strings.TrimSpace(output), "v")
}

func reset() {
	time.Sleep(2 * time.Second)
	clearScreen()
	app.Stop()
	main()
}

func runCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}
