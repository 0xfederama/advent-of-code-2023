package main

import (
	"bufio"
	"fmt"
	"os"
)

type coord struct {
	i int
	j int
}

func starting_directions(matrix [][]rune, start_i, start_j int) []rune {
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

func loop_len_dir(matrix [][]rune, start_i, start_j int, dir rune) (int, []coord) {
	loop := make([]coord, 0)
	i := start_i
	j := start_j
	switch dir {
	case 'N':
		i--
	case 'S':
		i++
	case 'E':
		j++
	case 'W':
		j--
	}
	curr_pipe := matrix[i][j]
	dist := 1
	coming_from := opp_dir(dir)
	for curr_pipe != 'S' {
		loop = append(loop, coord{i, j})
		switch curr_pipe {
		case '-':
			if coming_from == 'E' && j != 0 {
				j--
			} else if coming_from == 'W' && j != len(matrix[0])-1 {
				j++
			} else {
				fmt.Println("Wall on", string(curr_pipe), "coming from", string(coming_from), "dist", dist)
				return 0, loop
			}
		case '|':
			if coming_from == 'N' && i != len(matrix)-1 {
				i++
			} else if coming_from == 'S' && i != 0 {
				i--
			} else {
				fmt.Println("Wall on", string(curr_pipe), "coming from", string(coming_from), "dist", dist)
				return 0, loop
			}
		case 'F':
			if coming_from == 'S' && j != len(matrix[0])-1 {
				j++
				coming_from = 'W'
			} else if coming_from == 'E' && i != len(matrix)-1 {
				i++
				coming_from = 'N'
			} else {
				fmt.Println("Wall on", string(curr_pipe), "coming from", string(coming_from), "dist", dist)
				return 0, loop
			}
		case '7':
			if coming_from == 'S' && j != 0 {
				j--
				coming_from = 'E'
			} else if coming_from == 'W' && i != len(matrix)-1 {
				i++
				coming_from = 'N'
			} else {
				fmt.Println("Wall on", string(curr_pipe), "coming from", string(coming_from), "dist", dist)
				return 0, loop
			}
		case 'J':
			if coming_from == 'N' && j != 0 {
				j--
				coming_from = 'E'
			} else if coming_from == 'W' && i != 0 {
				i--
				coming_from = 'S'
			} else {
				fmt.Println("Wall on", string(curr_pipe), "coming from", string(coming_from), "dist", dist)
				return 0, loop
			}
		case 'L':
			if coming_from == 'E' && i != 0 {
				i--
				coming_from = 'S'
			} else if coming_from == 'N' && j != len(matrix[0])-1 {
				j++
				coming_from = 'W'
			} else {
				fmt.Println("Wall on", string(curr_pipe), "coming from", string(coming_from), "dist", dist)
				return 0, loop
			}
		case '.':
			fmt.Println("Wall on", string(curr_pipe), "coming from", string(coming_from), "dist", dist)
			return 0, loop
		}
		curr_pipe = matrix[i][j]
		dist++
	}
	loop = append(loop, coord{i, j})
	fmt.Println("Gone back to S with dist", dist)

	return dist, loop
}

func print_matrix(mat [][]rune) {
	fmt.Println("Matrix:")
	for _, line := range mat {
		fmt.Print("[ ")
		for _, r := range line {
			fmt.Print(string(r))
		}
		fmt.Println("]")
	}
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

	if len(matrix) < 32 {
		// don't print if too big
		print_matrix(matrix)
	}
	start_dir := starting_directions(matrix, start_i, start_j)
	fmt.Println("S is at", start_i, start_j, "and can go to:", string(start_dir))

	// each pipe goes on to one and only one pipe, there are no disjunctions
	// it should be easy to do with a single loop while we don't return to S
	maxdist := 0
	loop := make([]coord, 0)
	for _, dir := range start_dir {
		fmt.Println("Starting on direction", string(dir))
		dist, lp := loop_len_dir(matrix, start_i, start_j, dir)
		if dist > maxdist {
			maxdist = dist
			loop = lp
		}
	}
	fmt.Println("Result 1:", maxdist/2)
	fmt.Println("Loop", loop)

    // find enclosed with shoelace formula and pick's theorem
	area := 0
	for i := 0; i < len(loop); i++ {
		curr := loop[i]
		next := loop[(i+1)%len(loop)]

		area += curr.i*next.j - curr.j*next.i
	}
	if area < 0 {
		area = -area
	}
	area /= 2
    enclosed := area - len(loop)/2 + 1

	fmt.Println("Result 2:", enclosed)

}
