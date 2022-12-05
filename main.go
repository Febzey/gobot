package main

import (
	"flag"
	"log"
	"os"

	"github.com/Tnze/go-mc/bot"
	_ "github.com/Tnze/go-mc/data/lang/en-us"
	"github.com/Tnze/go-mc/data/packetid"

	"github.com/febzey/gobot/pkg/events"
)

var (
	client *bot.Client

	address = flag.String("address", "localhost:53994", "Server address")
)

type Config struct {
	useEvents   bool
	useCommands bool
}

type Bot struct {
	Client *bot.Client
	Logger *log.Logger
	Config *Config

	Events *events.Events

	PlayerInfo
	WorldInfo
	Settings *Settings
}

func newBot() *Bot {
	client = bot.NewClient()

	b := &Bot{
		Client: client,
		Logger: log.New(os.Stdout, "gobot", log.LstdFlags),
		Config: &Config{
			useEvents:   true,
			useCommands: true,
		},
		Events: &events.Events{
			Client: client,
		},
		Settings: &DefaultSettings,
	}

	b.Client.Events.AddListener(
		bot.PacketHandler{Priority: 0, ID: packetid.ClientboundLogin, F: b.Events.OnLogin},
		bot.PacketHandler{Priority: 64, ID: packetid.ClientboundSystemChat, F: b.Events.OnSystemMsg},
		bot.PacketHandler{Priority: 64, ID: packetid.ClientboundPlayerChat, F: b.Events.OnPlayerMsg},
		bot.PacketHandler{Priority: 64, ID: packetid.ClientboundSetHealth, F: b.Events.HandleHealth},
		bot.PacketHandler{Priority: 64, ID: packetid.ClientboundDisconnect, F: b.Events.HandleDisconnect},
		bot.PacketHandler{Priority: 64, ID: packetid.ClientboundKeepAlive, F: b.Events.HandleKeepAlive},
	)

	return b
}

func (b *Bot) Connect(host string) error {
	err := b.Client.JoinServer(host)
	if err != nil {
		return err
	}

	b.Logger.Println("Login success")

	return nil
}

func main() {
	flag.Parse()

	ibot := newBot()

	mauth, err := Authenticate()
	if err != nil {
		ibot.Logger.Fatal(err)
	}

	ibot.Client.Auth = mauth

	err = ibot.Connect(*address)
	if err != nil {
		ibot.Logger.Fatal(err)
	}

	err = ibot.Client.HandleGame()
	if err != nil {
		ibot.Logger.Fatalf("game error: %v", err)
	}

}
