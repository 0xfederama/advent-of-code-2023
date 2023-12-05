package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func get_adj(puzzle [][]rune, i int, j int) int {
	if i < 0 || j < 0 || i >= len(puzzle) || j >= len(puzzle[0]) {
		return 0
	}
	if unicode.IsDigit(puzzle[i][j]) {
		var currnum []rune
		for j >= 1 && unicode.IsDigit(puzzle[i][j-1]) {
			j--
		}
		for j < len(puzzle[0]) && unicode.IsDigit(puzzle[i][j]) {
			currnum = append(currnum, puzzle[i][j])
			puzzle[i][j] = '.'
			j++
		}
		num, _ := strconv.Atoi(string(currnum))
		return num
	}
	return 0
}

func print_p(puzzle [][]rune) {
	fmt.Println("PUZZLE with length", len(puzzle), len(puzzle[0]))
	for i, row := range puzzle {
		for _, c := range row {
			fmt.Print(string(c))
		}
		if i < len(puzzle)-1 {
			fmt.Println()
		}
	}
}

func count_diff_zero(arr []int) int {
	tot := 0
	for _, n := range arr {
		if n != 0 {
			tot += 1
		}
	}
	return tot
}

func mult_diff_zero(arr []int) int {
	mult := 1
	for _, n := range arr {
		if n != 0 {
			mult *= n
		}
	}
	return mult
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var puzzle [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		puzzle = append(puzzle, []rune(line))
	}
	tot_sum := 0
	tot_gear_rat := 0
	for i, row := range puzzle {
		for j, c := range row {
			if !unicode.IsDigit(c) && c != '.' {
				// search adjacent
				var adjacents []int
				adjacents = append(adjacents, get_adj(puzzle, i-1, j-1))
				adjacents = append(adjacents, get_adj(puzzle, i-1, j))
				adjacents = append(adjacents, get_adj(puzzle, i-1, j+1))
				adjacents = append(adjacents, get_adj(puzzle, i, j-1))
				adjacents = append(adjacents, get_adj(puzzle, i, j+1))
				adjacents = append(adjacents, get_adj(puzzle, i+1, j-1))
				adjacents = append(adjacents, get_adj(puzzle, i+1, j))
				adjacents = append(adjacents, get_adj(puzzle, i+1, j+1))
				// part1
				for _, n := range adjacents {
					tot_sum += n
				}
				// part2
				if c == '*' {
					if count_diff_zero(adjacents) == 2 {
						tot_gear_rat += mult_diff_zero(adjacents)
					}
				}
			}
		}
	}
	fmt.Println("Result 1:", tot_sum)
	fmt.Println("Result 2:", tot_gear_rat)
}
