package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Empty struct{}

type Calibration struct {
	numbers []int
	res     int
}

var (
	operands    = map[string]string{"0": "+", "1": "*", "2": "|"}
	operandsStr = "+*|"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("you must pass an input file name")
	}

	file, err := os.Open(args[0])
	if err != nil {
		log.Fatal("failed to open input file")
	}

	scanner := bufio.NewScanner(file)
	input := [][]byte{}
	calibrations := []Calibration{}
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")
		if len(parts) < 2 {
			continue
		}
		resStr := strings.TrimSpace(parts[0])
		numsStr := strings.TrimSpace(parts[1])
		resInt, err := strconv.Atoi(resStr)
		if err != nil {
			log.Fatal("failed to convert string into int 2")
		}

		calibrations = append(calibrations, Calibration{
			res:     resInt,
			numbers: StrNumsToInt(numsStr),
		})

		input = append(input, []byte(scanner.Text()))
		// fmt.Println(string(scanner.Text()))
	}

	// operations := cartesianProduct("+*|", 3)
	// fmt.Println("Calibrations:", calibrations)
	operations := FormOperations(calibrations)
	// for _, op := range operations {
	// 	fmt.Println("operations:", op)
	// }
	fmt.Println("Result:", operations)
}

func FormOperations(cs []Calibration) int {
	// operations := []string{}
	results := map[int]Empty{}
	for _, c := range cs {
		// expected := c.res
		// got := 0
		nl := len(c.numbers)
		// perm := int(math.Pow(float64(3), float64(nl-1)))
		// perm := len(c.numbers) - 1
		// fmt.Println(perm, "perm for ", c.numbers)
		// acc := [][]byte{}
		// combo := zeroSlice(nl - 1)
		// for k := len(combo) - 1; k >= 0; k-- {
		// 	acc = recPerm(acc, combo, k)
		// 	if k-1 >= 0 {
		// 		combo[k-1] = byte((combo[k-1]-'0')+1) + '0'
		// 		combo[k] = '0'
		// 	}
		// }
		acc := cartesianProduct("+*|", nl-1)
		// fmt.Println(acc)

		// Result: 11387
		// Result: 438027111276610
		for i := 0; i < len(acc); i++ {
			symbols := acc[i]
			// operation := ""
			tmpRes := c.numbers[0]
			for j, symbol := range symbols {
				// n := int64(c.numbers[j])
				// nstr := strconv.Itoa(int(n))
				// operation += nstr + operands[string(symbol)]
				tmpRes = calcOperation(tmpRes, c.numbers[j+1], string(symbol))
			}
			// nstr := strconv.Itoa(c.numbers[len(c.numbers)-1])
			// res := strconv.Itoa(tmpRes)
			// operation += nstr + " = " + string(res)
			// operations = append(operations, operation)

			if c.res == tmpRes {
				results[tmpRes] = Empty{}
			}
		}
	}

	sum := 0
	for key := range results {
		sum += key
	}
	return sum
}

func FormOperationsP1(cs []Calibration) int {
	// operations := []string{}
	results := map[int]Empty{}
	for _, c := range cs {
		// expected := c.res
		// got := 0
		nl := len(c.numbers)
		// perm := int(math.Pow(float64(3), float64(nl-1)))
		// perm := len(c.numbers) - 1
		// fmt.Println(perm, "perm for ", c.numbers)
		acc := [][]byte{}
		combo := zeroSlice(nl - 1)
		for k := len(combo) - 1; k >= 0; k-- {
			acc = recPerm(acc, combo, k)
			if k-1 >= 0 {
				combo[k-1] = byte((combo[k-1]-'0')+1) + '0'
				combo[k] = '0'
			}
		}
		// for i := 0; i < perm; i++ {
		// 	padding := strconv.Itoa(nl - 1)
		// 	rep := "%0" + padding + "b"
		// 	symbols := strings.Split(fmt.Sprintf(rep, i), "")
		// 	operation := ""
		// 	tmpRes := c.numbers[0]
		// 	for j, symbol := range symbols {
		// 		n := int64(c.numbers[j])
		// 		nstr := strconv.Itoa(int(n))
		// 		operation += nstr + operands[symbol]
		// 		tmpRes = calcOperation(tmpRes, c.numbers[j+1], operands[symbol])
		// 	}
		// 	nstr := strconv.Itoa(c.numbers[len(c.numbers)-1])
		// 	res := strconv.Itoa(tmpRes)
		// 	operation += nstr + " = " + string(res)
		// 	operations = append(operations, operation)
		//
		// 	if c.res == tmpRes {
		// 		results[tmpRes] = Empty{}
		// 	}
		// }
	}

	sum := 0
	for key := range results {
		sum += key
	}
	return sum
}

func cartesianProduct(chars string, n int) []string {
	if n <= 0 {
		return []string{""} // Base case: one empty string
	}

	// Get the Cartesian product of (n-1)
	smallerProduct := cartesianProduct(chars, n-1)

	var result []string
	for _, prefix := range smallerProduct {
		for _, char := range chars {
			result = append(result, prefix+string(char))
		}
	}

	return result
}

func StrNumsToInt(str string) []int {
	str = strings.TrimSpace(str)
	strNums := strings.Split(str, " ")
	intNums := []int{}

	for _, n := range strNums {
		resInt, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal("failed to convert string into int 3")
		}

		intNums = append(intNums, resInt)
	}

	return intNums
}

func calcOperation(op1, op2 int, symbol string) int {
	switch symbol {
	case "|":
		op1Str := strconv.Itoa(op1)
		op2Str := strconv.Itoa(op2)
		res, err := strconv.Atoi(op1Str + op2Str)
		if err != nil {
			panic("failed to concatenate int str")
		}
		return res
	case "*":
		return op1 * op2
	default:
		return op1 + op2
	}
}

func zeroSlice(l int) []byte {
	sl := make([]byte, l)
	for i := 0; i < len(sl); i++ {
		sl[i] = byte(0) + '0'
	}

	return sl
}

func recPerm(acc [][]byte, sl []byte, idx int) [][]byte {
	init := sl[idx] - '0'
	if idx < 0 {
		return acc
	}

	// fmt.Println("idx ->", idx)
	for j := init; j < 3; j++ {
		sl[idx] = byte(j) + '0'
		if idx+1 < len(sl) {
			acc = recPerm(acc, sl, idx+1)
		} else {
			dst := make([]byte, len(sl))
			copy(dst, sl)
			// fmt.Println(string(dst))
			acc = append(acc, dst)
		}
	}

	sl[idx] = '0'
	return acc
}
