package simulator

import (
	"testing"
)

func TestPlayer(t *testing.T) {
	player := CreateRandomPlayer()
	if player.Elo != 1500 {
		t.Errorf("Expected elo to be 1500, was %f", player.Elo)
	}
}
