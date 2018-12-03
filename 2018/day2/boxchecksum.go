package main

import (
	"fmt"
	"io/ioutil"
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
	boxIds := strings.Split(string(dat),"\n")
	var twos int
	var threes int
	var checksum int

	for _, boxId := range boxIds{
		two := DoesContainXOfTheSameLetters(boxId, 2)
		three := DoesContainXOfTheSameLetters(boxId, 3)

		if(two){ twos+=1}
		if(three){ threes+=1}
	}
	checksum = twos * threes
	fmt.Println(checksum)
}

func DoesContainXOfTheSameLetters(id string, noLetters int) bool {
	actualNoLetters := strings.Count(id, "a")
	return actualNoLetters == noLetters
}

func Contains(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}