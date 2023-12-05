package main

import (
	readinput "aoc_2023/inputreader"
	"fmt"
	"strconv"
	"strings"
)

func parseRawToMaps(data []string) map[string][][]int {
	var newData []string

	// Remove empty rows from indata
	for _, row := range data {
		if row != "" {
			newData = append(newData, row)
		}
	}

	fullMap := make(map[string][][]int)
	var key string

	for _, row := range newData {
		if strings.HasSuffix(row, " map:") {
			key = strings.Split(row, " map:")[0]
			fullMap[key] = nil
		} else {
			var tmpHold []int
			tmp := strings.Split(row, " ")
			for _, i := range tmp {
				val, _ := strconv.Atoi(i)
				tmpHold = append(tmpHold, val)
			}
			slc := fullMap[key]
			slc = append(slc, tmpHold)
			fullMap[key] = slc
		}
	}

	return fullMap
}

func parseSeeds(data string) []int {
	var intSeeds []int
	strSeeds := strings.Split(strings.Split(data, ": ")[1], " ")

	for _, ss := range strSeeds {
		val, _ := strconv.Atoi(ss)
		intSeeds = append(intSeeds, val)
	}

	return intSeeds
}

func calcNextSeed(seed int, mk string, maps map[string][][]int) int {
	for _, i := range maps[mk] {
		if seed >= i[1] && seed < (i[1]+i[2]) {
			// if mk == "fertilizer-to-water" {
			// 	fmt.Println(seed, i, seed+calcNextSeedAbs((i[0]-i[1])), seed+(i[0]-i[1]))
			// }
			return seed + (i[0] - i[1])
			// return seed + calcNextSeedAbs((i[0] - i[1]))
		}
	}
	return seed
}

func getLocations(seeds []int, maps map[string][][]int) []int {
	var locs []int
	order := []string{
		"seed-to-soil",
		"soil-to-fertilizer",
		"fertilizer-to-water",
		"water-to-light",
		"light-to-temperature",
		"temperature-to-humidity",
		"humidity-to-location",
	}

	for _, seed := range seeds {
		// fmt.Println("seed", seed)
		for _, o := range order {
			seed = calcNextSeed(seed, o, maps)
			// fmt.Println(o, seed)
		}
		locs = append(locs, seed)
		// fmt.Println("===========")
	}

	return locs
}

func getLocationsTwo(seedStart, seedStop int, maps map[string][][]int) int {
	// var locs []int
	var low int
	order := []string{
		"seed-to-soil",
		"soil-to-fertilizer",
		"fertilizer-to-water",
		"water-to-light",
		"light-to-temperature",
		"temperature-to-humidity",
		"humidity-to-location",
	}

	stopAt := seedStart + seedStop

	for startAt := seedStart; startAt < stopAt; startAt++ {
		seed := startAt
		for _, o := range order {
			seed = calcNextSeed(seed, o, maps)
		}

		if low == 0 || seed < low {
			low = seed
			fmt.Println(low)
		}
	}

	fmt.Println("returning")
	return low
}

func getLow(nums []int) int {
	var low int
	for _, i := range nums {
		if low == 0 || i < low {
			low = i
		}
	}
	return low
}

func one(data []string) {
	seeds := parseSeeds(data[0]) // Seeds is always on the first line
	// key: [[dst_range_start, src_range_start, range_len], ...]
	maps := parseRawToMaps(data[1:])

	locs := getLocations(seeds, maps)

	var low int
	for _, i := range locs {
		if low == 0 || i < low {
			low = i
		}
	}
	fmt.Printf("Solution 1: %d\n", low)
}

// SLOW AS A MOTHERFUCKER BUT IT RUNS
// There's supposedly a fancy way of doing this but I have no clue how so bruteforce it is
func two(data []string) {
	seeds := parseSeeds(data[0])
	maps := parseRawToMaps(data[1:])

	var actualSeeds [][]int
	for {
		if len(seeds) == 0 {
			break
		}
		actualSeeds = append(actualSeeds, seeds[:2])
		seeds = seeds[2:]
	}

	var res []int
	for _, i := range actualSeeds {
		res = append(res, getLocationsTwo(i[0], i[1], maps))
	}

	var low int
	for _, i := range res {
		if low == 0 || i < low {
			low = i
		}
	}
	fmt.Printf("Solution 2: %d\n", low)
}

func main() {
	// data := readinput.ReadText("example.txt")
	data := readinput.ReadText("input.txt")

	one(data)
	two(data)
}
