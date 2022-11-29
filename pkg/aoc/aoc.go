package aoc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

type Leaderboard struct {
	OwnerdID int                `json:"owner_id"`
	Members  map[string]Members `json:"members"`
}

type Members struct {
	ID                 int                `json:"id"`
	Name               string             `json:"name,omitempty"`
	LastStarTS         int                `json:"last_star_ts"`
	Stars              int                `json:"stars"`
	LocalScore         int                `json:"local_score"`
	GlobalScore        int                `json:"global_score"`
	CompletionDayLevel CompletionDayLevel `json:"completion_day_level"`
}

type CompletionDayLevel map[string]DayLevel

func (cdl CompletionDayLevel) GetDayLevel(day int) DayLevel {
	return cdl[strconv.Itoa(day)]
}

type DayLevel map[string]Level

func (dl DayLevel) LevelOneCompleted() bool {
	_, done := dl["1"]
	return done
}

func (dl DayLevel) LevelTwoCompleted() bool {
	_, done := dl["2"]
	return done
}

type Level struct {
	StarIndex int `json:"star_index"`
	GetStarTS int `json:"get_star_ts"`
}

type Client struct {
	Session string
	Client  *http.Client
}

func NewClient(session string) *Client {
	return &Client{
		Client:  http.DefaultClient,
		Session: session,
	}
}

func (c *Client) GetLeaderboard(year string) (*Leaderboard, error) {
	endpoint := fmt.Sprintf("https://adventofcode.com/%s/leaderboard/private/view/504742.json", year)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating GET request")
	}

	req.Header.Set("cookie", fmt.Sprintf("session=%s", c.Session))
	res, err := c.Client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error doing GET request")
	}
	defer res.Body.Close()

	fmt.Println(res.StatusCode)

	leaderBoard := &Leaderboard{}
	err = json.NewDecoder(res.Body).Decode(leaderBoard)
	if err != nil {
		return nil, errors.Wrap(err, "error decoding JSON")
	}

	return leaderBoard, nil
}
