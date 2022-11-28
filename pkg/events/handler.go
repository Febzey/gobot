package events

import (
	"github.com/Tnze/go-mc/bot"
	pktid "github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
)

type Events struct {
	Client *bot.Client
}

func RegisterEvents(client *bot.Client) {

	e := &Events{
		Client: client,
	}

	handlers := map[int32]func(pk.Packet) error{
		pktid.ClientboundLogin:      e.OnLogin,
		pktid.ClientboundSystemChat: e.OnSystemMsg,
		pktid.ClientboundPlayerChat: e.onPlayerMsg,
		pktid.ClientboundSetHealth:  e.handleHealth,
		pktid.ClientboundDisconnect: e.handleDisconnect,
		pktid.ClientboundKeepAlive:  e.handleKeepAlive,
	}

	for id, handler := range handlers {
		id := id
		handler := handler

		client.Events.AddListener(
			bot.PacketHandler{
				Priority: 64,
				ID:       id,
				F: func(p pk.Packet) error {
					return handler(p)
				},
			},
		)
	}
}
