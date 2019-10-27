package utils

import (
	"strings"

	"github.com/houssemcharf/Mika/config"
)

// ParseMessage parses ...
func ParseMessage(message string) (string, []string) {
	message = RemovePrefix(message)
	splitted := strings.Split(message, " ")
	if message == "" {
		return "", nil
	}

	if len(splitted) == 1 {
		return message, []string{}
	}
	return splitted[0], splitted[1:]
}

// RemovePrefix returns raw command
func RemovePrefix(invoked string) string {
	cfg, err := config.Grab()
	if err != nil {
		panic(err)
	}
	return strings.Replace(invoked, cfg.Token, "", 1)
}
