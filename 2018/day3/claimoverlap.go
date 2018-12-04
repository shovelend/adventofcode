package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Claim struct {
	id     string
	startX int
	startY int
	width  int
	length int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("input.txt")
	check(err)
	claimsString := strings.Split(string(dat), "\n")
	var claimList []Claim

	for _, claim := range claimsString {
		claimList = append(claimList, constructClaim(claim))
	}

	overlappedInches := getNumberOfOverLappedInches(getFilledFabric(claimList))
	fmt.Println(overlappedInches)
}

func constructClaim(claimString string) Claim {
	var claim Claim
	claim.id = claimString[1 : strings.Index(claimString, "@")-1]
	claim.startX, _ = strconv.Atoi(claimString[strings.Index(claimString, "@")+2 : strings.Index(claimString, ",")])
	claim.startY, _ = strconv.Atoi(claimString[strings.Index(claimString, ",")+1 : strings.Index(claimString, ":")])
	claim.width, _ = strconv.Atoi(claimString[strings.Index(claimString, ":")+2 : strings.Index(claimString, "x")])
	claim.length, _ = strconv.Atoi(claimString[strings.Index(claimString, "x")+1 : len(claimString)])

	return claim
}

func getFilledFabric(claimList []Claim) [1000][1000]string {
	var fabric [1000][1000]string
	for _, claim := range claimList {
		var x, y int

		for y = claim.startY; y < claim.startY+claim.length; y++ {
			for x = claim.startX; x < claim.startX+claim.width; x++ {
				if fabric[x][y] != "" || fabric[x][y] == "X" {
					fabric[x][y] = "X"
				} else {
					fabric[x][y] = claim.id
				}
			}
		}
	}
	return fabric
}

func getNumberOfOverLappedInches(fabric [1000][1000]string) int {
	var count, x, y int
	for y = 0; y < 1000; y++ {
		for x = 0; x < 1000; x++ {
			if fabric[x][y] == "X" {
				count++
			}
		}
	}
	return count
}
