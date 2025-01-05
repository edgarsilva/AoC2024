package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	args := flag.Args()
	file, err := os.Open(args[0])
	if err != nil {
		panic("failed to open input file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rules := []string{}
	upds := []string{}
	linebreak := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			linebreak = true
			continue
		}

		if linebreak {
			upds = append(upds, line)
			continue
		}

		rules = append(rules, line)
	}

	fmt.Println("Rules")
	for _, r := range rules {
		fmt.Println(r)
	}

	fmt.Println("\nUpdates")
	for _, u := range upds {
		fmt.Println(u)
	}
	fmt.Println("")

	correct := []string{}
	incorrect := []string{}
	for _, upd := range upds {
		isCorrect := true
		for i := 0; i < len(rules); i++ {
			r := rules[i]
			parts := strings.Split(r, "|")
			if idx1 := strings.Index(upd, parts[0]); idx1 > -1 {
				idx2 := strings.Index(upd, parts[1])
				// fmt.Println("idx1 ", idx1, "idx2", idx2)
				if idx2 > -1 && idx2 < idx1 {
					isCorrect = false
					line := []byte(upd)
					line[idx1], line[idx1+1], line[idx2], line[idx2+1] = line[idx2], line[idx2+1], line[idx1], line[idx1+1]
					upd = string(line)
					i = -1
					continue
				}
			}
		}

		if isCorrect {
			correct = append(correct, upd)
		} else {
			incorrect = append(incorrect, upd)
		}
	}

	fmt.Println("Incorrect ones ->")
	for _, u := range incorrect {
		fmt.Println(u)
	}

	sum := 0
	for _, u := range incorrect {
		parts := strings.Split(u, ",")
		middleNum, err := strconv.Atoi(parts[len(parts)/2])
		if err != nil {
			panic("failed to convert Asccii to int")
		}
		sum += middleNum
	}

	fmt.Println("The sum of middle ones ->", sum)
}
