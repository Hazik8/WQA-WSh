package packages

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"windroid/wqa/internal/format"
)

type AppInfo struct {
	WQA          string   `json:"wqa"`
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Version      string   `json:"version"`
	Publisher    string   `json:"publisher"`
	Runtime      string   `json:"runtime"`
	Entry        string   `json:"entry"`
	Architecture []string `json:"architecture"`
}

func Info(name string) error {
	appDir := filepath.Join(AppsDir, name)

	if _, err := os.Stat(appDir); err != nil {
		return fmt.Errorf("application not installed: %s", name)
	}

	wqaPath := filepath.Join(appDir, "package.wqa")
	exePath := filepath.Join(appDir, "app.exe")

	if _, err := os.Stat(wqaPath); err == nil {
		return infoWQA(wqaPath)
	}

	if _, err := os.Stat(exePath); err == nil {
		infoEXE(name, appDir, exePath)
		return nil
	}

	return fmt.Errorf("unknown application type")
}

func infoWQA(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	header, err := format.ParseHeader(data)

	if err != nil {
		return fmt.Errorf("cannot read WQA header: %w", err)
	}

	manifestStart := int(header.ManifestOffset)
	manifestEnd := manifestStart + int(header.ManifestSize)

	if manifestStart < 0 ||
		manifestEnd > len(data) ||
		manifestStart > manifestEnd {
		return fmt.Errorf("invalid manifest location")
	}

	manifestData := data[manifestStart:manifestEnd]

	var manifest AppInfo

	if err := json.Unmarshal(manifestData, &manifest); err != nil {
		return fmt.Errorf("cannot read manifest: %w", err)
	}

	fmt.Println("Application Information")
	fmt.Println("-----------------------")
	fmt.Println("Name:", manifest.Name)
	fmt.Println("ID:", manifest.ID)
	fmt.Println("Version:", manifest.Version)
	fmt.Println("Publisher:", manifest.Publisher)
	fmt.Println("Runtime:", manifest.Runtime)
	fmt.Println("Entry:", manifest.Entry)

	if len(manifest.Architecture) > 0 {
		fmt.Println("Architecture:", manifest.Architecture)
	}

	fmt.Println("Type: WQA")

	return nil
}

func infoEXE(
	name string,
	appDir string,
	exePath string,
) {
	fmt.Println("Application Information")
	fmt.Println("-----------------------")
	fmt.Println("Name:", name)
	fmt.Println("Type: EXE")
	fmt.Println("Location:", appDir)
	fmt.Println("Executable:", filepath.Base(exePath))
}
