package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type Quiz struct {
	Question string
	Answer   string
}

func main() {

	csvFilename := flag.String("csv", "problem.csv", "a csv file in the format of 'question,answer'")

	timeLimit := flag.Int("limit", 30, "Time limit")

	flag.Parse()

	file, err := os.Open(*csvFilename)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully CSV opened!")
	defer file.Close()

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	checkAnswers(file, timer)

}

func checkAnswers(file *os.File, timer *time.Timer) {

	var answer string
	var rightAnswer, wrongAnswer int64

	linesFile, err := csv.NewReader(file).ReadAll()

	if err != nil {
		fmt.Println(err)
	}

	for _, col := range linesFile {
		quiz := &Quiz{
			Question: col[0],
			Answer:   col[1],
		}

		fmt.Println("What?", quiz.Question)
		answerCh := make(chan string)
		go func() {
			fmt.Scanf("%s", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("The result was... %d correct and %d wrong!", rightAnswer, wrongAnswer)
			return
		case answer := <-answerCh:
			if answer == quiz.Answer {
				rightAnswer++
			} else {
				wrongAnswer++
			}
		}

	}
	fmt.Printf("The result was... %d correct and %d wrong!", rightAnswer, wrongAnswer)
}
