package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	inputfile, err := os.Open("./input.txt")
	if err != nil {
		panic("Can't read input file")
	}

	scanner := bufio.NewScanner(inputfile)

	lCol := make([]int, 0)
	rCol := make([]int, 0)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		textline := scanner.Text()
		cols := strings.SplitN(textline, " ", 2)
		if len(cols) < 2 {
			continue
		}

		val, err := strconv.Atoi(cols[0])
		if err != nil {
			panic("failed to convert string to int")
		}

		lCol = append(lCol, val)
		val, err = strconv.Atoi(strings.TrimSpace(cols[1]))
		if err != nil {
			panic("failed to convert string to int")
		}

		rCol = append(rCol, val)
	}

	slices.Sort(lCol)
	slices.Sort(rCol)

	occurrences := map[int]int{}
	diff := []int{}
	for i := 0; i < len(lCol); i++ {
		occurrences[rCol[i]]++
		val := lCol[i] - rCol[i]
		if val < 0 {
			val = -val
		}
		diff = append(diff, val)
	}

	score := 0
	for i := 0; i < len(lCol); i++ {
		num := lCol[i]
		n, ok := occurrences[num]
		if !ok {
			continue
		}

		score += (num * n)
	}

	// fmt.Println("Left Col", lCol)
	// fmt.Println("Right Col", rCol)
	// fmt.Println("Right Col", diff)
	fmt.Printf("Sim -> %+v", occurrences)
	fmt.Printf("Score -> %+v", score)
	fmt.Println("Res ->", Sum(diff))
}

func Sum(numbers []int) int {
	// Sum the values
	sum := 0
	for _, value := range numbers {
		sum += value
	}

	return sum
}
