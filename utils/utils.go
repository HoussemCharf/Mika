package utils

import (
	"log"
	"strings"

	"github.com/houssemcharf/Mika/config"
)

// ParseMessage parses ...
func ParseMessage(message string) (m string, sl []string) {
	m = RemovePrefix(message)
	splitted := strings.Split(m, " ")
	if m == "" {
		return
	}
	if len(splitted) == 1 {
		return
	}
	return splitted[0], splitted[1:]
}

// RemovePrefix returns raw command
func RemovePrefix(invoked string) (s string) {
	cfg, err := config.Grab()
	if err != nil {
		log.Println("[RemovePrefix]: ", err)
		return
	}
	return strings.Replace(invoked, cfg.Token, "", 1)
}
