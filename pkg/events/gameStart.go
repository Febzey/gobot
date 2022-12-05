package events

import (
	"log"

	pk "github.com/Tnze/go-mc/net/packet"
)

func (e *Events) OnLogin(p pk.Packet) error {
	log.Println("Game start")
	return nil
}
