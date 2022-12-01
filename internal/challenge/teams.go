package challenge

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Participant struct {
	ID   int    `yaml:"aoc_id"`
	Name string `yaml:"name"`
}

type Teams struct {
	Teams        []*Team       `yaml:"team"`
	Participants []Participant `yaml:"participants"`
}

type Team struct {
	ID      string   `yaml:"id"`
	Members []string `yaml:"members"`
	Stats   map[string]Stats
}

func LoadTeams() (*Teams, error) {
	f, err := os.ReadFile("teams.yaml")
	if err != nil {
		return nil, err
	}

	teams := &Teams{}

	err = yaml.Unmarshal(f, teams)
	if err != nil {
		return nil, err
	}

	return teams, nil
}
