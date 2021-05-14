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

// GetRankedTeams take league game scores as a string and outputs teams ordered by their rank as formatted text
func GetRankedTeams(input string) (string, error) {
	var result string
	var coolTeams []Team
	var teamOne, teamTwo *Team

	if strings.Trim(input, " ") == "" {
		return "", errors.New("no input provided")
	}

	getScoreAndAddTeam := func(gameScore string) (name string, score int, err error) {
		name, score, err = getTeamScore(gameScore)
		if err != nil {
			return "", 0, err
		}

		if i := findTeam(name, coolTeams); i == -1 {
			coolTeams = append(coolTeams, Team{name: name})
		}

		return
	}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		gameScores := strings.Split(line, ",")
		if len(gameScores) != 2 {
			return "", fmt.Errorf("team scores required = 2, found %d", len(gameScores))
		}

		teamOneName, teamOneScore, err := getScoreAndAddTeam(gameScores[0])
		if err != nil {
			return "", err
		}

		teamTwoName, teamTwoScore, err := getScoreAndAddTeam(gameScores[1])
		if err != nil {
			return "", err
		}

		var index int
		if index = findTeam(teamOneName, coolTeams); index == -1 {
			return "", fmt.Errorf("team %s not found", teamOneName)
		}
		teamOne = &coolTeams[index]

		if index = findTeam(teamTwoName, coolTeams); index == -1 {
			return "", fmt.Errorf("team %s not found", teamTwoName)
		}
		teamTwo = &coolTeams[index]

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
	sort.Slice(coolTeams, func(i, j int) bool {
		if coolTeams[i].points == coolTeams[j].points {
			return coolTeams[i].name < coolTeams[j].name
		} else {
			return coolTeams[i].points > coolTeams[j].points
		}
	})

	// Rank teams
	var lastPoints int
	tmpRank := 1
	for i := range coolTeams {
		team := &coolTeams[i]
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

	// Format output
	for i, team := range coolTeams {
		if i > 0 {
			result += "\n"
		}
		result += fmt.Sprintf("%d. %s, %d pts", team.rank, team.name, team.points)
	}
	return result, nil
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

	output, err := GetRankedTeams(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)
}
