package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("data/day1.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var dial int64 = 50
	var counter int64 = 0
	for scanner.Scan() {
		text := scanner.Text()
		rotation := text[:1]
		count, err := strconv.ParseInt(text[1:], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		dial = rotateDial(dial, rotation, count)

		if dial == 0 {
			counter++
		}
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(counter)
}

func rotateDial(dial int64, rotation string, count int64) int64 {
	switch rotation {
	case "R":
		return (dial + count) % 100
	case "L":
		tmp := (dial - count) % 100
		if tmp < 0 {
			return 100 + tmp
		}
		return tmp
	default:
		return 50
	}
}
