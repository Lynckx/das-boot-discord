package gametypes

import (
	"lynckx/das-boot-discord/pkg/games/gametypes/mexxen"
	"lynckx/das-boot-discord/pkg/games/player"
	"lynckx/das-boot-discord/pkg/games/round"
)

var gameMap map[string]gametypes.GameType

type GameType struct {
	Name    string
	Players []player.Player
	round   round.Round
}

func init() {
	gameMap = make(map[string]GameType)
	gameMap[mexxen.GameName] = &mexxen.newGame()
}
