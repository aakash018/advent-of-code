package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func removeSliceWithIndex[t any](slice []t, index int) []t {
	new_slice := make([]t, 0)
	new_slice = append(new_slice, slice...)
	if index+1 > len(slice) {
		empty_slice := []t{}
		new_slice = append(new_slice[:index], empty_slice...)
	} else {

		new_slice = append(new_slice[:index], new_slice[index+1:]...)
	}

	return new_slice

}

func is_safe(line []string) bool {
	var pole int
	var next_pole int

	realFirst, _ := strconv.Atoi(line[0])
	realSecond, _ := strconv.Atoi(line[1])

	mainDiff := realSecond - realFirst
	if mainDiff > 0 {
		pole = 1
	} else {
		pole = 0
	}

	for i := 0; i < len(line)-1; i++ {
		first, _ := strconv.Atoi(line[i])
		second, _ := strconv.Atoi(line[i+1])
		diff := second - first
		// fmt.Println(diff, first, second, line)
		if diff > 0 {
			next_pole = 1
		} else if diff < 0 {
			next_pole = 0
		} else {
			next_pole = pole
		}
		if absInt(diff) < 1 || absInt(diff) > 3 || (next_pole != pole) {
			return false
		}

	}
	return true
}

// 343
func main() {
	inputBlob, _ := os.ReadFile("./input.txt")
	input := strings.TrimSpace(string(inputBlob))

	lines := strings.Split(input, "\n")
	safe_count := 0
	for _, line := range lines {
		elements := strings.Split(line, " ")
		if is_safe(elements) {
			safe_count++
		} else {
			for i := 0; i < len(elements); i++ {
				temp_arr := removeSliceWithIndex(elements, i)
				if is_safe(temp_arr) {
					safe_count++
					break
				}
			}
		}
	}

	fmt.Println("count", safe_count)

}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
