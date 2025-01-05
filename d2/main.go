package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputReader, err := os.Open("input.txt")
	if err != nil {
		panic("failed to read file from disk")
	}

	scanner := bufio.NewScanner(inputReader)

	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, " ")
		if len(numbers) < 2 {
			continue
		}

		if ValidLine(numbers) {
			result++
			continue
		}

		dampen := make([]string, 0, len(numbers)-1)
		for i := range numbers {
			dampen = dampen[:0]
			dampen = append(dampen, numbers[:i]...)
			dampen = append(dampen, numbers[i+1:]...)

			if !ValidLine(dampen) {
				continue
			}

			result++
			break
		}
	}

	fmt.Println("Safe Results ->", result)
}

func ValidLine(numbers []string) bool {
	prev, err := strconv.Atoi(numbers[0])
	if err != nil {
		panic("failed to to convert string number to int")
	}

	prevDirection := 0
	isSafe := true
	for i := 1; i < len(numbers); i++ {
		current, err := strconv.Atoi(numbers[i])
		if err != nil {
			panic("failed to to convert string number to int")
		}

		change := current - prev
		if change < 0 {
			change = -change
		}

		if change > 3 {
			return false
		}

		currentDirection := 0
		if current > prev {
			currentDirection = 1
		} else if current < prev {
			currentDirection = -1
		} else {
			return false
		}

		if prevDirection == 0 {
			prevDirection = currentDirection
		}

		if currentDirection != prevDirection {
			return false
		}

		prev = current
		prevDirection = currentDirection
	}

	return isSafe
}
