package main

import (
	"testing"
)

var addCases = []struct {
	description string
	in          string
	want        string
}{
	{
		"No input",
		"",
		"no input provided",
	},
	{
		"No space between team and score",
		`Lions3, Snakes 3`,
		`invalid line format`,
	},
	{
		"No commas between team scores",
		`Lions 3 Snakes 1`,
		"team scores required = 2, found 1",
	},
	{
		"Too many teams",
		`Lions 3, Snakes 1, Sharks 5`,
		"team scores required = 2, found 3",
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
3. FC Awesome, 1 pts
3. Snakes, 1 pts
5. Grouches, 0 pts`,
	},
}

func TestGetRankedTeams(t *testing.T) {
	for _, tc := range addCases {
		got, err := GetRankedTeams(tc.in)

		if err != nil {
			if err.Error() != tc.want {
				t.Fatalf(`FAIL: %s
got:
%s
want:
%s`, tc.description, err.Error(), tc.want)
			}
		} else {
			if got != tc.want {
				t.Fatalf(`FAIL: %s
got:
%s
want:
%s`, tc.description, got, tc.want)
			}
		}
	}
}
