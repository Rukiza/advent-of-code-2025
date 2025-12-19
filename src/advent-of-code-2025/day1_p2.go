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
		var amount int64 = 0
		text := scanner.Text()
		rotation := text[:1]
		count, err := strconv.ParseInt(text[1:], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		dial, amount = rotateDial(dial, rotation, count)
		// fmt.Printf("%s %d\n", text, amount)

		counter = counter + amount
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(counter)
}

func rotateDial(dial int64, rotation string, count int64) (int64, int64) {
	switch rotation {
	case "R":
		tmp := (dial + count) % 100
		amount := count / 100
		if dial == 0 {
		} else if count%100 == 0 {
		} else if tmp <= dial {
			amount = amount + 1
		}
		return tmp, amount
	case "L":
		tmp := (dial - count) % 100
		amount := count / 100
		if dial == 0 {
		} else if count%100 == 0 {
		} else if tmp <= 0 {
			amount++
		}
		if tmp < 0 {
			return 100 + tmp, amount
		}
		return tmp, amount
	default:
		return 50, 0
	}
}
