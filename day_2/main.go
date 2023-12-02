package main

import (
	readinput "aoc_2023/inputreader"
	"fmt"
	"strconv"
	"strings"
)

type game struct {
	red, green, blue int
}

type full_game struct {
	game_id          int
	individual_games []*game
	valid_game       bool
}

// Validates cubes against the known maxes
func (g *game) validate() bool {
	max_red := 12
	max_green := 13
	max_blue := 14

	if g.red > max_red || g.green > max_green || g.blue > max_blue {
		return false
	}

	return true
}

// Run validation on each round
func (fg *full_game) validate() bool {
	for _, a := range fg.individual_games {
		if !a.validate() {
			return false
		}
	}

	return true
}

// Get the maximum seen count of red, green, and blue over all rounds
func (fg *full_game) maxColorCount() [3]int {
	r, g, b := 0, 0, 0
	for _, a := range fg.individual_games {
		if a.red > r {
			r = a.red
		}
		if a.green > g {
			g = a.green
		}
		if a.blue > b {
			b = a.blue
		}
	}
	return [3]int{r, g, b}
}

func parseGames(line string) full_game {
	// Could/should be done with fmt.Sscanf!

	// Separate game id from game data
	split_one := strings.Split(line, ": ") // [ "game x" games.. ]

	// Get game id from first part of split string
	game_id_string := strings.Split(split_one[0], " ") // [ "game" "x" ]
	game_id, _ := strconv.Atoi(game_id_string[1])

	// Separate the individual games playes
	individual_games := strings.Split(split_one[1], "; ") // [ "game1"; "game2"]
	var games_slice []*game

	// For every single game, get the counts of red, green, and blue cubes
	for _, round := range individual_games {
		r, g, b := 0, 0, 0
		colors_in_game := strings.Split(round, ", ") // [ "1 red" "2 blue" "3 green" ]

		for _, color_in_game := range colors_in_game {
			split_cig := strings.Split(color_in_game, " ") // [ "1" "red" ]
			color_count, _ := strconv.Atoi(split_cig[0])
			if split_cig[1] == "red" {
				r = color_count
			} else if split_cig[1] == "blue" {
				b = color_count
			} else {
				g = color_count
			}
		}
		// As one game can have several rounds, a slice is created to hold each
		// indivudual round
		games_slice = append(games_slice, &game{red: r, blue: b, green: g})
	}

	// Return and object with all the relevant information
	return full_game{game_id: game_id, individual_games: games_slice, valid_game: false}
}

func one(data []string) {
	var result int

	// Iterate over each game and check if it is valid
	for _, line := range data {
		fg := parseGames(line)
		if fg.validate() {
			result += fg.game_id
		}
	}

	fmt.Printf("Result 1: %v \n", result)
}

func two(data []string) {
	var result int

	// Iterate over all the games and get the max color counts from each game
	for _, line := range data {
		fg := parseGames(line)
		min_req := fg.maxColorCount()
		result += min_req[0] * min_req[1] * min_req[2]
	}

	fmt.Printf("Result 2: %v \n", result)
}

func main() {
	// data := readinput.ReadText("example1.txt")
	data := readinput.ReadText("input.txt")

	one(data)
	two(data)
}
