package dasbot

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/bwmarrin/discordgo"

	"lynckx/das-boot-discord/pkg/commands"
)

type Message struct {
	content []string
	dcb     DiscordCallback
}

type DiscordCallback struct {
	s *discordgo.Session
	m *discordgo.MessageCreate
}

type Bot struct {
	responseMessageChannel chan Message
	receivedMessageChannel chan Message
	wg                     *sync.WaitGroup
}

func CreateBot() Bot {
	var newBot Bot
	newBot.responseMessageChannel = make(chan Message, 1)
	newBot.receivedMessageChannel = make(chan Message, 1)
	newBot.wg = &sync.WaitGroup{}
	return newBot
}

func (b Bot) Run(ctx context.Context) {
	b.wg.Add(1)
	go b.ReceivedHandler(ctx)
	b.wg.Add(1)
	go b.ResponseHandler(ctx)
	b.wg.Wait()
}

func (b Bot) ReceivedHandler(ctx context.Context) {
	select {
	case <-ctx.Done():
		b.wg.Done()
		return
	case msg := <-b.receivedMessageChannel:
		fmt.Printf("recieved message: %v", msg.content)
	}
}

func (b Bot) ResponseHandler(ctx context.Context) {
	select {
	case <-ctx.Done():
		b.wg.Done()
		return
	case msg := <-b.responseMessageChannel:
		for _, response := range msg.content {
			msg.dcb.s.ChannelMessageSend(msg.dcb.m.ChannelID, response)
		}
	}
}

//discordCommandHandler takes a string that has the "/db " prefix and calls the right command
func (b Bot) CommandHandler(cmdString string) ([]string, error) {
	var responses []string
	cmds := strings.SplitAfter(cmdString, " ")
	for _, cmd := range cmds {
		switch cmd {
		case "play":
			/*
				if games.IsGame(cmd[i+1]){
					responses = append(responses,"Playing: %v", cmd[i+1])
					b.Game := games.LoadGame(cmd[i+1])
					responses = append(responses,)
				} else {

				}
			*/
			//Play a game
		case "Das Boot":
			return commands.DasBoot(), nil
		default:
			return responses, errors.New("Command does not exist")
		}
	}
	return responses, nil
}

func (b Bot) DiscordHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if strings.HasPrefix(m.Content, "/db ") {
		prefixCount := strings.Count(m.Content, "/db ")
		if prefixCount > 1 {
			fmt.Errorf("the prefix can only be once inside the command")
			return
		}
		b.receivedMessageChannel <- Message{
			content: strings.Split(m.Content, "/db ")[1:],
			dcb: DiscordCallback{
				s: s,
				m: m,
			},
		}
		/*
			fmt.Println("Das Bot was called")
			fmt.Printf("Message ID: 		%+v\n", m.ID)
			fmt.Printf("Message Content: 	%+v\n", m.Content)
			fmt.Printf("Message Guild ID: 	%+v\n", m.GuildID)
			fmt.Printf("Message Timestamp: 	%+v\n", m.Timestamp)
			fmt.Printf("Message Author: 	%+v\n", m.Author)
			fmt.Printf("Message Type: 		%+v\n", m.Type)
			fmt.Printf("Message Member: 	%+v\n", m.Member)

			responses, err := b.discordCommandHandler(strings.Split(m.Content, "/db "))
			if err != nil {
				fmt.Errorf("Tried to handle the command, got error: %v", err)
				return
			}
			for _, r := range responses{
				msg, err := s.ChannelMessageSend(m.ChannelID, r)
				if err != nil {
					fmt.Errorf("tried to send a message, got error: %v", err)
					fmt.Errorf("tried to send a message, got message: %v", msg)
				}
			}
		*/
	}
}
