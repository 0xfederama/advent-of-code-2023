package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type race struct {
    time int
    dist int
}

func ways_win(r race) int {
    ways := 0
    for i := 1; i < r.time; i++ {
        // i is the time holding the button (skipping 0 and total time)
        tot_dist := i * (r.time - i)
        if tot_dist > r.dist {
            ways++
        }
    }
    return ways
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var input [][]string
    // read both lines
    for scanner.Scan() {
        line := scanner.Text()
        input = append(input, strings.Fields(line))
    }
    // build the races array for part 1 and the single race for part 2
    var races []race
    tot_strtime := ""
    tot_strdist := ""
    for i := 1; i < len(input[0]); i++ {
        strtime := input[0][i]
        strdist := input[1][i]
        tot_strtime += strtime
        tot_strdist += strdist
        time, _ := strconv.Atoi(strtime)
        dist, _ := strconv.Atoi(strdist)
        rac := race {
            time,
            dist,
        }
        races = append(races, rac)
    }
    tot_time, _ := strconv.Atoi(tot_strtime)
    tot_dist, _ := strconv.Atoi(tot_strdist)
    race2 := race {
        tot_time,
        tot_dist,
    }
    
    // find number of ways to win for part1
    tot_ways_p1 := 1
    for _, r := range races {
        tot_ways_p1 *= ways_win(r)
    }
    fmt.Println("Result 1:", tot_ways_p1)

    // find number of ways to win for part 2
    tot_ways_p2 := ways_win(race2)
    fmt.Println("Result 2:", tot_ways_p2)

}
