package uscf

import gr "github.com/Kashomon/gorating"

//
// The USCF Rating system
//

// The provisional winning expectency.
//
// Basically, if the delta between the player and opponent's score is greater
// than 400, we consider a win to be guaranteed.
func provWinExpect(player, opp *EloRating) float64 {
	delta := player.Score - opp.Score
	if delta <= -400 {
		return 0
	} else if delta >= 400 {
		return 1
	} else {
		return 0.5 + delta/800.0
	}
}

type UscfSystem struct {
	// Map from player id to games played
	pmap map[string][]gr.Game
}

func (t *UscfSystem) RateAll(games []gr.Game) ([]gr.PlayerRating, error) {
	t.pmap = gr.PlayerMap(games)
	return nil, nil
}

func (t *UscfSystem) Rate(player gr.PlayerRating, games []gr.Game) (gr.PlayerRating, error) {
	t.pmap = gr.PlayerMap(games)

	_, ok := t.pmap[player.UnqiueId()]
	if !ok {
		return nil, nil
	}
	return nil, nil
}

var _ gr.RatingSystem = &UscfSystem{}
