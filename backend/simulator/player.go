package simulator

import "math/rand"

// Player describes a simulated player
type Player struct {
	ID     uint
	Name   string
	Elo    float64
	Skill  float64
	Played uint
	Wins   uint
}

var counter uint = 0

// CreatePlayer create a new player
func CreatePlayer(name string, skill float64) *Player {
	counter++
	return &Player{ID: counter, Name: name, Elo: 1500, Skill: skill, Played: 0, Wins: 0}
}

// CreateRandomPlayer create player with random skill value
func CreateRandomPlayer() *Player {
	skill := rand.NormFloat64()*200 + 1500
	name := randomString(10)
	return CreatePlayer(name, skill)
}

// GetWinrate Get player winrate
func (p Player) GetWinrate() float64 {
	return float64(p.Wins) / float64(p.Played)
}
