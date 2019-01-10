package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	// Set our CSV file by using CLI arg parsing
	var file string
	flag.StringVar(&file, "file", "problems.csv", "Specify a .CSV file to use for the questions")
	flag.Parse()

	// Read the CSV file contents
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("Something went wrong:\n%s", err)
	}
	r := csv.NewReader(f)

	fmt.Println("#############################")
	fmt.Println("# Gophercises -- Exercise 1 #")
	fmt.Printf("#############################\n\n")
	fmt.Printf("Welcome. We will be using questions from %s. Let's begin...\n\n", file)
	time.Sleep(3 * time.Second)

	score := 0
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something went wrong:\n%s", err)
		}

		fmt.Println(line[0], "= ?")
		var answer string
		fmt.Scanln(&answer)
		if answer == line[1] {
			score++
		}
	}

	if score >= 10 {
		fmt.Printf("Well done! Your score is %d. 😁", score)
	} else if score >= 5 && score <= 9 {
		fmt.Printf("Not bad! You got %d correct!", score)
	} else {
		// Don't let them know their score....
		fmt.Println("################################")
		fmt.Println("# ERROR READING SCORE - EXITING #")
		fmt.Println("################################")
		msg := os.Getenv("USER") + " scored a " + strconv.Itoa(score) + ". I didn't want to let them know..."
		for _, v := range []byte(msg) {
			fmt.Println(v)
		}
	}
}
