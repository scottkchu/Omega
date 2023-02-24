package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Configs struct {
	DiscordToken string `json:"discordToken"`
	CohereToken string	`json:"cohereToken"`
}

func GetConfigs() (configs Configs, err error) {
	configsJSON, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Error opening config.json: ", err)
		return configs, err
	}
	defer configsJSON.Close()

	bytes, err := ioutil.ReadAll(configsJSON)
	if err != nil {
		fmt.Println("Error reading config.json: ", err)
		return configs, err
	}

	err = json.Unmarshal(bytes, &configs)
	return configs, err
}

func GetCommand(m *discordgo.MessageCreate) (cmd string, err error){
	str := strings.Split(m.Content, " ")
	
	if len(str) == 0 {
		return "", nil
	}

	return str[0], nil
}

func GetPrompt(m *discordgo.MessageCreate) (prompt string, err error){
	str := strings.Split(m.Content, " ")
	
	if len(str) < 2 {
		return "", nil
	}

	return strings.Join(str[1:], " "), nil
}