package repository

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

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

func (r *Repository) Search(query string) []App {
	query = strings.ToLower(strings.TrimSpace(query))

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
			results = append(results, app)
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

	if !strings.HasPrefix(app.Download, "https://") {
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

		response, err = client.Get(app.Download)

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

	if _, err := io.Copy(file, response.Body); err != nil {
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

	if _, err := io.Copy(hash, file); err != nil {
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
