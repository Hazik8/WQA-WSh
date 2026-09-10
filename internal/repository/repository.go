package repository

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const DefaultRepositoryURL = "https://raw.githubusercontent.com/Hazik8/WinDroidApps/main/repository.json"

type Repository struct {
	Apps []App `json:"apps"`
}

func Load(path string) (*Repository, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var repo Repository

	if err := json.Unmarshal(data, &repo); err != nil {
		return nil, fmt.Errorf(
			"repository error: %w",
			err,
		)
	}

	return &repo, nil
}

func LoadURL(url string) (*Repository, error) {
	if !strings.HasPrefix(url, "https://") {
		return nil, fmt.Errorf(
			"repository URL must use HTTPS",
		)
	}

	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	response, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to download repository: %w",
			err,
		)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"failed to download repository: HTTP %d",
			response.StatusCode,
		)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to read repository: %w",
			err,
		)
	}

	var repo Repository

	if err := json.Unmarshal(data, &repo); err != nil {
		return nil, fmt.Errorf(
			"repository error: %w",
			err,
		)
	}

	return &repo, nil
}

func Save(path string, repo *Repository) error {
	if repo == nil {
		return fmt.Errorf("repository is nil")
	}

	data, err := json.MarshalIndent(
		repo,
		"",
		"  ",
	)
	if err != nil {
		return fmt.Errorf(
			"failed to encode repository: %w",
			err,
		)
	}

	if err := os.MkdirAll(
		filepath.Dir(path),
		0755,
	); err != nil {
		return fmt.Errorf(
			"failed to create cache directory: %w",
			err,
		)
	}

	if err := os.WriteFile(
		path,
		data,
		0644,
	); err != nil {
		return fmt.Errorf(
			"failed to save repository cache: %w",
			err,
		)
	}

	return nil
}

func (r *Repository) Search(query string) []App {
	query = strings.ToLower(
		strings.TrimSpace(query),
	)

	if query == "" {
		return r.Apps
	}

	var results []App

	for _, app := range r.Apps {
		if strings.Contains(
			strings.ToLower(app.Name),
			query,
		) || strings.Contains(
			strings.ToLower(app.ID),
			query,
		) {
			results = append(
				results,
				app,
			)
		}
	}

	return results
}

func (r *Repository) Find(id string) *App {
	for i := range r.Apps {
		if strings.EqualFold(
			r.Apps[i].ID,
			id,
		) || strings.EqualFold(
			r.Apps[i].Name,
			id,
		) {
			return &r.Apps[i]
		}
	}

	return nil
}

func Download(app *App, destination string) error {
	if app.Download == "" {
		return fmt.Errorf(
			"application has no download URL",
		)
	}

	if !strings.HasPrefix(
		app.Download,
		"https://",
	) {
		return fmt.Errorf(
			"download URL must use HTTPS",
		)
	}

	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	var response *http.Response
	var err error

	for attempt := 1; attempt <= 3; attempt++ {
		fmt.Printf(
			"[INFO] Download attempt %d/3...\n",
			attempt,
		)

		response, err = client.Get(
			app.Download,
		)

		if err == nil {
			break
		}

		if attempt < 3 {
			fmt.Println("[INFO] Retrying...")
			time.Sleep(2 * time.Second)
		}
	}

	if err != nil {
		return fmt.Errorf(
			"download failed: %w",
			err,
		)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf(
			"download failed: HTTP %d",
			response.StatusCode,
		)
	}

	file, err := os.Create(destination)
	if err != nil {
		return err
	}

	defer file.Close()

	if _, err := io.Copy(
		file,
		response.Body,
	); err != nil {
		return err
	}

	return nil
}

func VerifySHA256(path string, expected string) error {
	expected = strings.ToLower(
		strings.TrimSpace(expected),
	)

	if expected == "" {
		return fmt.Errorf(
			"SHA-256 is missing",
		)
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()

	hash := sha256.New()

	if _, err := io.Copy(
		hash,
		file,
	); err != nil {
		return err
	}

	actual := hex.EncodeToString(
		hash.Sum(nil),
	)

	if actual != expected {
		return fmt.Errorf(
			"SHA-256 mismatch: expected %s, got %s",
			expected,
			actual,
		)
	}

	return nil
}
