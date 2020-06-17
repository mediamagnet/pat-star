package commands

import (
	"fmt"
	"pat-star/lib"
	"strings"

	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/atlas"
)

// Help command.
type Help struct{ Command }

var prefix = "." // TODO: Import from yaml

// InitHelp initializes the help command.
func InitHelp() Help {
	return Help{Init(&CommandItem{
		Name:        "help",
		Description: "Shows help message",
		Aliases:     []string{"h", "hh"},
		Usage:       "help role",
		Parameters: []Parameter{
			{
				Name:        "command",
				Description: "Gets help on a specific command",
				Required:    false,
			},
		},
	})}
}

// Register registers and runs the help2 command.
func (c Help) Register() *atlas.Command {

	c.CommandInterface.Run = func(ctx atlas.Context) {
		if len(ctx.Args) > 0 {
			argcmd := ctx.Args[0]

			for _, command := range commands {
				// if argcmd == command.Name then
				// run help, otherwise, likely an
				// alias was used instead
				// which should also work
				if argcmd == command.Name {
					// c.processHelp(ctx, command)
					c.processHelp(ctx, command)
				} else {
					// check if argument was an alias
					for _, alias := range command.Aliases {
						if argcmd == alias {
							// c.processHelp(ctx, command)
							c.processHelp(ctx, command)
						}
					}
				}
			}
		} else {
			var cmdstrslc []string

			for _, command := range commands {
				descslc := strings.Split(command.Description, "\n") // don't want all lines of a command description, just the first
				cmdstrslc = append(cmdstrslc, fmt.Sprintf("`%s%s` - %s", prefix, command.Name, descslc[0]))
			}

			f, t := c.embedFooter(ctx)
			ctx.Atlas.CreateMessage(ctx.Context, ctx.Message.ChannelID, &disgord.CreateMessageParams{
				Embed: &disgord.Embed{
					Fields: []*disgord.EmbedField{
						{
							Name:  "Help",
							Value: "Listing all top-level commands. Specify a command to see more information.",
						},
						{
							Name:  "Commands",
							Value: lib.JoinString(cmdstrslc, "\n"),
						},
					},
					Color:  0x007FFF,
					Footer: f, Timestamp: t,
				},
			})
		}
	}

	return c.CommandInterface
}

// processHelp processes help info defined in each command for command specific help pages
func (c Help) processHelp(ctx atlas.Context, command CommandItem) {
	if command.Admin {
		ctx.Message.Reply(ctx.Context, ctx.Atlas, "Note: Requires some level of permissions")
	}

	embedFields := []*disgord.EmbedField{
		{
			Name:  fmt.Sprintf("%s Help", lib.Ucwords(command.Name)),
			Value: fmt.Sprintf("`%s%s`: %s", prefix, command.Name, command.Description),
		},
	}

	// Usage
	if command.Usage != "" {
		embedFields = append(embedFields, &disgord.EmbedField{
			Name:  "Example Usage",
			Value: fmt.Sprintf("`%s%s`", prefix, command.Usage),
		})
	}

	// Parameters
	if len(command.Parameters) > 0 {
		var params []string

		for _, param := range command.Parameters {
			var (
				paramStr  string
				paramName string
			)

			slParamName := strings.Split(param.Name, ",")

			if len(slParamName) > 1 {
				param.Name = fmt.Sprintf("%s, -%s", slParamName[0], slParamName[1])
			}

			if param.Value != "" {
				paramName = fmt.Sprintf("--%s %s", param.Name, param.Value)
			} else {
				paramName = param.Name
			}

			if param.Required {
				paramStr = fmt.Sprintf("<%s>", paramName)
			} else {
				paramStr = fmt.Sprintf("[%s]", paramName)
			}

			params = append(params, fmt.Sprintf("`%s` - %s",
				paramStr,
				param.Description,
			))
		}

		embedFields = append(embedFields, &disgord.EmbedField{
			Name:  "Parameters",
			Value: lib.JoinString(params, "\n"),
		})
	}

	// Aliases
	if len(command.Aliases) > 0 {
		var aliases []string

		for _, alias := range command.Aliases {
			aliases = append(aliases, fmt.Sprintf("`%s%s`", prefix, alias))
		}

		embedFields = append(embedFields, &disgord.EmbedField{
			Name:  "Aliases",
			Value: lib.JoinString(aliases, ", "),
		})
	}

	f, t := c.embedFooter(ctx)
	ctx.Atlas.CreateMessage(ctx.Context, ctx.Message.ChannelID, &disgord.CreateMessageParams{
		Embed: &disgord.Embed{
			Fields: embedFields,
			Color:  0x007FFF,
			Footer: f, Timestamp: t,
		},
	})
}
