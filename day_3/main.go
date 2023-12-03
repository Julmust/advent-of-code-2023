package main

import (
	readinput "aoc_2023/inputreader"
	"fmt"
	"strconv"
	"strings"
)

type coord struct {
	x, y int
}

func display(data [][]string) {
	for _, i := range data {
		fmt.Println(i)
	}
}

// Parse indata to 2d array where each sub array is one row
func parse(data []string) [][]string {
	var output [][]string
	for _, val := range data {
		output = append(output, strings.Split(val, ""))
	}

	return output
}

// Calculates boundries to look for, based on coordinates
// If a boundry would be out of range, it is set to -1 as all valid boundries
// are positive integers
func getBoundries(co []coord, w, h int) (int, int, int, int) {
	lboundry := co[0].y - 1
	if lboundry < 0 {
		lboundry = -1
	}

	rboundry := co[1].y + 1
	if rboundry > w {
		rboundry = -1
	}

	tboundry := co[0].x - 1
	if tboundry < 0 {
		tboundry = -1
	}

	bboundry := co[0].x + 1
	if bboundry > h {
		bboundry = -1
	}

	return lboundry, rboundry, tboundry, bboundry
}

func checkIfPart(coordinates [][]coord, data [][]string) []int {
	var output []int
	for _, co := range coordinates {
		var above, left, right, below string
		lb, rb, tb, bb := getBoundries(co, len(data[0])-1, len(data)-1)

		// Extract all the characters surrounding the possible part number into a string
		// For lb and rb, we ignore them if the calculated boundry is -1
		// but we still need correct values for them for the top and bottom rows
		// so we re-calculate them if they are reported as -1
		if lb != -1 {
			left = data[co[0].x][lb]
		} else {
			lb = 0
		}
		if rb != -1 {
			right = data[co[0].x][rb]
		} else {
			rb = len(data[0]) - 1
		}
		if tb != -1 {
			above = strings.Join(data[tb][lb:rb+1], "")
		}
		if bb != -1 {
			below = strings.Join(data[bb][lb:rb+1], "")
		}

		surrounding_chars := above + left + right + below

		for _, i := range surrounding_chars {
			if i != 46 { // If we find a character that's not a dot, we've found a valid part!
				res := data[co[0].x][co[0].y : co[1].y+1]
				val, _ := strconv.Atoi(strings.Join(res, ""))
				output = append(output, val)
			}
		}
	}
	return output
}

func getNumberCoordinates(data [][]string) [][]coord {
	var coords_slice [][]coord

	for idx, val := range data { // Gives []string
		var tmp_co []coord
		for sidx, sval := range val {
			_, err := strconv.Atoi(sval) // Trying to cast from string to int
			if err == nil {              // If the cast succeeded, we know the value is an integer
				tmp_co = append(tmp_co, coord{x: idx, y: sidx})
			} else {
				// If the cast failed, and we have values in our "cache"
				// we flush those to the main 2d array and clear the cache.
				if len(tmp_co) > 0 {
					coords_slice = append(coords_slice, tmp_co)
					tmp_co = []coord{}
				}
			}
			if sidx == len(val)-1 {
				if len(tmp_co) > 0 {
					coords_slice = append(coords_slice, tmp_co)
					tmp_co = []coord{}
				}
			}
		}
	}

	// Trim the output array to only contain the stop and start values
	// of the integers
	for idx, oval := range coords_slice {
		coords_slice[idx] = []coord{oval[0], oval[len(oval)-1]}
	}

	return coords_slice
}

func one(data []string) {
	parsed_data := parse(data)
	coordinates := getNumberCoordinates(parsed_data)
	parts := checkIfPart(coordinates, parsed_data)
	var res int
	for _, i := range parts {
		res += i
	}
	fmt.Printf("Result 1: %v\n", res)
}

func main() {
	// data := readinput.ReadText("example1.txt")
	data := readinput.ReadText("input.txt")

	one(data)
}
