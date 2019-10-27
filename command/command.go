package command

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var (
	handlers = make(map[string]Command)
)

// Command interface
type Command interface {
	Handle(message *discordgo.MessageCreate, cmd string, args ...string) error
}

// Register to subscribe to the bot
func Register(name string, c Command) {
	handlers[name] = c
}

type CommandFunc func(message *discordgo.MessageCreate, cmd string, args ...string) error

// Handle to attach
func (f CommandFunc) Handle(message *discordgo.MessageCreate, cmd string, args ...string) error {
	return f(message, cmd, args...)
}

// HandlerFor tp double check
func HandlerFor(name string) (Command, error) {
	if handler, ok := handlers[name]; ok {
		return handler, nil
	}
	return nil, fmt.Errorf("No handler for command: %s", name)
}
