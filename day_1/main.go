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
				// fmt.Printf("%v %v %T\n", line[idx], string(line[idx]), line[idx])
			}
		}
		res_string := tmp[0] + tmp[len(tmp)-1]
		res_int, _ := strconv.Atoi(res_string)
		output = append(output, res_int)
	}

	return output
}

func main() {

	data := readinput.ReadText("testdata.txt")
	// data := readinput.ReadText("small_one.txt")
	output := parseInput(data)

	var sum int
	for _, val := range output {
		sum += val
	}

	fmt.Printf("Solution 1: %v\n", sum)
}
