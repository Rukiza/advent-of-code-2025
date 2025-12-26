package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"math"
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

	numbers := list.New()

	numberln := 12

	// Find largest except the end value
	currentStrIndexStart := 0
	for numbers.Len() < numberln {
		bestIndex := currentStrIndexStart
		i := currentStrIndexStart
		current, err := strconv.ParseInt(l[currentStrIndexStart], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		for i < len(l)-(numberln-numbers.Len())+1 {
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
		numbers.PushBack(current)
		currentStrIndexStart = bestIndex + 1
	}

	// fmt.Println(v)
	var j int64 = 0
	var total int64 = 0
	value := numbers.Front()
	for j < int64(numbers.Len()) {
		fmt.Print(value.Value)
		total = total + value.Value.(int64)*int64(math.Pow(10, float64(int64(numbers.Len())-j-1)))
		value = value.Next()
		j++
	}
	fmt.Println("")
	return total
}
