package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type hand struct {
	cards string
	bid   int
}

var rank1 = map[byte]int{
	'A': 1,
	'K': 2,
	'Q': 3,
	'J': 4,
	'T': 5,
	'9': 6,
	'8': 7,
	'7': 8,
	'6': 9,
	'5': 10,
	'4': 11,
	'3': 12,
	'2': 13,
}

var rank2 = map[byte]int{
	'A': 1,
	'K': 2,
	'Q': 3,
	'T': 4,
	'9': 5,
	'8': 6,
	'7': 7,
	'6': 8,
	'5': 9,
	'4': 10,
	'3': 11,
	'2': 12,
	'J': 13,
}

func occ_to_str(strength int, occ int) int {
	switch occ {
	case 2:
		switch strength {
		case 7:
			return 6
		case 6:
			return 5
		case 4:
			return 3
		}
	case 3:
		switch strength {
		case 7:
			return 4
		case 6:
			return 3
		}
	case 4:
		return 2
	case 5:
		return 1
	}
	return 0
}

func hand_strength_p1(h string) int {
	// 1 five kind, 2 four kind, 3 full, 4 three kind, 5 two pair, 6 one pair, 7 high card
	cards := make(map[rune]int)
	for _, c := range h {
		cards[c]++
	}
	strength := 7
	for _, occ := range cards {
		str := occ_to_str(strength, occ)
		if str != 0 {
			strength = str
		}
	}
	return strength
}

func hand_strength_p2(h string) int {
	// 1 five kind, 2 four kind, 3 full, 4 three kind, 5 two pair, 6 one pair, 7 high card
	cards := make(map[rune]int)
	for _, c := range h {
		cards[c]++
	}
	strength := 7
    if cards['J'] == 5 {
        return 1
    }
	for c, occ := range cards {
		if c == 'J' {
			continue
		}
		str := occ_to_str(strength, occ)
		if str != 0 {
			strength = str
		}
	}
	// the strength doesn't consider jokers, so if e.g. i have a full, it means i have also 0 jokers
	switch strength {
	case 2:
		if cards['J'] == 1 {
			return 1
		}
	case 4:
		switch cards['J'] {
		case 1:
			return 2
		case 2:
			return 1
		}
	case 5:
		if cards['J'] == 1 {
            return 3
		}
    case 6:
        switch cards['J'] {
        case 1:
            return 4
        case 2:
            return 2
        case 3:
            return 1
        }
    case 7:
        switch cards['J'] {
        case 1:
            return 6
        case 2:
            return 4
        case 3:
            return 2
        case 4:
            return 1
        }
	}
	return strength
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hands []hand
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		bid, _ := strconv.Atoi(line[1])
		h := hand{
			line[0],
			bid,
		}
		hands = append(hands, h)
	}

	// sort hands for part1
	sort.Slice(hands, func(i, j int) bool {
		h1 := hands[i].cards
		h2 := hands[j].cards
		s1 := hand_strength_p1(h1)
		s2 := hand_strength_p1(h2)
		if s1 == s2 {
			for i := 0; i < len(h1); i++ {
				r1 := rank1[h1[i]]
				r2 := rank1[h2[i]]
				if r1 == r2 {
					continue
				}
				return r1 < r2
			}
		}
		return s1 < s2
	})
	// find the total win with bids
	numhands := len(hands)
	tot_win := 0
	for i, h := range hands {
		tot_win += (h.bid * (numhands - i))
	}
	fmt.Println("Result 1:", tot_win)

	// sort hands for part2
	sort.Slice(hands, func(i, j int) bool {
		h1 := hands[i].cards
		h2 := hands[j].cards
		s1 := hand_strength_p2(h1)
		s2 := hand_strength_p2(h2)
		if s1 == s2 {
			for i := 0; i < len(h1); i++ {
				r1 := rank2[h1[i]]
				r2 := rank2[h2[i]]
				if r1 == r2 {
					continue
				}
				return r1 < r2
			}
		}
		return s1 < s2
	})
	// find the total win with bids
	tot_win = 0
	for i, h := range hands {
		tot_win += (h.bid * (numhands - i))
	}
	fmt.Println("Result 2:", tot_win)

}
