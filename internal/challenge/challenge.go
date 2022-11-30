package challenge

import (
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/golangroma/aoc/pkg/aoc"
)

var years = []string{"2022", "2021", "2020"}

func Execute(session string) error {
	// load the teams.yaml file
	teams, err := LoadTeams()
	if err != nil {
		return err
	}

	// get users from teams.yaml file
	users := GetUsersFromTeams(teams)

	// load the AoC leaderboard
	client := aoc.NewClient(session)
	leaderboard, err := client.GetLeaderboard("2022")
	if err != nil {
		return err
	}

	// get scores and stars from the leaderboard
	err = AssignScores(leaderboard, users)
	if err != nil {
		return err
	}

	err = AssignSubmissions(users)
	if err != nil {
		return err
	}

	// sort users by score
	sort.Slice(users, func(i, j int) bool {
		if users[i].Score != users[j].Score {
			return users[i].Score > users[j].Score
		}
		return users[i].Username < users[j].Username
	})

	return UpdateReadme(users, teams.Teams)
}

func AssignScores(leaderboard *aoc.Leaderboard, users []User) error {
	for i, user := range users {

		// find user from leaderboard
		aocID := strconv.Itoa(user.AocID)
		aocMember, found := leaderboard.Members[aocID]
		if !found {
			continue
		}

		// assign the local_score to him
		user.Score = aocMember.LocalScore
		if user.Team != nil {
			user.Team.Score += user.Score
		}

		// check for the stars completions
		for day := 1; day <= 25; day++ {
			dayLevel := aocMember.CompletionDayLevel.GetDayLevel(day)

			if dayLevel.LevelOneCompleted() {
				user.Stars[day-1] = SilverStar
			}
			if dayLevel.LevelTwoCompleted() {
				user.Stars[day-1] = GoldStar
			}
		}

		users[i] = user
	}

	return nil
}

type UserSubmissionsByYear map[string]int

func AssignSubmissions(users []User) error {
	userMap := make(map[string]*User)
	for _, u := range users {
		userMap[u.Username] = &u
	}

	for _, year := range years {
		err := filepath.WalkDir(year, func(path string, d fs.DirEntry, err error) error {
			parts := strings.Split(path, string(os.PathSeparator))
			if len(parts) == 3 && d.IsDir() {
				username := d.Name()

				if u, found := userMap[username]; found {
					if u.Submissions == nil {
						u.Submissions = make(UserSubmissionsByYear)
					}
					u.Submissions[year]++

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
