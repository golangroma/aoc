package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)

var years = []string{"2022", "2021", "2020"}

type Teams struct {
	Teams        []Team   `yaml:"team"`
	Participants []string `yaml:"participants"`
}

type Team struct {
	ID          string   `yaml:"id"`
	Members     []string `yaml:"members"`
	Submissions int
}

type User struct {
	Username    string
	Avatar      string
	Team        Team
	ProfileURL  string
	Submissions map[string]int
}

func main() {
	f, err := os.ReadFile("teams.yaml")
	checkErr(err)

	teams := Teams{}
	err = yaml.Unmarshal(f, &teams)
	checkErr(err)

	users := GetUsersFromTeams(teams)
	userYearSubmissionsMap := findSubmissions()
	for i := range users {
		users[i].Submissions = userYearSubmissionsMap[users[i].Username]
	}

	// sort users by submissions
	sort.Slice(users, func(i, j int) bool {
		for _, year := range years {
			if users[i].Submissions[year] != users[j].Submissions[year] {
				return users[i].Submissions[year] > users[j].Submissions[year]
			}
		}
		return users[i].Username < users[j].Username
	})

	readmeTemplate, err := template.New("README.md.tmpl").Funcs(template.FuncMap{"add": add}).ParseFiles("README.md.tmpl")
	checkErr(err)

	readme, err := os.OpenFile("README.md", os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	checkErr(err)
	defer readme.Close()

	templateData := struct {
		Users []User
		Teams []Team
	}{
		Users: users,
		Teams: teams.Teams,
	}

	err = readmeTemplate.Execute(readme, templateData)
	checkErr(err)
}

func GetUsersFromTeams(teams Teams) []User {
	participantsMap := make(map[string]User)

	for _, p := range teams.Participants {
		participantsMap[p] = User{
			Username:   p,
			Avatar:     fmt.Sprintf("https://github.com/%s.png?size=60", p),
			Team:       Team{},
			ProfileURL: fmt.Sprintf("https://github.com/%s", p),
		}
	}

	for _, t := range teams.Teams {
		for _, member := range t.Members {
			if m, found := participantsMap[member]; found {
				m.Team = t
				participantsMap[member] = m
			}
		}
	}

	users := []User{}
	for _, v := range participantsMap {
		users = append(users, v)
	}

	return users
}

func findSubmissions() map[string]map[string]int {
	userYearSubmissionsMap := map[string]map[string]int{}

	for _, year := range years {
		err := filepath.WalkDir(year, func(path string, d fs.DirEntry, err error) error {
			parts := strings.Split(path, string(os.PathSeparator))
			if len(parts) == 3 && d.IsDir() {
				user := d.Name()

				if userYearSubmissionsMap[user] == nil {
					userYearSubmissionsMap[user] = make(map[string]int)
				}
				userYearSubmissionsMap[user][year]++
			}

			return nil
		})
		checkErr(err)
	}

	return userYearSubmissionsMap
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func add(x, y int) int {
	return x + y
}
