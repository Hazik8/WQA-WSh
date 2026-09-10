package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func historyFilePath() string {
	userProfile := os.Getenv("USERPROFILE")

	if userProfile == "" {
		userProfile, _ = os.UserHomeDir()
	}

	return filepath.Join(
		userProfile,
		".wsh",
		"history",
	)
}

func addHistory(command string) {
	command = strings.TrimSpace(command)

	if command == "" {
		return
	}

	commandHistory = append(commandHistory, command)

	// Keep the last 500 commands in memory.
	if len(commandHistory) > 500 {
		commandHistory = commandHistory[len(commandHistory)-500:]
	}

	saveHistory()
}

func loadHistory() {
	path := historyFilePath()

	data, err := os.ReadFile(path)

	if err != nil {
		return
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line != "" {
			commandHistory = append(commandHistory, line)
		}
	}

	if len(commandHistory) > 500 {
		commandHistory = commandHistory[len(commandHistory)-500:]
	}
}

func saveHistory() {
	path := historyFilePath()

	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return
	}

	data := strings.Join(commandHistory, "\n")

	_ = os.WriteFile(
		path,
		[]byte(data),
		0644,
	)
}

func showHistory(args []string) {
	if len(commandHistory) == 0 {
		fmt.Println("[INFO] History is empty.")
		return
	}

	start := 0

	if len(args) > 0 {
		count, err := strconv.Atoi(args[0])

		if err == nil && count > 0 && count < len(commandHistory) {
			start = len(commandHistory) - count
		}
	}

	for i := start; i < len(commandHistory); i++ {
		fmt.Printf("%4d  %s\n", i+1, commandHistory[i])
	}
}
