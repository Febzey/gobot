package main

import (
	"flag"
	"log"

	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/chat"
	_ "github.com/Tnze/go-mc/data/lang/zh-cn"
	"github.com/febzey/gobot/pkg/events"
	GMMAuth "github.com/maxsupermanhd/go-mc-ms-auth"

	pktid "github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
)

var (
	client *bot.Client
	// player *basic.Player

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

	handlers := map[int32]func(pk.Packet) error{
		pktid.ServerboundChat:       onMessage,
		pktid.ClientboundSystemChat: onMessage,
		pktid.ClientboundPlayerChat: onMessage,
	}

	for id, handler := range handlers {
		id := id
		handler := handler

		client.Events.AddListener(
			bot.PacketHandler{
				Priority: 0,
				ID:       id,
				F: func(p pk.Packet) error {
					return handler(p)
				},
			},
		)
	}

	err = client.JoinServer(*address)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Login success")

	//JoinGame
	err = client.HandleGame()
	if err != nil {
		log.Fatalf("game error: %v", err)
	}

	// for {
	// 	if err = client.HandleGame(); err == nil {
	// 		panic("HandleGame never return nil")
	// 	}

	// 	if err2 := new(bot.PacketHandlerError); errors.As(err, err2) {
	// 		if err := new(DisconnectErr); errors.As(err2, err) {
	// 			log.Print("Disconnect: ", err.Reason)
	// 			return
	// 		} else {
	// 			// print and ignore the error
	// 			log.Print(err2)
	// 		}
	// 	} else {
	// 		log.Fatal(err)
	// 	}
	// }

}

func onMessage(p pk.Packet) error {
	// log.Println("Packet:", p)
	var msg chat.Message

	err := p.Scan(&msg)
	if err != nil {
		return err
	}

	// packetDataByte := p.Data

	// packetDataString := string(packetDataByte)

	//turn p.Data into something readable, it is currently a type "Byte"

	log.Println("Message:", msg.ClearString(), "translate:", msg.Translate)

	return nil
}

// func packetHandler(p pk.Packet) error {
// 	var nodes []Node
// 	var root pk.VarInt
// 	err := p.Scan(pk.Array(&nodes), &root)
// 	if err != nil {
// 		return err
// 	}
// 	log.Printf("Root index: %d", root)
// 	return nil
// }

type Node struct {
}

type DisconnectErr struct {
	Reason chat.Message
}

func (d DisconnectErr) Error() string {
	return "disconnect: " + d.Reason.String()
}
