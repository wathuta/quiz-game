package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file containng mathematical problems in the formart of question,answer")
	timeLimit := flag.Int("limit", 30, "the time limit for user to answer questions")
	flag.Parse()
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("unable to open csv file %s ", csvFilename))
		return
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("unable to read the content in the file %s", csvFilename))
	}

	problem := parselines(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	counter := 0
	for i, p := range problem {
		fmt.Printf("problem #%v:%s = \n", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("you got %d questions right out of %d problems \n", counter, len(problem))
			return
		case answer := <-answerCh:
			if answer == p.a {
				counter++
			}
		}

	}
	fmt.Printf("you got %d questions right out of %d problems \n", counter, len(problem))
}

func parselines(lines [][]string) []problems {
	ret := make([]problems, len(lines))
	for i, line := range lines {
		ret[i] = problems{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problems struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	return
}
