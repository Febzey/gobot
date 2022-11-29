package main

import (
	"flag"
	"log"
	"os"

	"github.com/Tnze/go-mc/bot"
	_ "github.com/Tnze/go-mc/data/lang/en-us"

	"github.com/febzey/gobot/pkg/events"
)

var (
	client *bot.Client

	address = flag.String("address", "localhost:21656", "Server address")
)

type Config struct {
	useEvents   bool
	useCommands bool
}

type Bot struct {
	Client *bot.Client
	Logger *log.Logger
	Config *Config
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
	}

	events.RegisterEvents(client)

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

// for {
// 	if err := client.HandleGame(); err == nil {
// 		panic("HandleGame never return nil")
// 	}
// 	if err2 := new(bot.PacketHandlerError); errors.As(err, err2) {
// 		if err := new(DisconnectErr); errors.As(err2, err) {
// 			log.Print("Disconnect: ", err.Reason)
// 			return
// 		} else {
// 			log.Print(err2)
// 		}
// 	} else {
// 		log.Fatal(err)
// 	}
// }
// type Node struct {
// }

// type DisconnectErr struct {
// 	Reason chat.Message
// }

// func (d DisconnectErr) Error() string {
// 	return "disconnect: " + d.Reason.String()
// }
