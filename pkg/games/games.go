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
	ID        string
	players   map[users.User]player.Player
	gameTypes []gametypes.GameType
}

// Load returns all game commands
func LoadCommand() *commands.CommandLoader {
	return &commands.CommandLoader{
		Id:      constants.PLAY_COMMAND,
		Alias:   []string{"p", "speel"},
		CmdFunc: PlayCommand,
	}
}

func PlayCommand(cmd *commands.CommandHandler, msg *message.Message) error {
	// "/db play" or "/db play <gametype>"
	commands := msg.GetCommandListAfterPrefix()
	if len(commands) > 1 {

	} else {
		err := msg.Respond([]string{
			"What would you like to play?",
			GetGamesList(),
		})
		if err != nil {
			return err
		}
		ch := cmd.AddMessageListener()
		defer cmd.RemoveMessageListener(ch)
		for {
			response := <-ch
			if response.GetAuthor() == msg.GetAuthor() {
				selectedGame := response.GetCommandListAfterPrefix()[0]
				break
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
