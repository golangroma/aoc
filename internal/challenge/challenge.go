package challenge

import (
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/golangroma/aoc/pkg/aoc"
)

var years = []string{"2023", "2022", "2021", "2020"}

func Execute(session string) error {
	// load the teams.yaml file
	teams, err := LoadTeams()
	if err != nil {
		return err
	}

	// get users from teams.yaml file
	users := GetUsersFromTeams(teams)

	// load the AoC leaderboards by year
	client := aoc.NewClient(session)

	leaderboardsByYear := make(map[string]*aoc.Leaderboard)
	for _, year := range years {
		leaderboard, err := client.GetLeaderboard(year)
		if err != nil {
			return err
		}
		leaderboardsByYear[year] = leaderboard
	}

	// add the users missing in the teams.yaml but in the leaderboard (pick the first one, users are the same)
	users = MergeUsers(leaderboardsByYear[years[0]], users)

	// get scores and stars from the leaderboard
	err = AssignScores(leaderboardsByYear, users)
	if err != nil {
		return err
	}

	err = AssignSubmissions(users)
	if err != nil {
		return err
	}

	return UpdateReadme(users, teams.Teams)
}

// MergeUsers will add the users in the leaderboard that are missing in the teams.yaml
func MergeUsers(leaderboard *aoc.Leaderboard, users []*User) []*User {
	teamsUsersMap := make(map[string]*User)
	for _, user := range users {
		teamsUsersMap[strconv.Itoa(user.AocID)] = user
	}

	for aocID, member := range leaderboard.Members {
		if _, found := teamsUsersMap[aocID]; !found { // if not found
			users = append(users, &User{
				AocID: member.ID,
				Name:  member.Name,
			})
		}
	}

	return users
}

func AssignScores(leaderboards map[string]*aoc.Leaderboard, users []*User) error {
	for _, year := range years {
		leaderboard, found := leaderboards[year]
		if !found {
			continue
		}

		for i, user := range users {

			// find user from leaderboard
			aocID := strconv.Itoa(user.AocID)
			aocMember, found := leaderboard.Members[aocID]
			if !found {
				continue
			}

			// assign the local_score to him
			userStats := Stats{}

			userStats.Score = aocMember.LocalScore
			if user.Team != nil {
				if user.Team.Stats == nil {
					user.Team.Stats = make(map[string]Stats)
				}
				teamStats := user.Team.Stats[year]
				teamStats.Score += userStats.Score
				user.Team.Stats[year] = teamStats
			}

			// check for the stars completions
			for day := 1; day <= 25; day++ {
				dayLevel := aocMember.CompletionDayLevel.GetDayLevel(day)

				if dayLevel.LevelOneCompleted() {
					userStats.Stars[day-1] = SilverStar
				}
				if dayLevel.LevelTwoCompleted() {
					userStats.Stars[day-1] = GoldStar
				}
			}

			if user.Stats == nil {
				user.Stats = make(map[string]Stats)
			}
			user.Stats[year] = userStats

			users[i] = user
		}
	}

	return nil
}

type UserSubmissionsByYear map[string]int

func AssignSubmissions(users []*User) error {
	userMap := make(map[string]*User)
	for _, u := range users {
		userMap[u.Username] = u
	}

	for _, year := range years {
		err := filepath.WalkDir(year, func(path string, d fs.DirEntry, err error) error {
			parts := strings.Split(path, string(os.PathSeparator))
			if len(parts) == 3 && d.IsDir() {
				username := d.Name()

				if u, found := userMap[username]; found {

					userStats := u.Stats[year]
					userStats.Submissions++
					u.Stats[year] = userStats

					if u.Team != nil {
						teamStats := u.Team.Stats[year]
						teamStats.Submissions++
						u.Team.Stats[year] = teamStats
					}

					userMap[username] = u
				}
			}

			return nil
		})

		if err != nil {
			return err
		}
	}

	return nil
}
