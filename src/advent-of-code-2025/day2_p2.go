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
	file, err := os.Open("data/day2.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(splitOnComma)

	var counter int64 = 0
	for scanner.Scan() {
		text := scanner.Text()
		// fmt.Println(text)

		value := strings.Split(strings.TrimSpace(text), "-")
		start, err := strconv.ParseInt(value[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		end, err := strconv.ParseInt(value[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(start)
		// fmt.Println(end)

		var i int64 = 0
		for start+i <= end {
			counter += invalidId(start + i)
			i++
		}
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(counter)
}

func invalidId(id int64) int64 {
	sid := strconv.FormatInt(id, 10)

	// If you cant half something why bother.
	i := 1
	for i <= len(sid) {
		// Is a divisiable step
		if len(sid)%i != 0 {
			i++
			continue
		}
		// Get each step
		step := len(sid) / i

		j := 2 // Skip 1 due to prev
		prev := sid[:step]
		allMatch := true

		// Would just mean step is length
		if j*step > len(sid) {
			i++
			continue
		}
		for j*step <= len(sid) {
			toCheck := sid[step*(j-1) : step*j]
			if prev != toCheck {
				allMatch = false
			}
			j++
		}

		if allMatch {
			// We just return the value found ... good job
			return id
		}

		i++
	}

	return 0
}

// https://stackoverflow.com/questions/33068644/how-a-scanner-can-be-implemented-with-a-custom-split
func splitOnComma(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	i := 0
	if i = strings.Index(string(data), ","); i >= 0 {
		return i + 1, data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return
}
