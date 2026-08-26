package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
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
		) {
			return &r.Apps[i]
		}
	}

	return nil
}
