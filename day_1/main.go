package main

import (
	readinput "aoc_2023/inputreader"
	"fmt"
	"strconv"
)

func parseInput(data []string) []int {
	var output []int

	for _, line := range data {
		var tmp []string
		for idx, char := range line {
			if char <= 57 && char >= 48 { // filter on ascii value for digits
				tmp = append(tmp, string(line[idx]))
			}
		}
		res_string := tmp[0] + tmp[len(tmp)-1]
		res_int, _ := strconv.Atoi(res_string)
		output = append(output, res_int)
	}

	return output
}

func second(data []string) []int {
	var output []int
	string_to_val := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for _, line := range data {
		var tmp []string
		for idx, char := range line {
			if char <= 57 && char >= 48 { // filter on ascii value for digits
				tmp = append(tmp, string(line[idx]))
			} else {
				for key := range string_to_val {
					if len(line[idx:]) >= len(key) { // Checking that there's enough characters left for this case
						if line[idx:idx+len(key)] == key {
							tmp = append(tmp, string_to_val[key])
						}
					}
				}
			}
		}
		res_string := tmp[0] + tmp[len(tmp)-1]
		res_int, _ := strconv.Atoi(res_string)
		output = append(output, res_int)
	}

	return output
}

func sumIntSlice(data []int) int {
	var sum int
	for _, val := range data {
		sum += val
	}

	return sum
}

func one(data []string) {
	output := parseInput(data)

	fmt.Printf("Solution 1: %v\n", sumIntSlice(output))
}

func two(data []string) {
	output := second(data)

	fmt.Printf("Solution 2: %v\n", sumIntSlice(output))
}

func main() {

	data := readinput.ReadText("testdata.txt")
	// data := readinput.ReadText("small_two.txt")
	one(data)
	two(data)
}
