package player

import (
	"math/rand"
	"time"

	"lynckx/das-boot-discord/pkg/games/score"
	"lynckx/das-boot-discord/pkg/users"
)

type Player struct {
	User  users.User
	Score score.Score
}

func NewPlayer(u users.User) Player {
	var newPlayer Player
	newPlayer.User = u
	return newPlayer
}

func SelectRandomPlayer(players []Player) Player {
	var amountOfPlayers = len(players)
	rand.Seed(time.Now().UnixNano())
	return players[rand.Intn(amountOfPlayers)]
}
