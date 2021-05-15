package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Team struct {
	name   string
	points int
	rank   int
}

const (
	pointsDraw = 1
	pointsWin  = 3
	pointsLose = 0
)

// GetRankedTeams take league game scores as a string and outputs a slice of teams ordered by their rank
func GetRankedTeams(input string) ([]Team, error) {
	var teams []Team
	var teamOne, teamTwo *Team

	if strings.Trim(input, " ") == "" {
		return nil, errors.New("no input provided")
	}

	getScoreAndAddTeam := func(gameScore string) (name string, score int, err error) {
		name, score, err = getTeamScore(gameScore)
		if err != nil {
			return "", 0, err
		}

		if i := findTeam(name, teams); i == -1 {
			teams = append(teams, Team{name: name})
		}

		return
	}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		gameScores := strings.Split(line, ",")
		if len(gameScores) != 2 {
			return nil, fmt.Errorf("team scores required = 2, found %d", len(gameScores))
		}

		teamOneName, teamOneScore, err := getScoreAndAddTeam(gameScores[0])
		if err != nil {
			return nil, err
		}

		teamTwoName, teamTwoScore, err := getScoreAndAddTeam(gameScores[1])
		if err != nil {
			return nil, err
		}

		var index int
		if index = findTeam(teamOneName, teams); index == -1 {
			return nil, fmt.Errorf("team %s not found", teamOneName)
		}
		teamOne = &teams[index]

		if index = findTeam(teamTwoName, teams); index == -1 {
			return nil, fmt.Errorf("team %s not found", teamTwoName)
		}
		teamTwo = &teams[index]

		if teamOneScore == teamTwoScore {
			teamOne.points += pointsDraw
			teamTwo.points += pointsDraw
		} else if teamOneScore > teamTwoScore {
			teamOne.points += pointsWin
			teamTwo.points += pointsLose
		} else {
			teamOne.points += pointsLose
			teamTwo.points += pointsWin
		}
	}

	// Sort teams
	sort.Slice(teams, func(i, j int) bool {
		if teams[i].points == teams[j].points {
			return teams[i].name < teams[j].name
		} else {
			return teams[i].points > teams[j].points
		}
	})

	// Rank teams
	var lastPoints int
	tmpRank := 1
	for i := range teams {
		team := &teams[i]
		if i == 0 {
			lastPoints = team.points
			team.rank = tmpRank
			continue
		}

		if team.points != lastPoints {
			tmpRank = i + 1
		}

		team.rank = tmpRank
		lastPoints = team.points
	}

	return teams, nil
}

func getTeamScore(teamScoreStr string) (string, int, error) {
	teamScoreStr = strings.Trim(teamScoreStr, " ")

	// Find last space. Team name could contain spaces
	i := strings.LastIndex(teamScoreStr, " ")
	if i == -1 {
		return "", 0, errors.New("invalid line format")
	}

	name := teamScoreStr[:i]
	if strings.Trim(name, " ") == "" {
		return "", 0, errors.New("a valid team name is required")
	}

	score, err := strconv.Atoi(teamScoreStr[i+1:])
	if err != nil || score < 0 {
		return "", 0, fmt.Errorf("%s is not a valid score", teamScoreStr[i+1:])
	}

	return name, score, nil
}

func findTeam(name string, teams []Team) int {
	for i, team := range teams {
		if name == team.name {
			return i
		}
	}
	return -1
}

func GetFormattedRankedTeams(teams []Team) string {
	var result string
	for i, team := range teams {
		if i > 0 {
			result += "\n"
		}
		result += fmt.Sprintf("%d. %s, %d pts", team.rank, team.name, team.points)
	}

	return result
}

func main() {
	filePathPtr := flag.String("file", "", "Path to file containing league game results")
	flag.Parse()

	var input string
	if *filePathPtr != "" {
		b, err := ioutil.ReadFile(*filePathPtr)
		if err != nil {
			log.Fatal(err)
		}
		input = string(b)
	} else if len(os.Args) > 1 {
		input = os.Args[1]
	} else {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		input = string(b)
	}

	input = strings.Trim(input, "\n")

	teams, err := GetRankedTeams(input)
	if err != nil {
		log.Fatal(err)
	}

	output := GetFormattedRankedTeams(teams)

	fmt.Println(output)
}
