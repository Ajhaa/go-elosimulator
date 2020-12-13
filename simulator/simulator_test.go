package simulator

import (
	"math"
	"testing"
)

func TestPlayer(t *testing.T) {
	player := CreateRandomPlayer()
	if player.Elo != 1500 {
		t.Errorf("Expected elo to be 1500, was %f", player.Elo)
	}
}

func TestMatchEloChange(t *testing.T) {
	p1 := CreateRandomPlayer()
	p2 := CreateRandomPlayer()

	Match(p1, p2)

	sum := p1.Elo + p2.Elo

	if math.Abs(sum-3000) > 0.01 {
		t.Errorf("Expected elo sum to be 3000, was %f", sum)
	}
}
