package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"lynckx/das-boot-discord/pkg/dasbot"

	"github.com/bwmarrin/discordgo"
)

var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Owner Account Token")

	flag.Parse()

	if Token == "" {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	fmt.Println("Starting discord go")
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dasBot := dasbot.CreateBot()
	ctx, cnl := context.WithCancel(context.Background())
	defer cnl()
	go dasBot.Start(ctx)

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(dasBot.DiscordHandler)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
