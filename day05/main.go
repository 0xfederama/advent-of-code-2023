package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type converter struct {
	dst int
	src int
	len int
}

func str_to_int_arr(line string) []int {
	var arr []int
	seedsstr := strings.Fields(line)
	for _, s := range seedsstr {
		a, _ := strconv.Atoi(s)
		arr = append(arr, a)
	}
	return arr
}

func extract_range(line string) converter {
    ranges := str_to_int_arr(line)
    rng := converter {
        ranges[0],
        ranges[1],
        ranges[2],
    }
    return rng
}

func seed_to_loc(seed int, almanac [][]converter) int {
    val := seed
    for _, op := range almanac {
        for _, r := range op {
            if val >= r.src && val < r.src + r.len {
                // found the right spot
                val = r.dst + (val - r.src)
                break
            }
        }
        // if i'm here there is not a predefined range for the seed,
        //  so don't modify the value
    }
    return val
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	firstline := true
	operation := -1 // 0 seed_to_soil, 1 soil_to_fert, 2 fert_to_h20, 3 h2o_to_light,
	// 4 light_to_temp, 5 temp_to_hum, 6 hum_to_loc
	var seeds []int
	almanac := make([][]converter, 7)
	for scanner.Scan() {
		line := scanner.Text()
		if firstline {
            seeds = str_to_int_arr(line[6:])
			firstline = false
		} else {
			if line == "" {
				operation++
				continue
			}
			if !unicode.IsDigit(rune(line[0])) {
				continue
			}
			// add the ranges to almanac[operation]
            rng := extract_range(line)
            almanac[operation] = append(almanac[operation], rng)
		}
	}

    // PART 1
	minloc := math.MaxInt
	for _, seed := range seeds {
		// find location of seed
		loc := seed_to_loc(seed, almanac)
		minloc = min(minloc, loc)
	}
	fmt.Println("Result 1:", minloc)

    // PART 2
    minloc = math.MaxInt
    for i := 0; i < len(seeds); i+=2 {
        seed := seeds[i]
        rangelen := seeds[i+1]
        for j := seed; j < rangelen + seed; j++ {
            loc := seed_to_loc(j, almanac)
            minloc = min(minloc, loc)
        }
    }
    fmt.Println("Result 2:", minloc)

}
