package remove

import (
	"fmt"
	"os"
	"path/filepath"

	"windroid/wqa/internal/paths"
)

func Remove(id string) error {

	appsPath := paths.GetAppsPath(
		paths.Windows,
	)

	appPath := filepath.Join(
		appsPath,
		id,
	)

	_, err := os.Stat(appPath)

	if os.IsNotExist(err) {
		return fmt.Errorf(
			"application not found: %s",
			id,
		)
	}

	err = os.RemoveAll(appPath)

	if err != nil {
		return err
	}

	fmt.Println("[OK] Removed:", id)

	return nil
}
