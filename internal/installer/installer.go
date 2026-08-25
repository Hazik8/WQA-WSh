package installer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"windroid/wqa/internal/format"
	"windroid/wqa/internal/paths"
)

func Install(packagePath string) error {

	data, err := os.ReadFile(packagePath)

	if err != nil {
		return err
	}

	header, err := format.ParseHeader(data)

	if err != nil {
		return err
	}

	_ = header

	manifestStart := header.ManifestOffset
	manifestEnd := manifestStart + header.ManifestSize

	manifestData := data[manifestStart:manifestEnd]

	var manifest format.Manifest

	err = json.Unmarshal(
		manifestData,
		&manifest,
	)

	if err != nil {
		return err
	}

	appsPath := paths.GetAppsPath(
		paths.Windows,
	)

	appPath := filepath.Join(
		appsPath,
		manifest.ID,
	)

	err = os.MkdirAll(
		appPath,
		0755,
	)

	if err != nil {
		return err
	}

	err = os.WriteFile(
		filepath.Join(
			appPath,
			"package.wqa",
		),
		data,
		0644,
	)

	if err != nil {
		return err
	}

	fmt.Println("WQA Installer")
	fmt.Println("----------------")
	fmt.Println("Name:", manifest.Name)
	fmt.Println("ID:", manifest.ID)
	fmt.Println()
	fmt.Println("Installed to:")
	fmt.Println(appPath)
	fmt.Println()
	fmt.Println("Installation complete [OK]")

	return nil
}
