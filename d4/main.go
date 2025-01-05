package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	filereader, err := os.Open(args[0])
	if err != nil {
		panic("faild to read input file")
	}
	defer filereader.Close()

	data := [][]byte{}
	scanner := bufio.NewScanner(filereader)
	for scanner.Scan() {
		data = append(data, []byte(scanner.Text()))
	}

	count := 0
	rn := len(data)
	cn := len(data[0])
	for row := 1; row < rn-1; row++ {
		for col := 1; col < cn-1; col++ {
			if data[row][col] != 'A' {
				continue
			}

			if data[row-1][col-1] == 'M' && data[row+1][col-1] == 'M' && data[row-1][col+1] == 'S' && data[row+1][col+1] == 'S' {
				count += 1
			}
			if data[row-1][col-1] == 'S' && data[row+1][col-1] == 'S' && data[row-1][col+1] == 'M' && data[row+1][col+1] == 'M' {
				count += 1
			}
			if data[row-1][col-1] == 'M' && data[row+1][col-1] == 'S' && data[row-1][col+1] == 'M' && data[row+1][col+1] == 'S' {
				count += 1
			}
			if data[row-1][col-1] == 'S' && data[row+1][col-1] == 'M' && data[row-1][col+1] == 'S' && data[row+1][col+1] == 'M' {
				count += 1
			}
		}
	}

	fmt.Println("The XMAS count:", count)
}

// Lesson learned in reusing memory with scanner.Bytes()
// make sure to always make a copy of the byte slice.
func mainPart1() {
	flag.Parse()
	args := flag.Args()
	filereader, err := os.Open(args[0])
	if err != nil {
		panic("faild to read input file")
	}
	defer filereader.Close()

	data := [][]byte{}
	scanner := bufio.NewScanner(filereader)
	for scanner.Scan() {
		data = append(data, []byte(scanner.Text()))
	}

	count := 0
	rn := len(data)
	cn := len(data[0])
	count2 := 0
	fmt.Println("line count:\n", string(len(data)))
	fmt.Println("First line:\n", string(data[0]))
	fmt.Println("Last line:\n", string(data[len(data)-1]))
	fmt.Println("Row count ->", rn)
	fmt.Println("Col count ->", rn)
	for row := 0; row < rn; row++ {
		count2 += countXmas(data[row])
		for col := 0; col < cn; col++ {
			if data[row][col] != 'X' {
				continue
			}

			if row > 2 {
				if col-3 >= 0 && data[row-1][col-1] == 'M' && data[row-2][col-2] == 'A' && data[row-3][col-3] == 'S' {
					count += 1
				}
				if data[row-1][col] == 'M' && data[row-2][col] == 'A' && data[row-3][col] == 'S' {
					count += 1
				}
				if col+3 < cn && data[row-1][col+1] == 'M' && data[row-2][col+2] == 'A' && data[row-3][col+3] == 'S' {
					count += 1
				}
			}

			if col+3 < cn && data[row][col+1] == 'M' && data[row][col+2] == 'A' && data[row][col+3] == 'S' {
				count += 1
			}
			if col-3 >= 0 && data[row][col-1] == 'M' && data[row][col-2] == 'A' && data[row][col-3] == 'S' {
				count += 1
			}

			if row+3 < rn {
				if col-3 >= 0 && data[row+1][col-1] == 'M' && data[row+2][col-2] == 'A' && data[row+3][col-3] == 'S' {
					count += 1
				}
				if data[row+1][col] == 'M' && data[row+2][col] == 'A' && data[row+3][col] == 'S' {
					count += 1
				}
				if col+3 < cn && data[row+1][col+1] == 'M' && data[row+2][col+2] == 'A' && data[row+3][col+3] == 'S' {
					count += 1
				}
			}
		}
	}

	for col := 0; col < cn; col++ {
		count2 += countXmas(column(data, col))
	}

	diags := diagonals(data)
	for _, diag := range diags {
		// fmt.Println(string(diag))
		count2 += countXmas(diag)
	}

	fmt.Println("The single counter:", count2)
	fmt.Println("The XMAS count:", count)
}

func column(data [][]byte, colidx int) []byte {
	dst := make([]byte, len(data))
	for i := 0; i < len(data); i++ {
		dst[i] = data[i][colidx]
	}

	return dst
}

func diagonals(data [][]byte) [][]byte {
	diags := [][]byte{}
	// upwards from [0,0] <- this is half
	for y := 0; y < len(data); y++ {
		col := 0
		row := y
		diag := []byte{}
		for row >= 0 && col < len(data[y]) {
			point := data[row][col]
			diag = append(diag, point)
			row--
			col++
		}
		diags = append(diags, diag)
	}

	// upwards from [n,0] <- the other half
	for x := 1; x < len(data); x++ {
		col := x
		row := len(data) - 1
		diag := []byte{}
		for row >= 0 && col < len(data[x]) {
			point := data[row][col]
			diag = append(diag, point)
			row--
			col++
		}
		diags = append(diags, diag)
	}

	// downwards from [n, 0]
	for y := len(data); y >= 0; y-- {
		col := 0
		row := y
		diag := []byte{}
		for row < len(data) && col < len(data[y]) {
			point := data[row][col]
			diag = append(diag, point)
			row++
			col++
		}
		diags = append(diags, diag)
	}

	for x := 1; x < len(data); x++ {
		col := x
		row := 0
		diag := []byte{}
		for row <= len(data) && col < len(data[x]) {
			point := data[row][col]
			diag = append(diag, point)
			row++
			col++
		}
		diags = append(diags, diag)
	}

	return diags
}

func countXmas(data []byte) int {
	count := 0
	for i := 0; i < len(data)-3; i++ {
		word := string(data[i : i+4])
		if word == "XMAS" || word == "SAMX" {
			count++
		}
	}

	return count
}
