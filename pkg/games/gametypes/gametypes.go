package gametypes

import (
	"lynckx/das-boot-discord/pkg/games/gametypes/mexxen"
	"lynckx/das-boot-discord/pkg/games/player"
	"lynckx/das-boot-discord/pkg/games/round"
)

type GameType struct {
	Name    string
	Players []player.Player
	round   round.Round
}

type GameLibrary map[string]GameType

func GetGameLibrary() GameLibrary {
	gameMap = make(map[string]GameType)
	gameMap[mexxen.GameName] = &mexxen.Impl()
}
