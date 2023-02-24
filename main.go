package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"omega/cmd"

	"github.com/bwmarrin/discordgo"
)

// var (
// 	discordToken string
// )

// func init() {
// 	configs, err := common.GetConfigs()
// 	if err != nil {
// 		fmt.Println("Error getting configs: ", err)
// 		return
// 	}

// 	discordToken = configs.DiscordToken
// }

func main() {
	sess, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	sess.AddHandler(cmd.ChatCommand)

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = sess.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
		return
	}
	defer sess.Close()

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
