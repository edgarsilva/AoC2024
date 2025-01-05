package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic("failed to load input file")
	}

	multEnabled := true
	multResults := []int{}
	for i := 0; i < len(input)-4; i++ {
		condStr := input[i : i+len("do()")]
		if string(condStr) == "do()" {
			multEnabled = true
			i += 3
			continue
		}

		if i+7 < len(input) {
			condStr = input[i : i+len("don't()")]
		}
		if string(condStr) == "don't()" {
			multEnabled = false
			i += len("don't()") - 1
			continue
		}

		ridx := i + 4
		ins := input[i:ridx]
		if string(ins) != "mul(" {
			continue
		}

		if !multEnabled {
			continue
		}

		// fmt.Println("valid ins ->", string(ins))
		// until closing parenth
		numbers := []byte{}
		for j := ridx; j < len(input) && input[j] != ')'; j++ {
			numbers = append(numbers, input[j])
		}

		// fmt.Println("numbers ->", string(numbers))
		if !allDigits(numbers) {
			i = ridx
			continue
		}

		operands := strings.SplitN(string(numbers), ",", 2)
		if len(operands) != 2 {
			i = ridx
			continue
		}

		// fmt.Println("operands", operands)
		lop, err := strconv.Atoi(operands[0])
		if err != nil {
			panic("failed to parse number from string")
		}
		rop, err := strconv.Atoi(operands[1])
		if err != nil {
			panic("failed to parse number from string")
		}

		multResults = append(multResults, lop*rop)
		// fmt.Println("operation: ", string(numbers))
	}

	fmt.Println("All operations: ", multResults)
	fmt.Println("Sum total: ", Sum(multResults))
}

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}

	return sum
}

func allDigits(numbers []byte) bool {
	for _, num := range numbers {
		if (num < '0' || num > '9') && num != ',' {
			return false
		}
	}

	return true
}
