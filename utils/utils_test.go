package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMessage(t *testing.T) {
	assert := assert.New(t)

	msg1 := "!help"
	command, args := ParseMessage(msg1)
	assert.Equal("help", command)
	assert.Len(args, 0)

	msg2 := "!help 1 2"
	command, args = ParseMessage(msg2)
	assert.Equal("help", command)
	assert.Len(args, 2)
}
func TestRemovePrefix(t *testing.T) {
	assert := assert.New(t)
	invokeText := "!help"
	command := RemovePrefix(invokeText)
	assert.Equal("help", command)
}
