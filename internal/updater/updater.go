package updater

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"windroid/wqa/internal/installer"
	"windroid/wqa/internal/paths"
	"windroid/wqa/internal/repository"
)

func Update(name string) error {
	name = strings.TrimSpace(name)

	if name == "" {
		return fmt.Errorf("application name is required")
	}

	repositoryPath := paths.GetRepositoryPath()

	repo, err := repository.Load(repositoryPath)

	if err != nil {
		return fmt.Errorf(
			"cannot load repository: %w",
			err,
		)
	}

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

	if !strings.EqualFold(app.Type, "wqa") {
		return fmt.Errorf(
			"unsupported repository package type: %s",
			app.Type,
		)
	}

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

	fmt.Println("[INFO] Downloading...")

	if err := repository.Download(
		app,
		packagePath,
	); err != nil {
		return err
	}

	fmt.Println("[INFO] Verifying SHA-256...")

	if err := repository.VerifySHA256(
		packagePath,
		app.SHA256,
	); err != nil {
		return err
	}

	fmt.Println("[OK] SHA-256 verified")

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
