package dasbot

import (
	"context"
	"fmt"
	"sync"

	"github.com/bwmarrin/discordgo"

	"lynckx/das-boot-discord/pkg/commands"
	"lynckx/das-boot-discord/pkg/dasbot/message"
)

type Bot struct {
	discordMessageChannel chan message.Message
	cmd                   commands.CommandHandler
	wg                    *sync.WaitGroup
}

func CreateBot() Bot {
	var newBot Bot
	newBot.discordMessageChannel = make(chan message.Message, 1)
	newBot.wg = &sync.WaitGroup{}
	newBot.cmd = commands.CreateCommandHandler()
	return newBot
}

func (b Bot) Start(ctx context.Context) {
	b.initializeCommands()
	b.wg.Add(1)
	go b.BotHandler(ctx)
	b.wg.Wait()
}

func (b Bot) BotHandler(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			b.wg.Done()
			return
		case msg := <-b.discordMessageChannel:
			fmt.Printf("Got message from discordMessageChannel, %v\n", msg)
			go func() {
				if msg.HasBotPrefix() {
					if err := b.cmd.Emit(msg); err != nil {
						fmt.Printf("Got an error %v\n", err)
					}
					fmt.Printf("recieved message: %v\n", msg.GetMessageContent())
				}
			}()
		}
	}
}

func (b Bot) DiscordHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	fmt.Printf("Got message from discord, %v\n", m.Content)
	b.discordMessageChannel <- message.NewMessage(s, m)
	fmt.Printf("Got message from discord, %v\n", m.Content)
}

//discordCommandHandler takes a string that has the "/db " prefix and calls the right command
func (b Bot) initializeCommands() error {
	b.cmd.LoadStandartCommand()
	//initialize games
	return nil
}
