package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("data/day3.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var total int64 = 0
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(total)
}
