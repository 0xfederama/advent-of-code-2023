package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func part1(line string, re *regexp.Regexp) int {
	list := re.FindAllString(line, -1)
	str := list[0] + list[len(list)-1]
	number, _ := strconv.Atoi(str)

	return number
}

func convert(s string) string {
	switch s {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return s
	}
}

func part2(line string, re *regexp.Regexp) int {
	// findallstring doesn't work for overlapping words like oneight
	var list []string
	for i := range line {
		word := re.FindString(line[i:])
		if word != "" {
			list = append(list, word)
		}
	}

	str := convert(list[0]) + convert(list[len(list)-1])
	number, _ := strconv.Atoi(str)
	return number
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum1 := 0
	sum2 := 0
	reg1 := regexp.MustCompile("\\d")
	reg2 := regexp.MustCompile("\\d|one|two|three|four|five|six|seven|eight|nine")

	for scanner.Scan() {
		line := scanner.Text()
		res1 := part1(line, reg1)
		res2 := part2(line, reg2)
		sum1 += res1
		sum2 += res2
	}

	fmt.Println("Result 1:", sum1)
	fmt.Println("Result 2:", sum2)

}
