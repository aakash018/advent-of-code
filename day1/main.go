package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	inputText := strings.TrimSpace(string(input))

	pairs := strings.Split(inputText, "\n")

	var grp1 []int
	var grp2 []int

	for _, pair := range pairs {
		first := strings.Split(pair, "   ")[0]
		second := strings.Split(pair, "   ")[1]
		firstInt, _ := strconv.Atoi(first)
		secondInt, _ := strconv.Atoi(second)
		grp1 = append(grp1, firstInt)
		grp2 = append(grp2, secondInt)
	}

	sort.Ints(grp1)
	sort.Ints(grp2)
	sum := 0
	for i := 0; i < len(grp1); i++ {
		prod := grp1[i] - grp2[i]
		sum = sum + absInt(prod)
	}

	fmt.Println("Part 1 ", sum)

	part2sum := 0
	for _, val := range grp1 {
		totalCount := searchAndCount(val, grp2)
		product := val * totalCount

		part2sum = part2sum + product
	}

	fmt.Println("Part 2", part2sum)
}

func searchAndCount(number int, arr []int) int {
	count := 0
	for _, value := range arr {
		if value == number {
			count++
		}
	}
	return count
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
