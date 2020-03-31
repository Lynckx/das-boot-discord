package dasbot

import (
	"context"
	"fmt"
	"sync"

	"github.com/bwmarrin/discordgo"

	"lynckx/das-boot-discord/pkg/dasbot/commands"
	"lynckx/das-boot-discord/pkg/dasbot/message"
	"lynckx/das-boot-discord/pkg/games"
)

type Bot struct {
	discordMessageChannel chan *message.Message
	cmd                   *commands.CommandHandler
	wg                    *sync.WaitGroup
}

func CreateBot() Bot {
	var newBot Bot
	newBot.discordMessageChannel = make(chan *message.Message, 1)
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
			b.cmd.SendMessageToListeners(msg)
			go func() {
				if msg.HasBotPrefix() {
					if err := b.cmd.Emit(msg); err != nil {
						fmt.Printf("Got an error %v\n", err)
					}
					fmt.Printf("received message: %v From User: %+v\n", msg.GetMessageContent(), msg.GetAuthor())
				}
			}()
		}
	}
}

func (b Bot) DiscordHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}
	b.discordMessageChannel <- message.NewMessage(s, m)
}

//discordCommandHandler takes a string that has the "/db " prefix and calls the right command
func (b Bot) initializeCommands() error {
	b.cmd.LoadStandartCommand()
	//initialize games
	b.cmd.LoadCommand(games.LoadCommand())
	//initialize
	return nil
}
