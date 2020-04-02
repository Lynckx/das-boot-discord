package dasbot

import (
	"context"
	"fmt"
	"sync"

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

func (b Bot) initializeCommands() error {
	b.cmd.LoadStandartCommand()
	//initialize games
	b.cmd.LoadCommand(games.LoadCommand())
	//initialize
	return nil
}

func (b Bot) BotHandler(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			b.wg.Done()
			return
		case msg := <-b.discordMessageChannel:
			fmt.Printf("Got message from discordMessageChannel, %v\n", msg)
			b.cmd.Broadcast(msg)
		}
	}
}
