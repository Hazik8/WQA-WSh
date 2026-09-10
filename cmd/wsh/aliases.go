package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func aliasesFilePath() string {
	userProfile := os.Getenv("USERPROFILE")

	if userProfile == "" {
		userProfile, _ = os.UserHomeDir()
	}

	return filepath.Join(
		userProfile,
		".wsh",
		"aliases",
	)
}

func handleAlias(args []string) {
	if len(args) == 0 {
		showAliases()
		return
	}

	name := strings.ToLower(args[0])

	if len(args) == 1 {
		value, exists := aliases[name]

		if !exists {
			fmt.Println("[INFO] Alias not found:", name)
			return
		}

		fmt.Printf("%s = %s\n", name, value)
		return
	}

	value := strings.Join(args[1:], " ")

	aliases[name] = value
	saveAliases()

	fmt.Printf("[OK] Alias created: %s = %s\n", name, value)
}

func handleUnalias(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: unalias <name>")
		return
	}

	name := strings.ToLower(args[0])

	if _, exists := aliases[name]; !exists {
		fmt.Println("[INFO] Alias not found:", name)
		return
	}

	delete(aliases, name)
	saveAliases()

	fmt.Println("[OK] Alias removed:", name)
}

func showAliases() {
	if len(aliases) == 0 {
		fmt.Println("[INFO] No aliases configured.")
		return
	}

	fmt.Println("Aliases:")

	for name, value := range aliases {
		fmt.Printf("  %-12s = %s\n", name, value)
	}
}

func loadAliases() {
	path := aliasesFilePath()

	data, err := os.ReadFile(path)

	if err != nil {
		return
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		parts := strings.SplitN(line, "=", 2)

		if len(parts) != 2 {
			continue
		}

		name := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		if name == "" || value == "" {
			continue
		}

		aliases[strings.ToLower(name)] = value
	}
}

func saveAliases() {
	path := aliasesFilePath()

	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return
	}

	var builder strings.Builder

	for name, value := range aliases {
		builder.WriteString(name)
		builder.WriteString("=")
		builder.WriteString(value)
		builder.WriteString("\n")
	}

	_ = os.WriteFile(
		path,
		[]byte(builder.String()),
		0644,
	)
}
