package builder

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"windroid/wqa/internal/compiler"
	"windroid/wqa/internal/format"
	"windroid/wqa/internal/logger"
)

func Build(project string) error {

	manifestPath := filepath.Join(project, "wqa.json")

	manifest, err := format.ReadManifest(manifestPath)
	if err != nil {
		return err
	}

	manifestData, err := json.Marshal(manifest)
	if err != nil {
		return err
	}

	appPath := filepath.Join(
		project,
		manifest.Entry,
	)

	appData, err := compiler.Compile(appPath)
	if err != nil {
		return err
	}

	output := manifest.Name + ".wqa"

	file, err := os.Create(output)
	if err != nil {
		return err
	}

	defer file.Close()

	header := format.NewHeader()

	header.ManifestOffset = 64
	header.ManifestSize = uint32(len(manifestData))

	header.AppOffset =
		uint64(header.ManifestOffset) +
			uint64(header.ManifestSize)

	header.AppSize =
		uint64(len(appData))

	headerBytes, err := header.MarshalBinary()

	if err != nil {
		return err
	}

	// Header
	file.Write(headerBytes)

	// Manifest
	file.Write(manifestData)

	// Application
	file.Write(appData)

	logger.Success("Created: " + output)

	logger.Info(
		fmt.Sprintf(
			"Manifest size: %d bytes",
			len(manifestData),
		),
	)

	fmt.Println("App:",
		len(appData),
		"bytes")

	return nil
}
