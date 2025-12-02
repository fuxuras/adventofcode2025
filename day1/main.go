package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	part1()
	part2()
}

func part1() {

	file, err := os.Open("input.csv")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	zeroCrossing := 0
	dial := 50
	for scanner.Scan() {
		line := scanner.Text()

		direction := line[0]
		stepStr := line[1:]

		step, err := strconv.Atoi(stepStr)

		if err != nil {
			log.Fatal("Error parsing step size to integer")
		}

		switch direction {
		case 'L':
			dial = (dial - step) % 100
			if dial < 0 {
				dial += 100
			}
		case 'R':
			dial = (dial + step) % 100

		default:
			log.Fatal("Error line has no direction Left or Right")
		}

		if dial == 0 {
			zeroCrossing++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error during scan %s", err)
	}

	fmt.Printf("Result of Part1: %d\n", zeroCrossing)
}

func part2() {
	file, err := os.Open("input.csv")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	zeroCrossing := 0
	dial := 50
	for scanner.Scan() {
		line := scanner.Text()

		direction := line[0]
		stepStr := line[1:]

		step, err := strconv.Atoi(stepStr)

		if err != nil {
			log.Fatal("Error parsing step size to integer")
		}

		zeroCrossing += step / 100
		step = step % 100

		switch direction {
		case 'L':
			if step >= dial && dial != 0 {
				zeroCrossing++
			}
			dial = (dial - step + 100) % 100
		case 'R':
			if step >= 100-dial {
				zeroCrossing++
			}
			dial = (dial + step) % 100
		default:
			log.Fatal("Error line has no direction Left or Right")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error during scan %s", err)
	}

	fmt.Printf("Result of Part2: %d\n", zeroCrossing)
}
