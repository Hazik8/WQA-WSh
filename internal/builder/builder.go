package builder

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"windroid/wqa/internal/compiler"
	"windroid/wqa/internal/format"
	"windroid/wqa/internal/lexer"
	"windroid/wqa/internal/logger"
	"windroid/wqa/internal/parser"
)

func Build(project string) error {

	manifestPath := filepath.Join(
		project,
		"wqa.json",
	)

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

	source, err := os.ReadFile(appPath)
	if err != nil {
		return err
	}

	// WQA source header
	sourceText := string(source)

	if strings.TrimSpace(sourceText) == "" {
		return fmt.Errorf("WQA source cannot be empty")
	}

	lines := strings.Split(sourceText, "\n")

	firstLine := strings.TrimSpace(lines[0])

	if firstLine != "wqa" {
		return fmt.Errorf("WQA source must start with 'wqa'")
	}

	// Lexer
	lex := lexer.New(sourceText)

	tokens, err := lex.Tokenize()
	if err != nil {
		return err
	}

	// Parser
	p := parser.New(tokens)

	program, err := p.Parse()
	if err != nil {
		return err
	}
	// Compiler
	appData, err := compiler.Compile(program)
	if err != nil {
		return err
	}

	output := manifest.Name + ".wqa"

	file, err := os.Create(output)
	if err != nil {
		return err
	}

	defer file.Close()

	// ------------------------------
	// Header
	// ------------------------------

	header := format.NewHeader()

	header.ManifestOffset = 64
	header.ManifestSize = uint32(len(manifestData))

	header.AppOffset = uint64(
		header.ManifestOffset +
			uint32(len(manifestData)),
	)

	header.AppSize = uint64(len(appData))

	headerBytes, err := header.MarshalBinary()
	if err != nil {
		return err
	}

	// ------------------------------
	// Header
	// ------------------------------

	if _, err := file.Write(headerBytes); err != nil {
		return err
	}

	// ------------------------------
	// Manifest
	// ------------------------------

	if _, err := file.Write(manifestData); err != nil {
		return err
	}

	// ------------------------------
	// WQBC
	// ------------------------------

	if _, err := file.Write(appData); err != nil {
		return err
	}

	logger.Success(
		"Created: " + output,
	)

	logger.Info(
		fmt.Sprintf(
			"Manifest size: %d bytes",
			len(manifestData),
		),
	)

	logger.Info(
		fmt.Sprintf(
			"Bytecode size: %d bytes",
			len(appData),
		),
	)

	return nil
}
