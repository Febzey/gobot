package main

import (
	"errors"
	"flag"
	"log"

	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/bot/basic"
	"github.com/Tnze/go-mc/chat"
	_ "github.com/Tnze/go-mc/data/lang/en-us"
	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"

	"github.com/febzey/gobot/pkg/events"
	GMMAuth "github.com/maxsupermanhd/go-mc-ms-auth"
)

var (
	client *bot.Client
	player *basic.Player

	address = flag.String("address", "play.mcvpg.org", "Server address")
)

func authenticate() (bot.Auth, error) {
	mauth, er := GMMAuth.GetMCcredentials("./auth.cache", "88650e7e-efee-4857-b9a9-cf580a00ef43")
	if er != nil {
		return mauth, er
	}
	log.Print("Authenticated as ", mauth.Name, " (", mauth.UUID, ")")
	return mauth, nil
}

func main() {
	flag.Parse()

	mauth, err := authenticate()
	if err != nil {
		log.Fatal(err)
	}

	client = bot.NewClient()
	client.Auth = mauth

	events := events.Events{}
	events.RegisterEvents(client)

	player = basic.NewPlayer(client, basic.DefaultSettings)
	events.Player = player

	err = client.JoinServer(*address)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Login success")

	client.Events.AddListener(
		bot.PacketHandler{
			Priority: 0,
			ID:       packetid.ClientboundKeepAlive,
			F: func(p pk.Packet) error {

				return nil
			},
		},
	)

	//JoinGame
	// err = client.HandleGame()
	// if err != nil {
	// 	log.Fatalf("game error: %v", err)
	// }

	for {
		if err = client.HandleGame(); err == nil {
			panic("HandleGame never return nil")
		}

		if err2 := new(bot.PacketHandlerError); errors.As(err, err2) {
			if err := new(DisconnectErr); errors.As(err2, err) {
				log.Print("Disconnect: ", err.Reason)
				return
			} else {
				// print and ignore the error
				log.Print(err2)
			}
		} else {
			log.Fatal(err)
		}
	}

}

type Node struct {
}

type DisconnectErr struct {
	Reason chat.Message
}

func (d DisconnectErr) Error() string {
	return "disconnect: " + d.Reason.String()
}

// handlers := map[int32]func(pk.Packet) error{
// 	pktid.ClientboundSystemChat: onSystemMsg,
// 	pktid.ClientboundPlayerChat: onMessage,
// }

// for id, handler := range handlers {
// 	id := id
// 	handler := handler

// 	client.Events.AddListener(
// 		bot.PacketHandler{
// 			Priority: 0,
// 			ID:       id,
// 			F: func(p pk.Packet) error {
// 				return handler(p)
// 			},
// 		},
// 	)
// }
