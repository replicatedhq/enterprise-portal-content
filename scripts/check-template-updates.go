package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

type updateFrontmatter struct {
	ID          string   `yaml:"id"`
	PublishedAt string   `yaml:"published_at"`
	Title       string   `yaml:"title"`
	Impact      string   `yaml:"impact"`
	Summary     string   `yaml:"summary"`
	Affects     []string `yaml:"affects"`
}

func main() {
	updatesDir := os.Getenv("UPDATES_DIR")
	if updatesDir == "" {
		updatesDir = filepath.Join("..", "updates")
	}

	entries, err := os.ReadDir(updatesDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read %s: %v\n", updatesDir, err)
		os.Exit(1)
	}

	var failed int
	var checked int
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}
		checked++
		path := filepath.Join(updatesDir, entry.Name())
		if err := checkUpdateFile(path); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", path, err)
			failed++
		}
	}

	if checked == 0 {
		fmt.Fprintf(os.Stderr, "no updates/*.md files in %s\n", updatesDir)
		os.Exit(1)
	}
	if failed > 0 {
		os.Exit(1)
	}
	fmt.Printf("ok: parsed %d template update files\n", checked)
}

func checkUpdateFile(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	frontmatter, err := parseFrontmatter(string(content))
	if err != nil {
		return err
	}

	var parsed updateFrontmatter
	if err := yaml.Unmarshal([]byte(frontmatter), &parsed); err != nil {
		return fmt.Errorf("failed to parse update frontmatter: %w", err)
	}
	if parsed.ID == "" {
		return fmt.Errorf("update frontmatter missing id")
	}
	if parsed.Title == "" {
		return fmt.Errorf("update frontmatter missing title")
	}
	if parsed.PublishedAt == "" {
		return fmt.Errorf("update frontmatter missing published_at")
	}
	if parsed.Impact == "" {
		return fmt.Errorf("update frontmatter missing impact")
	}
	if parsed.Summary == "" {
		return fmt.Errorf("update frontmatter missing summary")
	}
	if _, err := time.Parse(time.RFC3339, parsed.PublishedAt); err != nil {
		return fmt.Errorf("failed to parse update published_at: %w", err)
	}
	return nil
}

func parseFrontmatter(content string) (string, error) {
	content = strings.TrimSpace(content)
	if !strings.HasPrefix(content, "---") {
		return "", fmt.Errorf("missing update frontmatter")
	}
	rest := content[3:]
	idx := strings.Index(rest, "\n---")
	if idx < 0 {
		return "", fmt.Errorf("missing update frontmatter")
	}
	frontmatter := strings.TrimSpace(rest[:idx])
	if frontmatter == "" {
		return "", fmt.Errorf("missing update frontmatter")
	}
	return frontmatter, nil
}
