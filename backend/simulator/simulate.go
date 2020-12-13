package simulator

import (
	"sort"
)

// Simulate N random players against each other
func Simulate(amount int) []*Player {
	players := make([]*Player, amount)

	for i := 0; i < amount; i++ {
		players[i] = CreateRandomPlayer()
	}

	n := 10000

	for i := 0; i < n; i++ {
		MultiMatch(players)
	}

	sort.Slice(players[:], func(i, j int) bool {
		return players[i].Elo > players[j].Elo
	})

	return players
}
