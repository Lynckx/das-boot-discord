package round

import (
	"lynckx/das-boot-discord/pkg/games/player"
	"lynckx/das-boot-discord/pkg/games/round/turn"
	"lynckx/das-boot-discord/pkg/users"
)

type Round struct {
	players map[users.User]player.Player
	turn    turn.Turn
}

func NewRound(players []player.Player) {
	var newRound Round
	newRound.players = players
	newRound.turn = turn.NewTurn(player.SelectRandomPlayer(players))
	return newRound
}

func (r Round) nextTurn() {
	currentPlayer := r.turn.Player
	for i, p := range r.players {
		if p == currentPlayer {
			newPlayer := r.players[i+1]
			break
		}
	}

}
