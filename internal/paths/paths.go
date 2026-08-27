package paths

import (
	"os"
	"path/filepath"
)

type Platform string

const (
	Windows Platform = "windows"
)

func GetAppsPath(platform Platform) string {
	switch platform {
	case Windows:
		return `C:\WinDroid\Apps`

	default:
		return filepath.Join(
			".",
			"WinDroid",
			"Apps",
		)
	}
}

func GetRepositoryPath() string {
	// 1. repository.json рядом с запущенным .exe
	if exe, err := os.Executable(); err == nil {
		path := filepath.Join(
			filepath.Dir(exe),
			"repository.json",
		)

		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	// 2. repository.json в текущей директории
	if path, err := filepath.Abs("repository.json"); err == nil {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	// 3. repository.json в корне проекта WQA
	projectPath := filepath.Join(
		`C:\WQA`,
		"repository.json",
	)

	if _, err := os.Stat(projectPath); err == nil {
		return projectPath
	}

	// Возвращаем стандартный путь,
	// чтобы repository.Load() показал нормальную ошибку.
	return projectPath
}
