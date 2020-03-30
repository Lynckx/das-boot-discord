package games

import (
	"lynckx/das-boot-discord/pkg/games/player"
	"lynckx/das-boot-discord/pkg/gametypes"
	"lynckx/das-boot-discord/pkg/users"
)

type Game struct {
	players  map[users.User]player.Player
	gameType gametypes.GameType
}

/*
func LoadGame(gameType string) (string, Game) {
	init()
	var newGames Game

	return newGames
}
*/
func IsGame(gameType string) bool {
	switch gameType {
	case "mexxen":
		return true
	default:
		return false
	}

}
