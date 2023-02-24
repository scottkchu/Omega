package cmd

import (
	"fmt"
	"omega/common"

	"github.com/bwmarrin/discordgo"
)

func ChatCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	cmd, err := common.GetCommand(m)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}


	if cmd == "/chat" {
		fmt.Println("Chat command received")

		prompt, err := common.GetPrompt(m)
		if err != nil {
			fmt.Println("Error creating Discord session: ", err)
			return
		}

		response, err := common.Generate(prompt)
		if err != nil {
			fmt.Println("Error creating Discord session: ", err)
			return
		}

		s.ChannelMessageSend(m.ChannelID, response)
	}
}