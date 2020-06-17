package commands

import "github.com/pazuzu156/atlas"

// Shove command should work like Donnybrook's .scatter
type Shove struct{ Command }

// InitShove basic stuff for the commmand
func InitShove() Shove {
	return Shove{Init(&CommandItem{
		Name:        "shove",
		Description: "Takes people and shoves them somewhere else",
		Aliases:     []string{"s"},
		Usage:       ".shove <Channel A>...",
		Admin:       true,
		Parameters: []Parameter{
			{
				Name:        "Channel",
				Description: "The channel you want to shove to.",
				Required:    true,
			},
		},
	})}
}

// Register Shove
func (c Shove) Register() *atlas.Command {
	c.CommandInterface.Run = func(ctx atlas.Context) {
		ctx.Message.Reply(ctx.Context, ctx.Atlas, "What if we take these users and shove them somewhere else?")

	}
	return c.CommandInterface
}
