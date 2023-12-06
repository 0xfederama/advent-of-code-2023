package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum_p1 := 0.0
	card_i := 0
	cards := make(map[int]int, 0)
	cards[0] = 0
	tot_copies := 0

	for scanner.Scan() {
		line := scanner.Text()

		// split line and get winning numbers and my numbers
		splitted := strings.Split(strings.Split(line, ":")[1], "|")
		winnums := make(map[string]struct{}, 0)
		mynum := strings.Fields(splitted[1])
		for _, n := range strings.Fields(splitted[0]) {
			winnums[n] = struct{}{}
		}

		// compute winning numbers
		matches := 0
		for _, n := range mynum {
			if _, ok := winnums[n]; ok {
				matches++
			}
		}

		// end of part1, part2
		copies := cards[card_i] + 1 // the map only counts the duplicated ones, hence the +1
		if matches > 0 {
			points := math.Pow(2, float64(matches-1))
			sum_p1 += points
			for i := 0; i < matches; i++ {
				cards[card_i+i+1] += copies
			}
		}
		tot_copies += copies
		card_i++
	}

	fmt.Println("Result 1:", sum_p1)
	fmt.Println("Result 2:", tot_copies)

}
