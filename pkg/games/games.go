package games

import (
	"lynckx/das-boot-discord/domain/constants"
	"lynckx/das-boot-discord/pkg/dasbot/commands"
	"lynckx/das-boot-discord/pkg/dasbot/message"
	"lynckx/das-boot-discord/pkg/games/gametypes"
	"lynckx/das-boot-discord/pkg/games/player"
	"lynckx/das-boot-discord/pkg/users"
)

type Games struct {
	ID          string
	players     map[users.User]player.Player
	GameLibrary gametypes.GameLibrary
}

// Load returns all game commands
func LoadCommand() *commands.CommandLoader {
	return &commands.CommandLoader{
		Id:      constants.PLAY_COMMAND,
		Alias:   []string{"p", "speel"},
		CmdFunc: PlayCommand,
	}
}

func PlayCommand(cmd *commands.CommandHandler, msg *message.Message, args commands.Arguments) error {
	// "/db play" or "/db play <gametype>"
	commands := msg.GetCommandListAfterPrefix()
	if len(commands) > 1 {
		if commands[0] == constants.BOT_PREFIX {
			if IsGame(commands[1]) {
				//
				return nil
			} else {
				err := msg.Respond([]string{
					"Unknown Game, Please enter the following",
					GetGamesList(),
				})
				if err != nil {
					return err
				}
			}
		}
	} else {
		err := msg.Respond([]string{
			"What would you like to play?",
			GetGamesList(),
		})
		if err != nil {
			return err
		}
	}
	return GetGameTypeFromUser(cmd, msg, args)
}
func GetGameTypeFromUser(cmd *commands.CommandHandler, msg *message.Message, _ commands.Arguments) error {
	// if we did not get a game, we either did not recognize the GameTypeName or it was not given
	listener := cmd.AddMessageListener()
	defer cmd.RemoveMessageListener(listener)
	for {
		newMsg := <-listener.Ch
		if newMsg.GetAuthor() == msg.GetAuthor() {
			selectedGame := newMsg.GetCommandListAfterPrefix()[0]
			if IsGame(selectedGame) {
				//
				break
			} else {
				err := newMsg.Respond([]string{
					"Unknown Game, Please enter the following",
					GetGamesList(),
				})
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func IsGame(gameType string) bool {
	switch gameType {
	case "mexxen":
		return true
	default:
		return false
	}
}

func GetGamesList() string {
	return "```" + `
	Games:	
		mexxen
	` + "```"
}
