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
	part1()
	part2()

}

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("file cant open")
	}
	defer file.Close()

	var invalidIdSum int64 = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		values := strings.Split(line, ",")
		for _, val := range values {
			ranges := strings.Split(val, "-")
			start, err := strconv.ParseInt(ranges[0], 10, 64)
			if err != nil {
				log.Fatalf("range %v is cant convert to int64", ranges[0])
			}
			end, err := strconv.ParseInt(ranges[1], 10, 64)
			if err != nil {
				log.Fatalf("range %v is cant convert to int64", ranges[1])
			}

			var i int64
			for i = start; i <= end; i++ {
				numberString := strconv.FormatInt(i, 10)
				if numberString[len(numberString)/2:] == numberString[:len(numberString)/2] {
					invalidIdSum += i
				}
			}

		}
	}
	fmt.Printf("Result of part1: %d\n", invalidIdSum)
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("file cant open")
	}
	defer file.Close()

	var invalidIdSum int64 = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		values := strings.Split(line, ",")
		for _, val := range values {
			ranges := strings.Split(val, "-")
			start, err := strconv.ParseInt(ranges[0], 10, 64)
			if err != nil {
				log.Fatalf("range %v is cant convert to int64", ranges[0])
			}
			end, err := strconv.ParseInt(ranges[1], 10, 64)
			if err != nil {
				log.Fatalf("range %v is cant convert to int64", ranges[1])
			}

			var i int64
			for i = start; i <= end; i++ {
				numberString := strconv.FormatInt(i, 10)
				for j := 1; j <= len(numberString)/2; j++ {
					subStrings := strings.Split(numberString, numberString[:j])
					isRepeating := true
					for _, subString := range subStrings {
						if subString != "" {
							isRepeating = false
							break
						}
					}
					if isRepeating {
						invalidIdSum += i
						fmt.Println(i)
						break
					}
				}
			}

		}
	}
	fmt.Printf("Result of part2: %d\n", invalidIdSum)

}
