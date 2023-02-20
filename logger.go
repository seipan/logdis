package logdis

import (
	"github.com/bwmarrin/discordgo"
)

type Fields map[string]interface{}

func (b *Bot) Info(f Fields, message string, user string) {
	str := "Info : "
	str += parse(f, message, user)

	for _, guild := range b.bot.State.Guilds {
		// Get channels for this guild
		channels, _ := b.bot.GuildChannels(guild.ID)

		for _, c := range channels {
			// Check if channel is a guild text channel and not a voice or DM channel
			if c.Type != discordgo.ChannelTypeGuildText {
				continue
			}

			// Send text message
			b.bot.ChannelMessageSend(
				c.ID,
				str,
			)
		}
	}
}

func (b *Bot) Infoln(f Fields, message string, user string) {
	str := "Info : "
	str += parse(f, message, user)
	str += "\n"

	for _, guild := range b.bot.State.Guilds {
		// Get channels for this guild
		channels, _ := b.bot.GuildChannels(guild.ID)

		for _, c := range channels {
			// Check if channel is a guild text channel and not a voice or DM channel
			if c.Type != discordgo.ChannelTypeGuildText {
				continue
			}

			// Send text message
			b.bot.ChannelMessageSend(
				c.ID,
				str,
			)
		}
	}
}

func (b *Bot) Debug(f Fields, message string, user string) {
	str := "Debug : "
	str += parse(f, message, user)

	for _, guild := range b.bot.State.Guilds {
		// Get channels for this guild
		channels, _ := b.bot.GuildChannels(guild.ID)

		for _, c := range channels {
			// Check if channel is a guild text channel and not a voice or DM channel
			if c.Type != discordgo.ChannelTypeGuildText {
				continue
			}

			// Send text message
			b.bot.ChannelMessageSend(
				c.ID,
				str,
			)
		}
	}
}

func (b *Bot) Debugln(f Fields, message string, user string) {
	str := "Debug : "
	str += parse(f, message, user)
	str += "\n"

	for _, guild := range b.bot.State.Guilds {
		// Get channels for this guild
		channels, _ := b.bot.GuildChannels(guild.ID)

		for _, c := range channels {
			// Check if channel is a guild text channel and not a voice or DM channel
			if c.Type != discordgo.ChannelTypeGuildText {
				continue
			}

			// Send text message
			b.bot.ChannelMessageSend(
				c.ID,
				str,
			)
		}
	}
}

func (b *Bot) Error(f Fields, message string, user string) {
	str := "Error : "
	str += parse(f, message, user)

	for _, guild := range b.bot.State.Guilds {
		// Get channels for this guild
		channels, _ := b.bot.GuildChannels(guild.ID)

		for _, c := range channels {
			// Check if channel is a guild text channel and not a voice or DM channel
			if c.Type != discordgo.ChannelTypeGuildText {
				continue
			}

			// Send text message
			b.bot.ChannelMessageSend(
				c.ID,
				str,
			)
		}
	}
}

func (b *Bot) Errorln(f Fields, message string, user string) {
	str := "Error : "
	str += parse(f, message, user)
	str += "\n"

	for _, guild := range b.bot.State.Guilds {
		// Get channels for this guild
		channels, _ := b.bot.GuildChannels(guild.ID)

		for _, c := range channels {
			// Check if channel is a guild text channel and not a voice or DM channel
			if c.Type != discordgo.ChannelTypeGuildText {
				continue
			}

			// Send text message
			b.bot.ChannelMessageSend(
				c.ID,
				str,
			)
		}
	}
}

func (b *Bot) Warn(f Fields, message string, user string) {
	str := "Warn : "
	str += parse(f, message, user)

	for _, guild := range b.bot.State.Guilds {
		// Get channels for this guild
		channels, _ := b.bot.GuildChannels(guild.ID)

		for _, c := range channels {
			// Check if channel is a guild text channel and not a voice or DM channel
			if c.Type != discordgo.ChannelTypeGuildText {
				continue
			}

			// Send text message
			b.bot.ChannelMessageSend(
				c.ID,
				str,
			)
		}
	}
}

func (b *Bot) Warnln(f Fields, message string, user string) {
	str := "Warn : "
	str += parse(f, message, user)
	str += "\n"

	for _, guild := range b.bot.State.Guilds {
		// Get channels for this guild
		channels, _ := b.bot.GuildChannels(guild.ID)

		for _, c := range channels {
			// Check if channel is a guild text channel and not a voice or DM channel
			if c.Type != discordgo.ChannelTypeGuildText {
				continue
			}

			// Send text message
			b.bot.ChannelMessageSend(
				c.ID,
				str,
			)
		}
	}
}

func (b *Bot) Fatal(f Fields, message string, user string) {
	str := "Fatal : "
	str += parse(f, message, user)

	for _, guild := range b.bot.State.Guilds {
		// Get channels for this guild
		channels, _ := b.bot.GuildChannels(guild.ID)

		for _, c := range channels {
			// Check if channel is a guild text channel and not a voice or DM channel
			if c.Type != discordgo.ChannelTypeGuildText {
				continue
			}

			// Send text message
			b.bot.ChannelMessageSend(
				c.ID,
				str,
			)
		}
	}
}

func (b *Bot) Fatalln(f Fields, message string, user string) {
	str := "Fatal : "
	str += parse(f, message, user)
	str += "\n"

	for _, guild := range b.bot.State.Guilds {
		// Get channels for this guild
		channels, _ := b.bot.GuildChannels(guild.ID)

		for _, c := range channels {
			// Check if channel is a guild text channel and not a voice or DM channel
			if c.Type != discordgo.ChannelTypeGuildText {
				continue
			}

			// Send text message
			b.bot.ChannelMessageSend(
				c.ID,
				str,
			)
		}
	}
}

func (b *Bot) Panic(f Fields, message string, user string) {
	str := "Panic : "
	str += parse(f, message, user)

	for _, guild := range b.bot.State.Guilds {
		// Get channels for this guild
		channels, _ := b.bot.GuildChannels(guild.ID)

		for _, c := range channels {
			// Check if channel is a guild text channel and not a voice or DM channel
			if c.Type != discordgo.ChannelTypeGuildText {
				continue
			}

			// Send text message
			b.bot.ChannelMessageSend(
				c.ID,
				str,
			)
		}
	}
}

func (b *Bot) Panicln(f Fields, message string, user string) {
	str := "Panic : "
	str += parse(f, message, user)
	str += "\n"

	for _, guild := range b.bot.State.Guilds {
		// Get channels for this guild
		channels, _ := b.bot.GuildChannels(guild.ID)

		for _, c := range channels {
			// Check if channel is a guild text channel and not a voice or DM channel
			if c.Type != discordgo.ChannelTypeGuildText {
				continue
			}

			// Send text message
			b.bot.ChannelMessageSend(
				c.ID,
				str,
			)
		}
	}
}
