package dgoc

import "github.com/bwmarrin/discordgo"

type Context struct {
	Name string

	Session *discordgo.Session
	Message *discordgo.Message
}
