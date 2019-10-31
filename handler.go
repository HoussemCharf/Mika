package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/houssemcharf/Mika/command"
	"github.com/houssemcharf/Mika/utils"
)

// HandleCommand process the message and calls the appropriate function to be exectuted.
func handleCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	defer func() {
		err := recover()
		if err != nil {

		}
	}()

	cmd, args := utils.ParseMessage(m.Content)

	handler, err := command.HandlerFor(cmd)
	if err != nil {

	}
	err = handler.Handle(m, cmd, args...)

}
