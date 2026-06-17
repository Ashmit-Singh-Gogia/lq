package config

import (
	"net/url"
	"os"
	"strings"
)

type Config struct {
	GitignoreListURL string
	GitignoreGetURL  string
	LicenseListURL   string
	LicenseGetURL    string
}

// Load reads configuration from environment variables, falling back to defaults if not set.
func Load() (*Config, error) {
	cfg := &Config{
		GitignoreListURL: getEnv("LQ_GITIGNORE_LIST_URL", "https://www.toptal.com/developers/gitignore/api/list?format=json"),
		GitignoreGetURL:  getEnv("LQ_GITIGNORE_GET_URL", "https://www.toptal.com/developers/gitignore/api/%s"),
		LicenseListURL:   getEnv("LQ_LICENSE_LIST_URL", "https://api.github.com/licenses"),
		LicenseGetURL:    getEnv("LQ_LICENSE_GET_URL", "https://api.github.com/licenses/%s"),
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return fallback
}

// Validate ensures all custom or default URLs are well-formed on startup
func (c *Config) Validate() error {
	urls := []string{c.GitignoreListURL, c.GitignoreGetURL, c.LicenseListURL, c.LicenseGetURL}

	for _, u := range urls {
		testURL := strings.ReplaceAll(u, "%s", "dummy-template")

		if _, err := url.ParseRequestURI(testURL); err != nil {
			return err
		}
	}

	return nil
}
