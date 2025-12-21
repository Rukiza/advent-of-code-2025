package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("data/day3-test.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)

	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
}
