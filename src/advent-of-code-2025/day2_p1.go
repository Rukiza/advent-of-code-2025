package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("data/day1.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
}
