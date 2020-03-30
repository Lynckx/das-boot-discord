package message

import (
	"fmt"
	"lynckx/das-boot-discord/domain/constants"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Message struct {
	s *discordgo.Session
	m *discordgo.MessageCreate
}

func (msg Message) HasBotPrefix() bool {
	if strings.HasPrefix(msg.m.Content, constants.BOT_PREFIX) {
		if strings.Count(msg.m.Content, constants.BOT_PREFIX) == 1 {
			return true
		}
		fmt.Errorf("the prefix can only be once inside the command")
	}
	// the msg does not contain the prefix
	return false
}

func (msg Message) GetFirstCommand() string {
	firstString := strings.SplitN(strings.Split(msg.m.Content, constants.BOT_PREFIX)[1], " ", 2)[0]
	return firstString
}

func (msg Message) GetAfterPrefix() string {
	stringAfterPrefix := strings.SplitAfterN(msg.m.Content, constants.BOT_PREFIX, 2)[1]
	return stringAfterPrefix
}

/*
func (msg Message) HasBotCommands() ([]commands.Command, error) {
	var botCommands []commands.Command
	cmdStrings := strings.Split(msg.m.Content, constants.BOT_PREFIX)[1:]
	for _, cmdString := range cmdStrings{
		cmd, err := commands.IsBotCommand(cmdString)
		if err != nil{
			fmt.Printf("Got error: %v",err)
		}
		botCommands = append(botCommands, cmd)
	}
	return botCommands
}
*/
func (msg Message) Respond(responses []string) error {
	for _, resp := range responses {
		if resp != "" {
			if m, err := msg.s.ChannelMessageSend(msg.m.ChannelID, resp); err != nil {
				fmt.Println("Got an error, message recieved: %v", m)
				return err
			}
		}
	}
	return nil
}

func (msg Message) GetMessageContent() string {
	return msg.m.Content
}

func NewMessage(s *discordgo.Session, m *discordgo.MessageCreate) Message {
	return Message{
		s: s,
		m: m,
	}
}
