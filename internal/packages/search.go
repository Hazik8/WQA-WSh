package packages

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Search(query string) error {
	entries, err := os.ReadDir(AppsDir)

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No applications installed")
			return nil
		}

		return err
	}

	query = strings.ToLower(strings.TrimSpace(query))

	fmt.Println("Search results")
	fmt.Println("--------------")

	found := false

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		name := entry.Name()

		if query != "" &&
			!strings.Contains(
				strings.ToLower(name),
				query,
			) {
			continue
		}

		appDir := filepath.Join(
			AppsDir,
			name,
		)

		kind := detectType(appDir)

		fmt.Printf(
			"%-24s %s\n",
			name,
			kind,
		)

		found = true
	}

	if !found {
		fmt.Println("No applications found")
	}

	return nil
}
