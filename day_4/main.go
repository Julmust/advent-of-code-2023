package main

import (
	readinput "aoc_2023/inputreader"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type card struct {
	card_id                  int
	winning_card, given_card []int
}

func parseIndividualCard(card string) []int {
	// Split string into values, note that some double spaces exist
	raw_split := strings.Split(card, " ")
	var refined_split []int
	for _, i := range raw_split {
		if i != "" {
			val, _ := strconv.Atoi(i)
			refined_split = append(refined_split, val)
		}
	}

	return refined_split
}

func parseCards(row string) (int, []int, []int) {
	var card_id int
	fmt.Sscanf(row, "Card %d", &card_id)

	cards_info := strings.Split(row, ": ")[1]
	w_g := strings.Split(cards_info, " | ")
	winning, given := w_g[0], w_g[1]
	winning_card := parseIndividualCard(winning)
	given_card := parseIndividualCard(given)

	return card_id, winning_card, given_card
}

func checkMatches(winning, given []int) int {
	var matches_found int
	for _, gnum := range given {
		for _, wnum := range winning {
			if gnum == wnum {
				matches_found += 1
			}
		}
	}

	return matches_found
}

func checkResMatches(matches_found int) int {
	if matches_found != 0 {
		return int(math.Pow(2, float64(matches_found)-1))
	}
	return 0
}

func one(data []string) {
	var res int
	for _, row := range data {
		_, winning, given := parseCards(row)
		matches := checkMatches(winning, given)
		res += checkResMatches(matches)
	}

	fmt.Printf("Solution 1: %d\n", res)
}

func two(data []string) {
	var total int
	// Create map with all parsed cards as maps and a slice containing all cards to be checked,
	// effectively acting as a queue
	cardMap := make(map[int]card)
	var cardSlice []card
	for _, row := range data {
		card_id, winning, given := parseCards(row)
		newCard := card{card_id: card_id, winning_card: winning, given_card: given}
		cardMap[card_id] = newCard
		cardSlice = append(cardSlice, newCard)
	}

	// Not the prettiest implementation but I'm tired
	loopVar := 0
	for {
		if len(cardSlice) == loopVar {
			break
		}
		total += 1
		c := cardSlice[loopVar]
		matches := checkMatches(c.winning_card, c.given_card)
		// If matches are found, append their corresponding card to the end of the queue
		if matches > 0 {
			for i := 1; i <= matches; i++ {
				cardSlice = append(cardSlice, cardMap[c.card_id+i])
			}
		}
		loopVar += 1
	}

	fmt.Printf("Solution 2: %d\n", total)
}

func main() {
	// data := readinput.ReadText("example.txt")
	data := readinput.ReadText("input.txt")

	one(data)
	two(data)
}
