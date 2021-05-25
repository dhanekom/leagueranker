package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"leagueranker"
	"log"
	"os"
)

func main() {
	var r io.Reader

	filepath := flag.String("file", "", "file path to input file")
	flag.Parse()
	if *filepath != "" {
		f, err := os.Open(*filepath)
		if err != nil {
			log.Fatalln(err)
		}
		defer f.Close()

		r = f
	} else {
		r = bufio.NewReader(os.Stdin)
	}

	ranker, err := leagueranker.NewRanker()
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		err := ranker.Parse(scanner.Text())
		if err != nil {
			log.Fatalln(err)
		}
	}

	teams := ranker.RankedTeams()
	fmt.Println(leagueranker.GetOutput(teams))
}
