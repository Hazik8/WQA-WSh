package repository

type App struct {
	Name     string `json:"name"`
	ID       string `json:"id"`
	Version  string `json:"version"`
	Type     string `json:"type"`
	Download string `json:"download"`
	SHA256   string `json:"sha256"`
}
