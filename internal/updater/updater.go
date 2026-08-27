package updater

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"windroid/wqa/internal/installer"
	"windroid/wqa/internal/repository"
)

func Update(name string) error {
	name = strings.TrimSpace(name)

	if name == "" {
		return fmt.Errorf("application name is required")
	}

	// Загружаем репозиторий.
	repo, err := repository.Load("repository.json")
	if err != nil {
		return fmt.Errorf(
			"cannot load repository: %w",
			err,
		)
	}

	// Ищем приложение по имени или ID.
	app := repo.Find(name)

	if app == nil {
		return fmt.Errorf(
			"application not found: %s",
			name,
		)
	}

	fmt.Println("[INFO] Found:", app.Name)
	fmt.Println("[INFO] Version:", app.Version)
	fmt.Println("[INFO] Type:", app.Type)

	if app.Download == "" {
		return fmt.Errorf(
			"application has no download URL",
		)
	}

	// Сейчас обновляем только WQA.
	if !strings.EqualFold(app.Type, "wqa") {
		return fmt.Errorf(
			"unsupported repository package type: %s",
			app.Type,
		)
	}

	// Создаём временную папку.
	tempDir, err := os.MkdirTemp(
		"",
		"wqa-update-*",
	)
	if err != nil {
		return fmt.Errorf(
			"cannot create temporary directory: %w",
			err,
		)
	}

	defer os.RemoveAll(tempDir)

	packagePath := filepath.Join(
		tempDir,
		app.Name+".wqa",
	)

	// Загружаем новую версию.
	fmt.Println("[INFO] Downloading...")

	if err := repository.Download(
		app,
		packagePath,
	); err != nil {
		return err
	}

	// Проверяем SHA-256.
	fmt.Println("[INFO] Verifying SHA-256...")

	if err := repository.VerifySHA256(
		packagePath,
		app.SHA256,
	); err != nil {
		return err
	}

	fmt.Println("[OK] SHA-256 verified")

	// Устанавливаем новую версию.
	fmt.Println("[INFO] Updating...")

	if err := installer.Install(packagePath); err != nil {
		return fmt.Errorf(
			"update installation failed: %w",
			err,
		)
	}

	fmt.Println()
	fmt.Println("[OK] Application updated:", app.Name)
	fmt.Println("[OK] Version:", app.Version)

	return nil
}
