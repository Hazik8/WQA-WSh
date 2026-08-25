package format

import (
	"encoding/json"
	"fmt"
	"os"
)

type Manifest struct {
	WQA          string   `json:"wqa"`
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Version      string   `json:"version"`
	Publisher    string   `json:"publisher"`
	Runtime      string   `json:"runtime"`
	Entry        string   `json:"entry"`
	Architecture []string `json:"architecture"`
}

func ReadManifest(path string) (*Manifest, error) {

	data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var manifest Manifest

	err = json.Unmarshal(
		data,
		&manifest,
	)

	if err != nil {
		return nil, fmt.Errorf(
			"manifest error: %w",
			err,
		)
	}

	if manifest.WQA != "1.0" {
		return nil, fmt.Errorf(
			"unsupported WQA version",
		)
	}

	return &manifest, nil
}
