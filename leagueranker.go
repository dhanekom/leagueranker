package leagueranker

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

type Team struct {
	name string
	pts  int
	rank int
}

type Ranker struct {
	regex *regexp.Regexp
	teams map[string]*Team
}

func NewRanker() (Ranker, error) {
	r, err := regexp.Compile(`^([a-zA-Z0-9-_\s]+)\s+(\d+),\s*([a-zA-Z0-9-_\s]+)\s+(\d+)$`)
	if err != nil {
		return Ranker{}, err
	}
	ranker := Ranker{}
	ranker.regex = r
	ranker.teams = make(map[string]*Team)
	return ranker, nil
}

func (r Ranker) Parse(line string) error {
	scoreSubMatches := r.regex.FindStringSubmatch(line)
	if len(scoreSubMatches) != 5 {
		return errors.New("invalid line format")
	}

	scoreSections := scoreSubMatches[1:]

	var teamA, teamB *Team
	var ok bool

	teamAName, teamBName := scoreSections[0], scoreSections[2]
	teamAScore, _ := strconv.Atoi(scoreSections[1])
	teamBScore, _ := strconv.Atoi(scoreSections[3])

	if teamA, ok = r.teams[teamAName]; !ok {
		teamA = &Team{name: teamAName}
		r.teams[teamAName] = teamA
	}
	if teamB, ok = r.teams[teamBName]; !ok {
		teamB = &Team{name: teamBName}
		r.teams[teamBName] = teamB
	}

	if teamAScore == teamBScore {
		teamA.pts += 1
		teamB.pts += 1
	} else if teamAScore > teamBScore {
		teamA.pts += 3
	} else {
		teamB.pts += 3
	}

	return nil
}

func (r Ranker) RankedTeams() []Team {
	result := make([]Team, 0)
	for _, team := range r.teams {
		result = append(result, *team)
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].pts == result[j].pts {
			return result[i].name < result[j].name
		} else {
			return result[i].pts > result[j].pts
		}
	})

	rank := 1
	for i := 0; i < len(result); i++ {
		if i > 0 && result[i].pts != result[i-1].pts {
			rank = i + 1
		}
		result[i].rank = rank
	}

	return result
}

func GetOutput(teams []Team) string {
	var result string
	var ptsDesc string
	sep := ""
	for _, team := range teams {
		if team.pts == 1 {
			ptsDesc = "pt"
		} else {
			ptsDesc = "pts"
		}
		result += fmt.Sprintf("%s%d. %s, %d %s", sep, team.rank, team.name, team.pts, ptsDesc)
		sep = "\n"
	}
	return result
}
