package loader

import (
	"encoding/json"
	"os"

	"windroid/wqa/internal/format"
)

type Package struct {
	Manifest format.Manifest
	App      []byte
}

func Load(path string) (*Package, error) {

	data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	header, err := format.ParseHeader(data)

	if err != nil {
		return nil, err
	}

	manifestData :=
		data[header.ManifestOffset : header.ManifestOffset+header.ManifestSize]

	var manifest format.Manifest

	err = json.Unmarshal(
		manifestData,
		&manifest,
	)

	if err != nil {
		return nil, err
	}

	appStart := header.AppOffset

	appEnd := appStart + header.AppSize

	app := data[appStart:appEnd]

	return &Package{
		Manifest: manifest,
		App:      app,
	}, nil
}
