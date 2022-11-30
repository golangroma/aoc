package challenge

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

func UpdateReadme(users []User, teams []Team) error {
	utilityFuncs := template.FuncMap{
		"add":        add,
		"mdImage":    mdImage,
		"mdUsername": mdUsername,
		"mdStars":    mdStars,
		"mdTeamID":   mdTeamID,
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
		Users []User
		Teams []Team
	}{
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

func mdStars(u User) string {
	builder := strings.Builder{}

	for _, star := range u.Stars {
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
