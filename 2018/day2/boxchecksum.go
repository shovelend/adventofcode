package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("input.txt")
	check(err)
	boxIDs := strings.Split(string(dat), "\n")
	a(boxIDs)
	b(boxIDs)
}

func a(boxIDs []string) {
	var twos int
	var threes int
	var checksum int

	for _, boxID := range boxIDs {
		two := doesContainXOfTheSameLetters(boxID, 2)
		three := doesContainXOfTheSameLetters(boxID, 3)

		if two {
			twos++
		}
		if three {
			threes++
		}
	}
	checksum = twos * threes
	fmt.Println(checksum)
}

func b(boxIDs []string) {
	for index1, boxID1 := range boxIDs {
		for index2, boxID2 := range boxIDs {
			if index2 > index1 {
				isOneLetterDiff(boxID1, boxID2)
			}
		}
	}
}

func doesContainXOfTheSameLetters(id string, nOfLetters int) bool {
	symbols := frequentSym(id)
	for _, v := range symbols {
		if v == nOfLetters {
			return true
		}
	}
	return false
}

func frequentSym(id string) map[string]int {
	symbols := make(map[string]int)

	for _, sym := range id {
		char := strconv.QuoteRune(sym)
		count := symbols[char]

		if count == symbols[char] {
			symbols[char] = count + 1
		} else {
			symbols[char] = 1
		}
	}

	return symbols
}

func isOneLetterDiff(boxID1 string, boxID2 string) bool {
	lenval := len(boxID1)
	var commonLetters strings.Builder
	for index1, sym := range boxID1 {
		for index2, sym2 := range boxID2 {
			if index1 == index2 {
				char1 := strconv.QuoteRune(sym)
				char2 := strconv.QuoteRune(sym2)
				if char1 == char2 {
					lenval--
					commonLetters.WriteString(char1)
				}
			}
		}
	}
	if lenval == 1 {
		fmt.Println(commonLetters.String())
	}
	return false
}
