package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/engelsjk/criticalityscore/criticalityscore"
	"gopkg.in/yaml.v2"
)

type Landscape struct {
	Landscape []Category `yaml:"landscape"`
}

type Category struct {
	Name          string        `yaml:"name"`
	Subcategories []Subcategory `yaml:"subcategories"`
}

type Subcategory struct {
	Name  string `yaml:"name"`
	Items []Item `yaml:"items"`
}

type Item struct {
	Name    string `yaml:"name"`
	RepoURL string `yaml:"repo_url"`
}

func main() {
	// Retrieve YAML data from the CNCF landscape
	resp, err := http.Get("https://raw.githubusercontent.com/cncf/landscape/master/landscape.yml")
	if err != nil {
		log.Fatal("Error retrieving landscape YAML:", err)
	}
	defer resp.Body.Close()

	// Read YAML content
	yamlData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading YAML data:", err)
	}

	// Unmarshal YAML into struct
	var landscape Landscape
	if err := yaml.Unmarshal(yamlData, &landscape); err != nil {
		log.Fatal("Error unmarshalling YAML:", err)
	}

	// Extract GitHub repo URLs
	var githubRepos []string
	for _, category := range landscape.Landscape {
		for _, subcategory := range category.Subcategories {
			for _, item := range subcategory.Items {
				if strings.HasPrefix(item.RepoURL, "https://github.com/") {
					githubRepos = append(githubRepos, item.RepoURL)
				}
			}
		}
	}

	// Calculate criticality score for each GitHub repository
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		fmt.Println("warning: env variable GITHUB_AUTH_TOKEN not provided")
	}

	//for _, repoURL := range githubRepos {
	for i, repoURL := range githubRepos {
		if i >= 10 {
			break // Just for test: Exit the loop after the 10th element
		}
		repo, err := criticalityscore.LoadRepository(repoURL, token)
		if err != nil {
			fmt.Printf("Error loading repository %s: %s\n", repoURL, err)
			continue
		}

		score, err := criticalityscore.RepositoryStats(repo, nil) // No additional parameters for now
		if err != nil {
			fmt.Printf("Error calculating score for repository %s: %s\n", repoURL, err)
			continue
		}

		// Print GitHub repo URL along with criticality score
		fmt.Printf("Repository: %s, Criticality Score: %f\n", repoURL, score)
	}
}
