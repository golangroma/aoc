package main

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

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
	Submissions int
}

func main() {
	f, err := os.ReadFile("teams.yaml")
	checkErr(err)

	teams := Teams{}
	err = yaml.Unmarshal(f, &teams)
	checkErr(err)

	users := GetUsersFromTeams(teams)

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
			Username:    p,
			Avatar:      fmt.Sprintf("https://github.com/%s.png?size=40", p),
			Team:        Team{},
			ProfileURL:  fmt.Sprintf("https://github.com/%s", p),
			Submissions: 0,
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

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func add(x, y int) int {
	return x + y
}
