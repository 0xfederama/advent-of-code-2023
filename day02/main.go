package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func is_game_good_p1(game [][]string) (bool) {
	red := 12
	green := 13
	blue := 14

	for _, d := range game {
		i := 0
		for i < len(d) {
			switch d[i+1] {
			case "red":
				reddrawn, _ := strconv.Atoi(d[i])
				if reddrawn > red {
					return false
				}
			case "green":
				greendrawn, _ := strconv.Atoi(d[i])
				if greendrawn > green {
					return false
				}
			case "blue":
				bluedrawn, _ := strconv.Atoi(d[i])
				if bluedrawn > blue {
					return false
				}
			}
			i += 2
		}
	}

	return true
}

func max_cubes_p2(game [][]string) int {
	maxred := -1
	maxgreen := -1
	maxblue := -1

	for _, d := range game {
		i := 0
		for i < len(d) {
			switch d[i+1] {
			case "red":
				reddrawn, _ := strconv.Atoi(d[i])
				if reddrawn > maxred {
					maxred = reddrawn
				}
			case "green":
				greendrawn, _ := strconv.Atoi(d[i])
				if greendrawn > maxgreen {
					maxgreen = greendrawn
				}
			case "blue":
				bluedrawn, _ := strconv.Atoi(d[i])
				if bluedrawn > maxblue {
					maxblue = bluedrawn
				}
			}
			i += 2
		}
	}

	return maxred * maxgreen * maxblue
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	gameid := 1
	idtotal := 0
	maxtotal := 0
	for scanner.Scan() {
		line := scanner.Text()
		var game [][]string
		// split line and send it to helper
		draws := strings.Split(strings.Split(line, ":")[1], ";")
		for _, d := range draws {
			// each draw containes something like "6 red, 1 blue, 3 green"
			// remove commas and split by spaces
			d = strings.TrimSpace(strings.ReplaceAll(d, ",", ""))
			game = append(game, strings.Split(d, " "))
		}

		if is_game_good_p1(game) {
			idtotal += gameid
		}

		maxtotal += max_cubes_p2(game)

		gameid++
	}

	fmt.Println("Result 1:", idtotal)
	fmt.Println("Result 1:", maxtotal)

}
