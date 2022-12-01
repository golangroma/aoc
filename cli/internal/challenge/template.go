package challenge

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"text/template"
)

func UpdateReadme(users []*User, teams []*Team) error {
	utilityFuncs := template.FuncMap{
		"add":             add,
		"mdImage":         mdImage,
		"mdUsername":      mdUsername,
		"mdStars":         mdStars,
		"mdTeamID":        mdTeamID,
		"sortByYearScore": sortByYearScore,
	}

	readmeTemplate, err := template.New("README.md.tmpl").Funcs(utilityFuncs).ParseFiles("assets/README.md.tmpl")
	if err != nil {
		return err
	}

	readme, err := os.OpenFile("README.md", os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer readme.Close()

	templateData := struct {
		Years []string
		Users []*User
		Teams []*Team
	}{
		Years: years,
		Users: users,
		Teams: teams,
	}

	return readmeTemplate.Execute(readme, templateData)
}

func add(x, y int) int {
	return x + y
}

func mdImage(url string) string {
	if url == "" {
		return ""
	}
	return fmt.Sprintf("![%s](%s)", url, url)
}

func mdUsername(u User) string {
	if u.Username != "" {
		return fmt.Sprintf("[%s](%s)", u.Username, u.ProfileURL)
	}
	return u.Name
}

func mdTeamID(u User) string {
	if u.Team == nil {
		return ""
	}
	return u.Team.ID
}

func mdStars(userStats Stats) string {
	builder := strings.Builder{}

	for _, star := range userStats.Stars {
		switch star {
		case NoStar:
			builder.WriteString("➖")
		case SilverStar:
			builder.WriteString("⭐")
		case GoldStar:
			builder.WriteString("🌟")
		}
	}

	return builder.String()
}

func sortByYearScore(year string, users []*User) []*User {
	sort.Slice(users, func(i, j int) bool {
		if users[i].Stats[year].Score != users[j].Stats[year].Score {
			return users[i].Stats[year].Score > users[j].Stats[year].Score
		}

		if users[i].Username == "" && users[j].Username == "" {
			return users[i].Name < users[j].Name
		}
		if users[i].Username != "" && users[j].Username != "" {
			return users[i].Username < users[j].Username
		}

		return users[i].Username != ""
	})

	return users
}
