package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/houssemcharf/Mika/config"
)

func main() {
	cfg, err := config.Grab()
	if err != nil {
		panic(err)
	}
	// instancing a new session for discord
	session, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		return
	}
	// add subscription of commands here
	session.AddHandler(handleCommand)
	// opening a websocket connection to discord and start listening.
	err = session.Open()
	if err != nil {
		fmt.Println("Error opening a websock connection,", err)
		return
	}
	fmt.Println("Bot is running. Press CTRL+c to exit.")
	// handling CTRL+C and other term signals
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	// close session
	session.Close()
}
