package updater

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"windroid/wqa/internal/format"
	"windroid/wqa/internal/installer"
	"windroid/wqa/internal/paths"
	"windroid/wqa/internal/repository"
)

func Update(name string) error {
	name = strings.TrimSpace(name)

	if name == "" {
		return fmt.Errorf("application name is required")
	}

	repo, err := loadRepository()
	if err != nil {
		return err
	}

	app := repo.Find(name)

	if app == nil {
		return fmt.Errorf(
			"application not found: %s",
			name,
		)
	}

	return updateApp(app)
}

func UpdateAll() error {
	repo, err := loadRepository()
	if err != nil {
		return err
	}

	appsPath := paths.GetAppsPath(paths.Windows)

	entries, err := os.ReadDir(appsPath)

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("[INFO] No applications installed")
			return nil
		}

		return fmt.Errorf(
			"cannot read applications directory: %w",
			err,
		)
	}

	fmt.Println("[INFO] Checking for application updates...")
	fmt.Println()

	checked := 0
	updated := 0

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		appDir := filepath.Join(
			appsPath,
			entry.Name(),
		)

		packagePath := filepath.Join(
			appDir,
			"package.wqa",
		)

		if _, err := os.Stat(packagePath); err != nil {
			continue
		}

		installed, err := readInstalledManifest(
			packagePath,
		)

		if err != nil {
			fmt.Printf(
				"[WARN] Cannot read %s: %v\n",
				entry.Name(),
				err,
			)
			continue
		}

		app := repo.Find(installed.ID)

		if app == nil {
			app = repo.Find(installed.Name)
		}

		if app == nil {
			fmt.Printf(
				"[INFO] %s: not found in repository\n",
				installed.Name,
			)
			continue
		}

		checked++

		if compareVersions(
			app.Version,
			installed.Version,
		) <= 0 {
			fmt.Printf(
				"[OK] %s: %s is up to date\n",
				installed.Name,
				installed.Version,
			)
			continue
		}

		fmt.Printf(
			"[INFO] %s: %s -> %s\n",
			installed.Name,
			installed.Version,
			app.Version,
		)

		if err := updateApp(app); err != nil {
			fmt.Printf(
				"[ERROR] Failed to update %s: %v\n",
				app.Name,
				err,
			)
			continue
		}

		updated++
		fmt.Println()
	}

	fmt.Println()

	if checked == 0 {
		fmt.Println("[INFO] No repository applications to check")
		return nil
	}

	if updated == 0 {
		fmt.Println("[OK] All applications are up to date")
		return nil
	}

	fmt.Printf(
		"[OK] Updated %d application(s)\n",
		updated,
	)

	return nil
}

func loadRepository() (*repository.Repository, error) {
	repositoryPath := paths.GetRepositoryPath()

	repo, err := repository.Load(repositoryPath)

	if err != nil {
		return nil, fmt.Errorf(
			"cannot load repository: %w",
			err,
		)
	}

	return repo, nil
}

func updateApp(app *repository.App) error {
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

type installedManifest struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

func readInstalledManifest(
	packagePath string,
) (*installedManifest, error) {
	data, err := os.ReadFile(packagePath)

	if err != nil {
		return nil, err
	}

	header, err := format.ParseHeader(data)

	if err != nil {
		return nil, fmt.Errorf(
			"invalid WQA header: %w",
			err,
		)
	}

	start := int(header.ManifestOffset)
	end := start + int(header.ManifestSize)

	if start < 0 ||
		end > len(data) ||
		start > end {
		return nil, fmt.Errorf(
			"invalid manifest location",
		)
	}

	var manifest installedManifest

	if err := json.Unmarshal(
		data[start:end],
		&manifest,
	); err != nil {
		return nil, fmt.Errorf(
			"invalid manifest: %w",
			err,
		)
	}

	return &manifest, nil
}

func compareVersions(a, b string) int {
	aParts := parseVersion(a)
	bParts := parseVersion(b)

	max := len(aParts)

	if len(bParts) > max {
		max = len(bParts)
	}

	for i := 0; i < max; i++ {
		var av int
		var bv int

		if i < len(aParts) {
			av = aParts[i]
		}

		if i < len(bParts) {
			bv = bParts[i]
		}

		if av > bv {
			return 1
		}

		if av < bv {
			return -1
		}
	}

	return 0
}

func parseVersion(version string) []int {
	version = strings.TrimSpace(version)
	version = strings.TrimPrefix(version, "v")

	parts := strings.Split(version, ".")

	result := make([]int, len(parts))

	for i, part := range parts {
		n := 0

		for _, char := range part {
			if char < '0' || char > '9' {
				break
			}

			n = n*10 + int(char-'0')
		}

		result[i] = n
	}

	return result
}
