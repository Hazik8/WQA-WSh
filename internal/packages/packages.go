package packages

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const AppsDir = `C:\WinDroid\Apps`

func Install(packageFile string) error {
	if _, err := os.Stat(packageFile); err != nil {
		return fmt.Errorf("package not found: %w", err)
	}

	ext := strings.ToLower(filepath.Ext(packageFile))

	switch ext {
	case ".wqa":
		return installWQA(packageFile)

	case ".exe":
		return installEXE(packageFile)

	default:
		return fmt.Errorf("unsupported package type: %s", ext)
	}
}

func installWQA(packageFile string) error {
	name := strings.TrimSuffix(
		filepath.Base(packageFile),
		filepath.Ext(packageFile),
	)

	target := filepath.Join(AppsDir, name)

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	destination := filepath.Join(target, "package.wqa")

	data, err := os.ReadFile(packageFile)
	if err != nil {
		return err
	}

	if err := os.WriteFile(destination, data, 0644); err != nil {
		return err
	}

	fmt.Println("[OK] Installed WQA:", name)

	return nil
}

func installEXE(exeFile string) error {
	name := strings.TrimSuffix(
		filepath.Base(exeFile),
		filepath.Ext(exeFile),
	)

	target := filepath.Join(AppsDir, name)

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	destination := filepath.Join(target, "app.exe")

	data, err := os.ReadFile(exeFile)
	if err != nil {
		return err
	}

	if err := os.WriteFile(destination, data, 0755); err != nil {
		return err
	}

	fmt.Println("[OK] Installed EXE:", name)
	fmt.Println("Location:", target)

	return nil
}

func Remove(name string) error {
	target := filepath.Join(AppsDir, name)

	if _, err := os.Stat(target); os.IsNotExist(err) {
		return fmt.Errorf("application not installed: %s", name)
	}

	if err := os.RemoveAll(target); err != nil {
		return err
	}

	fmt.Println("[OK] Removed:", name)

	return nil
}

func List() error {
	entries, err := os.ReadDir(AppsDir)

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No applications installed")
			return nil
		}

		return err
	}

	fmt.Println("Installed Applications")
	fmt.Println("----------------------")

	found := false

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		appDir := filepath.Join(AppsDir, entry.Name())

		kind := detectType(appDir)

		fmt.Printf("%-24s %s\n", entry.Name(), kind)

		found = true
	}

	if !found {
		fmt.Println("No applications installed")
	}

	return nil
}

func detectType(appDir string) string {
	wqaPath := filepath.Join(appDir, "package.wqa")
	exePath := filepath.Join(appDir, "app.exe")

	if _, err := os.Stat(wqaPath); err == nil {
		return "WQA"
	}

	if _, err := os.Stat(exePath); err == nil {
		return "EXE"
	}

	return "UNKNOWN"
}
