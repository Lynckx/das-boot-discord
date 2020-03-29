package mexxen

import (
	"errors"
	"lynckx/das-boot-discord/pkg/games/dice"
	"lynckx/das-boot-discord/pkg/games/player"
	"lynckx/das-boot-discord/pkg/games/round"
	"lynckx/das-boot-discord/pkg/games/round/turn"
	"lynckx/das-boot-discord/pkg/users"
)

var GameName = "mexxen"

type Mexxen interface {
	AddNewPlayer(user users.User)
	RemovePlayer(user users.User) error
	isUserKnown(user users.User) bool
	IsUserPlaying(r round.Round, user users.User) bool
	loadNewRound
}

type MexxenImpl struct {
	Name    string
	Players []player.Player
	round   round.Round
}

func (m *MexxenImpl) newGame() {
	// load new round
	m.loadNewRound()

}

func (m *MexxenImpl) AddNewPlayer(user users.User) {
	if !m.isUserKnown(user) {
		m.Players = append(m.Players, player.NewPlayer(user))
	}
}

func (m *MexxenImpl) RemovePlayer(user users.User) error {
	for i, p := range m.Players {
		if p.User == user {
			m.Players = append(m.Players[:i], m.Players[i+1:]...)
			return nil
		}
	}
	return errors.New("Player not found")
}

func (m MexxenImpl) isUserKnown(user users.User) bool {
	for _, p := range m.Players {
		if p.User == user {
			return true
		}
	}
	return false
}

func (m MexxenImpl) IsUserPlaying(r round.Round, user users.User) bool {
	for _, p := range m.round.players {
		if p.User == user {
			return true
		}
	}
	return false
}

func (m *MexxenImpl) loadNewRound() []player.Player {
	// collect all players
	m.round.players = m.Players
	m.round.turn = turn.Turn{}
}

func (m MexxenImpl) newRound() {
	// load new round
	m.loadNewRound()

}

func (m MexxenImpl) newTurn() Results {

}

type mexxenDices struct {
	diceOne dice.Dice
	DiceTwo dice.Dice
}

type Status string

type GameRule struct {
	Name
}

const (
	KansOp = "Kans Op"
	Mex    = "Mex"
)

func GameType() string {
	return "mexxen"
}
