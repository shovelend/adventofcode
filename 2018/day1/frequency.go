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
	frequencies := strings.Split(string(dat), "\n")

	var total int
	var totalList []int
	var found bool
	for found == false {
		for _, frequency := range frequencies {

			current, err := strconv.Atoi(frequency)
			check(err)
			total += current

			if contains(totalList, total) {
				fmt.Print(total)
				return
			} else {
				totalList = append(totalList, total)
			}
		}
	}
}

func contains(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
