package main

// This is what happens when you don't think ahead, kids

import (
	readinput "aoc_2023/inputreader"
	"fmt"
	"strconv"
	"strings"

	ms "github.com/Julmust/merge-sort"
)

type handBid struct {
	hand        []int
	bid, fullID int
}

func charToInt(hand []string) []int {
	var intHand []int
	for idx := range hand {
		val := hand[idx]

		switch val {
		case "T":
			intHand = append(intHand, 10)
		case "J":
			intHand = append(intHand, 11)
		case "Q":
			intHand = append(intHand, 12)
		case "K":
			intHand = append(intHand, 13)
		case "A":
			intHand = append(intHand, 14)
		default:
			v, _ := strconv.Atoi(val)
			intHand = append(intHand, v)
		}
	}
	return intHand
}

func parseInput(data string) handBid {
	f := strings.Split(data, " ")
	bid, _ := strconv.Atoi(f[1])
	hand := charToInt(strings.Split(f[0], ""))

	var fullid string
	for _, i := range hand {
		j := strconv.Itoa(i)
		if len(j) == 1 { // If single digit, prepend a 0 (eg. 7 -> 07)
			fullid += "0" + j
		} else {
			fullid += j
		}
	}
	fullidint, _ := strconv.Atoi(fullid)

	newHand := handBid{hand, bid, fullidint}

	return newHand
}

func bucketHand(hand []int) int {
	var maxCnt, scndCnt int
	set := make(map[int]bool)
	for _, k := range hand {
		set[k] = true
	}

	for k := range set {
		var cnt int
		for _, i := range hand {
			if i == k {
				cnt++
			}
		}

		if cnt > maxCnt {
			scndCnt = maxCnt
			maxCnt = cnt
		} else if cnt > scndCnt {
			scndCnt = cnt
		}
	}

	// Five           Four          Three of a kind
	if maxCnt == 5 || maxCnt == 4 {
		return maxCnt + 1
	} else if maxCnt == 3 {
		if scndCnt == 2 {
			return 4
		}
		return maxCnt
	} else if maxCnt == 2 {
		if scndCnt == 2 {
			return 2
		}
		return 1
	}

	return 0
}

func sortHands(hands []handBid) []int {
	// Build map from fullintid & bid
	m := make(map[int]int)
	var unsorted []int

	for _, h := range hands {
		m[h.fullID] = h.bid
		unsorted = append(unsorted, h.fullID)
	}

	var output []int

	if len(hands) > 0 {
		sortedList := ms.Sort(unsorted)

		for _, i := range sortedList {
			output = append(output, m[i])
		}
	}

	return output
}

func one(data []string) {
	var hc, op, tp, tok, fook, fiok, fh []handBid
	buckets := [7][]handBid{hc, op, tp, tok, fook, fiok, fh}

	for _, r := range data {
		newHand := parseInput(r)
		score_idx := bucketHand(newHand.hand)

		buckets[score_idx] = append(buckets[score_idx], newHand)
	}

	var rankedList []int
	for _, b := range buckets {
		rankedList = append(rankedList, sortHands(b)...)
	}

	var output int
	for idx, i := range rankedList {
		output += (idx + 1) * i
	}

	fmt.Printf("Solution 1: %d\n", output)

}

func main() {
	// data := readinput.ReadText("example.txt")
	data := readinput.ReadText("input.txt")

	one(data)
}
