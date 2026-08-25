package list

import (
	"fmt"
	"os"

	"windroid/wqa/internal/paths"
)

func Show() error {

	appsPath := paths.GetAppsPath(
		paths.Windows,
	)

	fmt.Println("WQA Installed Applications")
	fmt.Println("--------------------------")

	entries, err := os.ReadDir(appsPath)

	if err != nil {

		if os.IsNotExist(err) {
			fmt.Println("No applications installed")
			return nil
		}

		return err
	}

	if len(entries) == 0 {
		fmt.Println("No applications installed")
		return nil
	}

	for _, entry := range entries {

		if entry.IsDir() {
			fmt.Println("[OK]", entry.Name())
		}
	}

	return nil
}
