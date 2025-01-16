package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

var CommandHandlers = map[string]func(session *discordgo.Session, msg *discordgo.MessageCreate, args *[]string){
	"!add":    Add,
	"!remove": Remove,
}

func HandleCommand(session *discordgo.Session, msg *discordgo.MessageCreate) {

	content := strings.Fields(msg.Content)
	if len(content) == 0 {
		return
	}

	command := content[0]
	args := content[1:]

	if handler, exists := CommandHandlers[command]; exists {
		handler(session, msg, &args)
	}
}
