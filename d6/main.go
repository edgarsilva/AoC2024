package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type Empty struct{}

type Guard struct {
	dir byte
	pos Pos
}

type Pos struct {
	y int
	x int
}

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
	for scanner.Scan() {
		input = append(input, []byte(scanner.Text()))
		fmt.Println(string(scanner.Text()))
	}

	initialGuard := NewGuard(input)

	loopCount := 0
	// PrintMap(input)
	fmt.Println("Loop Count ", loopCount)
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if input[y][x] != '.' {
				continue
			}
			input[y][x] = '#'

			route := map[Guard]Empty{}
			guard := NewGuard(input)
			for WithinBounds(input, guard) {
				if _, seen := route[guard]; seen {
					input[guard.pos.y][guard.pos.x] = '.'
					loopCount++
					break
				}
				route[guard] = Empty{}
				guard = guard.Move(input)
				PrintMap(input)
				// fmt.Println("Loop Count ", loopCount)
				time.Sleep(500 * time.Millisecond)
			}
			input[y][x] = '.'
			input[initialGuard.pos.y][initialGuard.pos.x] = initialGuard.dir
			// PrintMap(input)
			// fmt.Println("Loop Count ", loopCount)
			time.Sleep(50 * time.Millisecond)
		}
	}

	fmt.Println("Loops:", loopCount)
}

func NewGuard(input [][]byte) Guard {
	g := Guard{}
outer:
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			if input[i][j] == '^' || input[i][j] == '>' || input[i][j] == '<' || input[i][j] == 'v' {
				g.dir = input[i][j]
				g.pos.y = i
				g.pos.x = j
				break outer
			}
		}
	}

	return g
}

func (g Guard) Move(input [][]byte) Guard {
	ng := Guard{}
	rowN := len(input)
	colN := len(input[0])
	switch input[g.pos.y][g.pos.x] {
	case '^':
		if g.pos.y-1 < 0 {
			ng = Guard{dir: g.dir, pos: Pos{y: -1, x: g.pos.x}}
			break
		}
		if input[g.pos.y-1][g.pos.x] == '#' {
			ng = Guard{dir: '>', pos: g.pos}
			break
		}
		ng = Guard{dir: g.dir, pos: Pos{y: g.pos.y - 1, x: g.pos.x}}
	case '>':
		if g.pos.x+1 >= colN {
			ng = Guard{dir: g.dir, pos: Pos{y: g.pos.y, x: g.pos.x + 1}}
			break
		}
		if input[g.pos.y][g.pos.x+1] == '#' {
			ng = Guard{dir: 'v', pos: g.pos}
			break
		}
		ng = Guard{dir: g.dir, pos: Pos{y: g.pos.y, x: g.pos.x + 1}}
	case 'v':
		if g.pos.y+1 >= rowN {
			ng = Guard{dir: g.dir, pos: Pos{y: g.pos.y + 1, x: g.pos.x}}
			break
		}
		if input[g.pos.y+1][g.pos.x] == '#' {
			ng = Guard{dir: '<', pos: g.pos}
			break
		}
		ng = Guard{dir: g.dir, pos: Pos{y: g.pos.y + 1, x: g.pos.x}}
	case '<':
		if g.pos.x-1 < 0 {
			ng = Guard{dir: g.dir, pos: Pos{y: g.pos.y, x: g.pos.x - 1}}
			break
		}
		if input[g.pos.y][g.pos.x-1] == '#' {
			ng = Guard{dir: '^', pos: g.pos}
			break
		}
		ng = Guard{dir: g.dir, pos: Pos{y: g.pos.y, x: g.pos.x - 1}}
	default:
		return g
	}

	input[g.pos.y][g.pos.x] = '.'

	if WithinBounds(input, ng) {
		input[ng.pos.y][ng.pos.x] = ng.dir
	}

	return ng
}

func PrintMap(input [][]byte) {
	n := len(input)
	fmt.Print("\033[", n, "A\033[1000D")
	for _, line := range input {
		fmt.Println(string(line))
	}
}

func WithinBounds(input [][]byte, g Guard) bool {
	y := g.pos.y
	x := g.pos.x
	return y >= 0 && y < len(input) && x >= 0 && x < len(input[0])
}

func CountVisited(input [][]byte) int {
	count := 0
	for _, line := range input {
		for _, char := range line {
			if char == 'X' {
				count++
			}
		}
	}

	return count
}
