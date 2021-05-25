package leagueranker

import (
	"bufio"
	"log"
	"strings"
	"testing"
)

var addCases = []struct {
	description string
	in          string
	want        string
}{
	{
		"No space between team and score",
		`Lions3, Snakes 3`,
		`invalid line format`,
	},
	{
		"No commas between team scores",
		`Lions 3 Snakes 1`,
		"invalid line format",
	},
	{
		"Too many teams",
		`Lions 3, Snakes 1, Sharks 5`,
		"invalid line format",
	},
	{
		"Successful sorting",
		`Cats 1, Astronaut 1
Bats 1, Astronaut 1
Cats 1, Bats 1`,
		`1. Astronaut, 2 pts
1. Bats, 2 pts
1. Cats, 2 pts`,
	},
	{
		"Case of team names matter",
		`Lions 5, Snakes 0
lions 1, FC Awesome 0`,
		`1. Lions, 3 pts
1. lions, 3 pts
3. FC Awesome, 0 pts
3. Snakes, 0 pts`,
	},
	{
		"Success test case",
		`Lions 3, Snakes 3
Tarantulas 1, FC Awesome 0
Lions 1, FC Awesome 1
Tarantulas 3, Snakes 1
Lions 4, Grouches 0`,
		`1. Tarantulas, 6 pts
2. Lions, 5 pts
3. FC Awesome, 1 pt
3. Snakes, 1 pt
5. Grouches, 0 pts`,
	},
}

func TestRankedTeams(t *testing.T) {
	displayResult := func(description, got, want string) {
		t.Fatalf(`FAIL: %s
got:
%s
want:
%s`, description, got, want)
	}

	for _, tc := range addCases {
		r := bufio.NewReader(strings.NewReader(tc.in))
		ranker, err := NewRanker()
		if err != nil {
			log.Fatalln(err)
		}
		scanner := bufio.NewScanner(r)
		hasErrors := false
		for scanner.Scan() {
			err = ranker.Parse(scanner.Text())
			if err != nil {
				hasErrors = true
				if err.Error() != tc.want {
					displayResult(tc.description, err.Error(), tc.want)
					continue
				}
			}
		}

		if hasErrors {
			continue
		}

		teams := ranker.RankedTeams()
		got := GetOutput(teams)

		if got != tc.want {
			displayResult(tc.description, got, tc.want)
		}
	}
}

func BenchmarkParse(b *testing.B) {
	ranker, _ := NewRanker()
	for i := 0; i < b.N; i++ {
		ranker.Parse("Lions 3, Snakes 3")
	}
}

func BenchmarkRankedTeams(b *testing.B) {
	ranker, _ := NewRanker()
	ranker.Parse("Lions 3, Snakes 3")
	ranker.Parse("Tarantulas 1, FC Awesome 0")
	ranker.Parse("Lions 1, FC Awesome 1")
	ranker.Parse("Tarantulas 3, Snakes 1")
	ranker.Parse("Lions 4, Grouches 0")
	for i := 0; i < b.N; i++ {
		ranker.RankedTeams()
	}
}
