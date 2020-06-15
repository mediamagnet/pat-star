package main

import (
	"pat-star/commands"
	"pat-star/config"

	"github.com/andersfylling/disgord"
	"github.com/sirupsen/logrus"

	"os"

	"github.com/pazuzu156/atlas"
	"github.com/spf13/viper"
)

var log = &logrus.Logger{
	Out:       os.Stderr,
	Formatter: new(logrus.TextFormatter),
	Hooks:     make(logrus.LevelHooks),
	Level:     logrus.InfoLevel,
}

func main() {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	var cfg config.Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	client := atlas.New(&atlas.Options{
		DisgordOptions: disgord.Config{
			BotToken: cfg.Bot.Token,
			Logger:   log,
		},
		OwnerID: cfg.Owner.ID,
	})

	client.Use(atlas.DefaultLogger())
	client.GetPrefix = func(m *disgord.Message) string {
		return cfg.Bot.Prefix
	}

	if err := client.Init(); err != nil {
		panic(err)
	}
}

func init() {
	atlas.Use(commands.InitPing().Register())
	// atlas.Use(commands.InitTiny().Register())
	atlas.Use(commands.InitHelp().Register())
	// atlas.Use(commands.InitRole().Register())
	// atlas.Use(commands.InitPlayer().Register())

}
