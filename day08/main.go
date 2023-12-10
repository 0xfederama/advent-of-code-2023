package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type coord struct {
	left  string
	right string
}

func lcm(nums []int) int {
	lowmult := 1

	for _, n := range nums {
		lowmult = (lowmult * n) / gcd(lowmult, n)
	}

	return lowmult
}

func gcd(a int, b int) int {
	for b > 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func steps_to_z(instr string, network map[string]coord, start string, end_cond func(string) bool) int {
	steps := 0
	curr_place := start
	for !end_cond(curr_place) {
		for _, c := range instr {
			if end_cond(curr_place) {
				break
			}
			if c == 'L' {
				curr_place = network[curr_place].left
				steps++
			} else if c == 'R' {
				curr_place = network[curr_place].right
				steps++
			}
		}
	}
	return steps
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	linenum := 0
	var instr string
    
    // read input and build the network map
	network := make(map[string]coord)
	for scanner.Scan() {
		line := scanner.Text()
		if linenum == 0 {
			instr = line
			linenum++
			continue
		}
		if linenum > 1 {
			re := regexp.MustCompile("[A-Z0-9]+")
			coords := re.FindAllString(line, -1)
			network[coords[0]] = coord{
				coords[1],
				coords[2],
			}
		}
		linenum++
	}

	// part 1
	steps := steps_to_z(instr, network, "AAA", func(s string) bool {
		return s == "ZZZ"
	})
	fmt.Println("Result 1:", steps)

	// part 2
	var start_places []string
    var places_steps []int
	for key := range network {
		if strings.HasSuffix(key, "A") {
			start_places = append(start_places, key)
		}
	}
    // compute steps to reach __Z for each starting position
	for _, p := range start_places {
		places_steps = append(places_steps, steps_to_z(instr, network, p,
			func(s string) bool {
				return strings.HasSuffix(s, "Z")
			}))
	}
	// find lcm of places_steps
	tot_steps := lcm(places_steps)
	fmt.Println("Result 2:", tot_steps)

}
