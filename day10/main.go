package main

import (
	"bufio"
	"fmt"
	"os"
)

func starting_positions(matrix [][]rune, start_i, start_j int) []rune {
	var goto_pos []rune
	// check north
	if start_i != 0 {
		ni := start_i - 1
		nj := start_j
		north := matrix[ni][nj]
		if north == '|' || north == 'F' || north == '7' {
			goto_pos = append(goto_pos, 'N')
		}
	}
	// check south
	if start_i != len(matrix)-1 {
		si := start_i + 1
		sj := start_j
		south := matrix[si][sj]
		if south == '|' || south == 'L' || south == 'J' {
			goto_pos = append(goto_pos, 'S')
		}
	}
	// check west
	if start_j != 0 {
		wi := start_i
		wj := start_j - 1
		west := matrix[wi][wj]
		if west == '-' || west == 'L' || west == 'F' {
			goto_pos = append(goto_pos, 'W')
		}
	}
	// check east
	if start_j != len(matrix[0])-1 {
		ei := start_i
		ej := start_j + 1
		east := matrix[ei][ej]
		if east == '-' || east == '7' || east == 'J' {
			goto_pos = append(goto_pos, 'E')
		}
	}
	return goto_pos
}

func opp_dir(dir rune) rune {
	switch dir {
	case 'N':
		return 'S'
	case 'S':
		return 'N'
	case 'E':
		return 'W'
	case 'W':
		return 'E'
	}
	return 'X'
}

func loop_len_dir(matrix [][]rune, start_i, start_j int, dir rune) int {
	curr_i := start_i
	curr_j := start_j
	switch dir {
	case 'N':
		curr_i--
	case 'S':
		curr_i++
	case 'E':
		curr_j++
	case 'W':
		curr_i--
	}
	curr_pipe := matrix[curr_i][curr_j]
	dist := 1
	coming_from := opp_dir(dir)
	fmt.Println("Starting from", curr_i, curr_j, "coming from", string(coming_from))
	for curr_pipe != 'S' {
		fmt.Println("Curr pipe is", string(curr_pipe))
		switch curr_pipe {
		case '-':
			if coming_from == 'E' && curr_j != 0 {
				curr_j--
			} else if coming_from == 'W' && curr_j != len(matrix[0])-1 {
				curr_j++
				fmt.Println("WALL on", string(curr_pipe), "coming from", string(coming_from))
			} else {
				return 0
			}
		case '|':
			if coming_from == 'N' && curr_i != len(matrix)-1 {
				curr_i++
			} else if coming_from == 'S' && curr_i != 0 {
				curr_i--
			} else {
				fmt.Println("WALL on", string(curr_pipe), "coming from", string(coming_from))
				return 0
			}
		case 'F':
			if coming_from == 'S' && curr_j != len(matrix[0])-1 {
				curr_j++
				coming_from = 'W'
			} else if coming_from == 'E' && curr_i != len(matrix)-1 {
				curr_i++
				coming_from = 'N'
			} else {
				fmt.Println("WALL on", string(curr_pipe), "coming from", string(coming_from))
				return 0
			}
		case '7':
			if coming_from == 'S' && curr_j != 0 {
				curr_j--
				coming_from = 'E'
			} else if coming_from == 'W' && curr_i != len(matrix)-1 {
				curr_i++
				coming_from = 'N'
			} else {
				fmt.Println("WALL on", string(curr_pipe), "coming from", string(coming_from))
				return 0
			}
		case 'J':
			if coming_from == 'N' && curr_j != 0 {
				curr_j--
				coming_from = 'E'
			} else if coming_from == 'W' && curr_i != 0 {
				curr_i--
				coming_from = 'S'
			} else {
				fmt.Println("WALL on", string(curr_pipe), "coming from", string(coming_from))
				return 0
			}
		case 'L':
			if coming_from == 'E' && curr_i != 0 {
				curr_i--
				coming_from = 'S'
			} else if coming_from == 'N' && curr_j != len(matrix[0])-1 {
				curr_j++
				coming_from = 'W'
			} else {
				fmt.Println("WALL on", string(curr_pipe), "coming from", string(coming_from))
				return 0
			}
		case '.':
			fmt.Println("WALL on", string(curr_pipe), "coming from", string(coming_from))
			return 0
		}
		curr_pipe = matrix[curr_i][curr_j]
		dist++
	}
	fmt.Println("Gone back to S with dist", dist)

	return dist
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var matrix [][]rune
	start_i := 0
	start_j := 0
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []rune(line))
		for j, r := range line {
			if r == 'S' {
				start_i = i
				start_j = j
			}
		}
		i++
	}
	fmt.Println("S is", start_i, start_j)

	start_dir := starting_positions(matrix, start_i, start_j)
	fmt.Println("From S start on directions", string(start_dir))

	// each pipe goes on to one and only one pipe, there are no disjunctions
	// it should be easy to do with a single loop while we don't return to S
	maxdist := 0
	for _, dir := range start_dir {
		maxdist = max(maxdist, loop_len_dir(matrix, start_i, start_j, dir))
	}

	fmt.Println("Result 1:", maxdist/2)

}
