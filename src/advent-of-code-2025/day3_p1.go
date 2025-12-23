package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
		// fmt.Println(text)

		total += findMeBigNumber(text)
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(total)
}

func findMeBigNumber(text string) int64 {
	l := strings.Split(strings.TrimSpace(text), "")

	i := 1
	current, err := strconv.ParseInt(l[0], 10, 64)
	bestIndex := 0
	if err != nil {
		log.Fatal(err)
	}

	// Find largest except the end value
	for i < len(l)-1 {
		temp, err := strconv.ParseInt(l[i], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		if current < temp {
			current = temp
			bestIndex = i
		}
		i++
	}

	// Find largest value using current as start
	secondBest := bestIndex + 1
	j := secondBest + 1
	current2, err := strconv.ParseInt(l[secondBest], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	for j < len(l) {
		temp, err := strconv.ParseInt(l[j], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		if current2 < temp {
			current2 = temp
			secondBest = j
		}
		j++
	}
	// v := current*10 + current2
	// fmt.Println(v)
	return current*10 + current2
}
