package round

import (
	"lynckx/das-boot-discord/pkg/games/player"
	"lynckx/das-boot-discord/pkg/games/round/turn"
)

type Round struct {
	players []player.Player
	turn    turn.Turn
}

func NewRound(players []player.Player) Round {
	var newRound Round
	newRound.players = players
	newRound.turn = turn.NewTurn(player.SelectRandomPlayer(players))
	return newRound
}

func (r *Round) nextTurn() {
	currentPlayer := r.turn.Player
	var nextPlayer player.Player
	for i, p := range r.players {
		if p == currentPlayer {
			nextPlayer = r.players[i+1]
			break
		}
	}
	r.turn = turn.NewTurn(nextPlayer)
}
