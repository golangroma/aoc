package challenge

import "fmt"

type User struct {
	AocID       int
	Name        string
	Username    string
	AvatarURL   string
	Team        *Team
	ProfileURL  string
	Submissions map[string]int
	Score       int
	Stars       [25]Star
}

type Star int

const (
	NoStar Star = iota
	SilverStar
	GoldStar
)

func GetUsersFromTeams(teams *Teams) []User {
	participantsMap := make(map[string]User)

	// load users from the participants
	for _, p := range teams.Participants {
		participantsMap[p.Name] = User{
			AocID:      p.ID,
			Username:   p.Name,
			AvatarURL:  fmt.Sprintf("https://github.com/%s.png?size=60", p.Name),
			ProfileURL: fmt.Sprintf("https://github.com/%s", p.Name),
		}
	}

	// check if a user is part of a team
	for _, t := range teams.Teams {
		for _, member := range t.Members {
			if m, found := participantsMap[member]; found {
				m.Team = &t
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
