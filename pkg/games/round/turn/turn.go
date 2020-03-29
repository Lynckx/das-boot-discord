package turn

import "lynckx/das-boot-discord/pkg/games/player"

type Turn struct {
	Player player.Player
}

func NewTurn(player player.Player) Turn {
	var newTurn Turn
	newTurn.Player = player
	return newTurn
}
