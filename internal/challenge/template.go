package challenge

import (
	"os"
	"text/template"
)

func UpdateReadme(users []User, teams []Team) error {
	readmeTemplate, err := template.New("README.md.tmpl").Funcs(template.FuncMap{"add": add}).ParseFiles("assets/README.md.tmpl")
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
