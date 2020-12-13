package simulator

import (
	"math/rand"
	"sort"
)

// Match two players
func Match(p1 *Player, p2 *Player) float64 {
	prediction := Predict(p1.Elo, p2.Elo)
	result := rand.Float64()

	var score float64
	if result < Predict(p1.Skill, p2.Skill) {
		score = 1
		p1.Wins++
	} else {
		score = 0
		p2.Wins++
	}

	p1.Elo += 24 * (score - prediction)
	p2.Elo += 24 * ((1 - score) - (1 - prediction))
	p1.Played++
	p2.Played++
	return score
}

// MultiMatch Match N players against each other
func MultiMatch(players []*Player) {
	predictions := make(map[uint]float64)

	sort.Slice(players, func(i, j int) bool {
		return players[i].Elo < players[j].Elo
	})

	for i := 0; i < len(players); i++ {
		player := players[i]
		var total float64 = 0
		for j := 0; j < len(players); j++ {
			if j == i {
				continue
			}

			opponent := players[j]
			total += Predict(player.Elo, opponent.Elo)
		}

		prediction := total / float64(len(players)-1)
		predictions[player.ID] = prediction

	}

	for i := 0; i < len(players); i++ {
		for j := 0; j < len(players)-1; j++ {
			p1 := players[j]
			p2 := players[j+1]

			result := rand.Float64()
			if result < Predict(p1.Skill, p2.Skill) {
				players[j], players[j+1] = players[j+1], players[j]
			}
		}
	}

	for i := 0; i < len(players); i++ {
		player := players[i]
		prediction := predictions[player.ID]

		score := float64(i) / float64(len(players)-1)

		eloChange := 24 * (score - prediction)
		player.Elo += eloChange
	}
}
