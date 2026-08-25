package run

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"windroid/wqa/internal/format"
	"windroid/wqa/internal/runtime"
)

const appsDir = `C:\WinDroid\Apps`

type manifest struct {
	WQA          string   `json:"wqa"`
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Version      string   `json:"version"`
	Publisher    string   `json:"publisher"`
	Runtime      string   `json:"runtime"`
	Entry        string   `json:"entry"`
	Architecture []string `json:"architecture"`
}

func Run(target string) error {
	// Путь к .wqa
	if strings.HasSuffix(strings.ToLower(target), ".wqa") ||
		strings.Contains(target, `\`) ||
		strings.Contains(target, `/`) {
		return runFile(target)
	}

	// ID установленного приложения
	return runInstalled(target)
}

func runInstalled(id string) error {
	packagePath := filepath.Join(
		appsDir,
		id,
		"package.wqa",
	)

	if _, err := os.Stat(packagePath); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf(
				"application %q is not installed",
				id,
			)
		}

		return err
	}

	fmt.Println("[INFO] Loading:", id)
	fmt.Println("[INFO] Package:", packagePath)

	return executePackage(packagePath)
}

func runFile(path string) error {
	if _, err := os.Stat(path); err != nil {
		return err
	}

	fmt.Println("[INFO] Loading:", path)

	return executePackage(path)
}

func executePackage(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf(
			"cannot read WQA package: %w",
			err,
		)
	}

	// ------------------------------
	// Header
	// ------------------------------

	header, err := format.ParseHeader(data)

	if err != nil {
		return fmt.Errorf(
			"WQA header error: %w",
			err,
		)
	}

	// ------------------------------
	// Manifest
	// ------------------------------

	manifest, err := readManifest(
		data,
		header.ManifestOffset,
		header.ManifestSize,
	)

	if err != nil {
		return err
	}

	fmt.Println("[INFO] Name:", manifest.Name)
	fmt.Println("[INFO] Version:", manifest.Version)
	fmt.Println("[INFO] Runtime:", manifest.Runtime)

	// Проверяем Runtime.
	if strings.ToLower(manifest.Runtime) != "wqbc" {
		return fmt.Errorf(
			"unsupported runtime: %s",
			manifest.Runtime,
		)
	}

	// ------------------------------
	// WQBC application section
	// ------------------------------

	app, err := readSection(
		data,
		header.AppOffset,
		header.AppSize,
	)

	if err != nil {
		return fmt.Errorf(
			"cannot read application section: %w",
			err,
		)
	}

	fmt.Println("[INFO] Starting WQBC Runtime")
	fmt.Println("[INFO] Bytecode size:", len(app), "bytes")

	// ------------------------------
	// VM
	// ------------------------------

	vm := runtime.New(app)

	vm.Run()

	return nil
}

func readManifest(
	data []byte,
	offset uint32,
	size uint32,
) (*manifest, error) {

	start := uint64(offset)
	end := start + uint64(size)

	if start > uint64(len(data)) ||
		end > uint64(len(data)) ||
		end < start {
		return nil, fmt.Errorf(
			"invalid manifest section",
		)
	}

	raw := data[start:end]

	var m manifest

	if err := json.Unmarshal(raw, &m); err != nil {
		return nil, fmt.Errorf(
			"manifest error: %w",
			err,
		)
	}

	if m.WQA != "1.0" {
		return nil, fmt.Errorf(
			"unsupported WQA version: %s",
			m.WQA,
		)
	}

	if m.ID == "" {
		return nil, fmt.Errorf(
			"manifest has no application ID",
		)
	}

	return &m, nil
}

func readSection(
	data []byte,
	offset uint64,
	size uint64,
) ([]byte, error) {

	end := offset + size

	if offset > uint64(len(data)) ||
		end > uint64(len(data)) ||
		end < offset {
		return nil, fmt.Errorf(
			"invalid section boundaries",
		)
	}

	return data[offset:end], nil
}
