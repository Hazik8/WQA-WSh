package paths

import (
	"os"
	"path/filepath"
)

type Platform string

const (
	Windows Platform = "windows"
	WDOS    Platform = "wdos"
	Linux   Platform = "linux"
)

func GetAppsPath(platform Platform) string {

	switch platform {

	case Windows:
		return "C:\\WinDroid\\Apps"

	case WDOS:
		return "/users/default/apps"

	case Linux:
		home, _ := os.UserHomeDir()

		return filepath.Join(
			home,
			".windroid",
			"apps",
		)

	default:
		return "./apps"
	}
}
