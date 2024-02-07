package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

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

type RepoScore struct {
	Name             string
	URL              string
	CriticalityScore float64
	ScoredOn         time.Time
}

type ByCriticalityScore []RepoScore

func (a ByCriticalityScore) Len() int           { return len(a) }
func (a ByCriticalityScore) Less(i, j int) bool { return a[i].CriticalityScore > a[j].CriticalityScore }
func (a ByCriticalityScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	fmt.Println("RANKSCAPE\nRetrieving CNCF landscape projects...")
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

	// Calculate and collect criticality score for each GitHub repository
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		fmt.Println("warning: env variable GITHUB_AUTH_TOKEN not provided")
	}

	var repoScores []RepoScore
	fmt.Println("Calculating scores for CNCF landscape projects...")
	for i, repoURL := range githubRepos {
		// just for testing in order to avoid reaching rete limit on github api
		if i >= 10 {
			break
		}
		fmt.Printf("Project %d of %d\n", i+1, len(githubRepos))
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

		// Parse ScoredOn string to time.Time
		scoredOn, err := time.Parse("Mon Jan 2 15:04:05 MST 2006", score.ScoredOn)
		if err != nil {
			fmt.Printf("Error parsing ScoredOn for repository %s: %s\n", repoURL, err)
			continue
		}

		repoScores = append(repoScores, RepoScore{
			Name:             score.Name,
			URL:              repoURL,
			CriticalityScore: score.CriticalityScore,
			ScoredOn:         scoredOn,
		})
	}

	// Sort repositories by criticality score
	sort.Sort(ByCriticalityScore(repoScores))

	// Generate HTML table
	htmlTable := "<table><tr><th>Rank</th><th>Name</th><th>Score</th><th>Scored On</th></tr>"
	for i, repoScore := range repoScores {
		rank := i + 1
		htmlTable += fmt.Sprintf("<tr><td>%d</td><td><a href=\"%s\">%s</a></td><td>%f</td><td>%s</td></tr>",
			rank, repoScore.URL, repoScore.Name, repoScore.CriticalityScore, repoScore.ScoredOn.Format("Mon Jan _2 15:04:05 2006"))
	}
	htmlTable += "</table>"

	// Update README.md with the generated HTML table
	updateReadme(htmlTable)
	fmt.Println("Rankings updates successfully")
}

func updateReadme(htmlTable string) {
	// Read the existing content of README.md
	readmeFile, err := os.ReadFile("README.md")
	if err != nil {
		log.Fatal("Error reading README.md:", err)
	}

	// Replace the placeholder with the generated HTML table
	readmeContent := strings.Replace(string(readmeFile), "<!--TABLE_PLACEHOLDER-->", htmlTable, 1)

	// Write the updated content back to README.md
	err = os.WriteFile("README.md", []byte(readmeContent), 0644)
	if err != nil {
		log.Fatal("Error writing README.md:", err)
	}
}
