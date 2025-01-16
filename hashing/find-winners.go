package main

import (
	"sort"
)

type player struct {
	won   int
	loose int
}

// this solution is not fully optimized, this just solved the problem statement
func findWinners(matches [][]int) [][]int {
	players := map[int]*player{}

	// O(n)
	for _, match := range matches {
		winner := players[match[0]]

		if winner == nil {
			winner = &player{}
			players[match[0]] = winner
		}

		winner.won++

		looser := players[match[1]]

		if looser == nil {
			looser = &player{}
			players[match[1]] = looser
		}

		looser.loose++
	}

	loosers := []int{}
	winners := []int{}

	// O(m)
	for index, player := range players {
		if player.loose == 1 {
			loosers = append(loosers, index)
		}

		if player.loose == 0 && player.won > 0 {
			winners = append(winners, index)
		}
	}

	sort.IntSlice(winners).Sort()
	sort.IntSlice(loosers).Sort()

	result := [][]int{
		winners,
		loosers,
	}

	return result
}
