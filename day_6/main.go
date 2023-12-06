package main

import (
	readinput "aoc_2023/inputreader"
	"fmt"
	"strconv"
	"strings"
)

// Very basic zip function, only works when slices are equal in length
func zipSlices(a, b []int) [][]int {
	var output [][]int
	for idx := range a {
		output = append(output, []int{a[idx], b[idx]})
	}

	return output
}

func parseData(data []string, task_id int) [][]int {
	time_str_slc, dist_str_slc := strings.Fields(data[0]), strings.Fields(data[1])
	time_str_slc, dist_str_slc = time_str_slc[1:], dist_str_slc[1:]

	if task_id == 2 {
		var time_str, dist_str string
		for _, a := range time_str_slc {
			time_str += a
		}
		time_str_slc = []string{time_str}

		for _, a := range dist_str_slc {
			dist_str += a
		}
		dist_str_slc = []string{dist_str}
	}

	var time_int_slc, dist_int_slc []int
	for _, i := range time_str_slc {
		val, _ := strconv.Atoi(i)
		time_int_slc = append(time_int_slc, val)
	}

	for _, i := range dist_str_slc {
		val, _ := strconv.Atoi(i)
		dist_int_slc = append(dist_int_slc, val)
	}

	return zipSlices(time_int_slc, dist_int_slc)
}

func calc(race []int, task_id int) int {
	var new_recs int
	race_dur, race_rec := race[0], race[1]

	for i := 0; i <= race_dur; i++ {
		spd := i
		dist := spd * (race_dur - spd)

		if task_id == 1 && dist > race_rec {
			new_recs += 1
		} else if task_id == 2 && dist > race_rec {
			return (race_dur - (spd * 2) + 1)
		}
	}

	return new_recs
}

func one(data []string) {
	parsedData := parseData(data, 1)

	res := 1
	for _, i := range parsedData {
		res *= calc(i, 1)
	}

	fmt.Printf("Solution 1: %d\n", res)
}

func two(data []string) {
	parsedData := parseData(data, 2)

	fmt.Printf("Solution 2: %d\n", calc(parsedData[0], 2))
}

func main() {
	// data := readinput.ReadText("example.txt")
	data := readinput.ReadText("input.txt")

	one(data)
	two(data)
}
